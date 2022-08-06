package crtsh

import (
	"context"
	"crypto/x509"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableCrtshCertificate() *plugin.Table {
	return &plugin.Table{
		Name:        "crtsh_certificate",
		Description: "Certificates recorded in transparency logs.",
		List: &plugin.ListConfig{
			Hydrate: listCertificate,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.AnyOf},
				{Name: "query", Require: plugin.AnyOf, CacheMatch: "exact"},
				{Name: "not_after", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("CertificateID"), Description: "Unique ID of the certificate in crt.sh."},
			{Name: "dns_names", Type: proto.ColumnType_JSON, Hydrate: parseCertificate, Description: "DNS names represented by the certificate, e.g. steampipe.io"},
			{Name: "not_before", Type: proto.ColumnType_TIMESTAMP, Hydrate: parseCertificate, Description: "The certificate invalid before this time."},
			{Name: "not_after", Type: proto.ColumnType_TIMESTAMP, Description: "The certificate is invalid after this time."},
			{Name: "subject", Type: proto.ColumnType_JSON, Hydrate: parseCertificate, Description: "Details about the Subject of the certificate, e.g. CommonName, OrganizationalUnit, etc."},
			// Other columns
			{Name: "email_addresses", Type: proto.ColumnType_JSON, Hydrate: parseCertificate, Description: "Email addresses associated with the certificate."},
			{Name: "fingerprint_sha1", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate").Transform(sha1Fingerprint), Description: "SHA1 fingerprint of the certificate, e.g. abcd12..."},
			{Name: "fingerprint_sha256", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate").Transform(sha256Fingerprint), Description: "SHA256 fingerprint of the certificate, e.g. abcd12..."},
			{Name: "ip_addresses", Type: proto.ColumnType_JSON, Hydrate: parseCertificate, Description: "IP addresses associated with the certificate."},
			{Name: "is_ca", Type: proto.ColumnType_BOOL, Hydrate: parseCertificate, Transform: transform.FromField("IsCA"), Description: "True if this certificate is a Certificate Authority."},
			{Name: "issuer", Type: proto.ColumnType_JSON, Hydrate: parseCertificate, Description: "Details about the Certificate Authority who issued the certificate, e.g. CommonName, "},
			{Name: "issuer_ca_id", Type: proto.ColumnType_INT, Description: "ID of the Certificate Authority who issued the certificate."},
			{Name: "public_key", Type: proto.ColumnType_STRING, Hydrate: parseCertificate, Transform: transform.FromField("PublicKey").Transform(publicKeyToPem), Description: "Public key of the certificate in PEM format."},
			{Name: "public_key_algorithm", Type: proto.ColumnType_STRING, Hydrate: parseCertificate, Description: "Algorithm used for the public key. e.g. RSA."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "The query provided for the certificate search."},
			{Name: "serial_number", Type: proto.ColumnType_STRING, Hydrate: parseCertificate, Transform: transform.FromField("SerialNumber").Transform(serialNumberToHex), Description: "Unique identifier assigned by the Certificate Authority who issued the certificate."},
			{Name: "signature_algorithm", Type: proto.ColumnType_STRING, Hydrate: parseCertificate, Description: "Algorithm used for the signature, e.g. SHA256-RSA."},
			{Name: "uris", Type: proto.ColumnType_JSON, Hydrate: parseCertificate, Description: "URIs associated with the certificate."},
			{Name: "version", Type: proto.ColumnType_INT, Hydrate: parseCertificate, Description: "Version of the certificate, e.g. 3."},
			// Large columns
			{Name: "certificate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate").Transform(byteArrayToString), Description: "Full raw certificate string in hex format."},
		},
	}
}

type certificateRow struct {
	CertificateID int        `db:"certificate_id"`
	IssuerCaID    int        `db:"issuer_ca_id"`
	NameType      string     `db:"name_type"`
	NameValue     string     `db:"name_value"`
	Certificate   []byte     `db:"certificate"`
	NotAfter      *time.Time `db:"not_after"`
}

func parseCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cr := h.Item.(certificateRow)
	cert, err := x509.ParseCertificate(cr.Certificate)
	if err != nil {
		return x509.Certificate{}, nil
	}
	return cert, nil
}

func listCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_certificate.listCertificate", "connection_error", err)
		return nil, err
	}

	q := `
		select distinct on (certificate_id)
			certificate_id,
			issuer_ca_id,
			name_type,
			name_value,
			certificate,
			x509_notAfter(certificate) as not_after
		from
			certificate_and_identities
	`

	quals := d.KeyColumnQuals
	whereClauses := []string{}
	args := []interface{}{}

	if quals["id"] != nil {
		args = append(args, quals["id"].GetInt64Value())
		whereClauses = append(whereClauses, fmt.Sprintf("certificate_id = $%d", len(args)))
	}

	if quals["query"] != nil {
		args = append(args, quals["query"].GetStringValue())
		whereClauses = append(whereClauses, fmt.Sprintf("plainto_tsquery('certwatch', $%d) @@ identities(certificate)", len(args)))
		whereClauses = append(whereClauses, fmt.Sprintf("name_value ilike ('%%' || $%d || '%%')", len(args)))
	}

	q = q + " where " + strings.Join(whereClauses, " and ")

	plugin.Logger(ctx).Debug("crtsh_ca_issuer.listCertificate", "query", regexp.MustCompile(`(?m)[\s\n]+`).ReplaceAllString(q, " "))
	plugin.Logger(ctx).Debug("crtsh_ca_issuer.listCertificate", "args", args)

	i := certificateRow{}
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_certificate.listCertificate", "query_error", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&i)
		if err != nil {
			plugin.Logger(ctx).Error("crtsh_certificate.listCertificate", "row_error", err)
			continue
		}
		d.StreamListItem(ctx, i)
	}

	return nil, err
}
