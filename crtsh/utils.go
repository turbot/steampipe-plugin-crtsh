package crtsh

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func byteArrayToString(_ context.Context, d *transform.TransformData) (interface{}, error) {
	ba := d.Value.([]byte)
	hexString := hex.EncodeToString(ba)
	re := regexp.MustCompile("..")
	return strings.TrimRight(re.ReplaceAllString(hexString, "$0:"), ":"), nil
}

// Format is lowercase string with no colons (e.g. abcd01...). This is the
// consistent with crt.sh and SSLLabs.
// Note that Google Chrome does AB CD 01 format.
func sha1Fingerprint(_ context.Context, d *transform.TransformData) (interface{}, error) {
	ba := d.Value.([]byte)
	sum := sha1.Sum(ba)
	sumSlice := sum[:]
	hexString := hex.EncodeToString(sumSlice)
	return hexString, nil
}

// Format is lowercase string with no colons (e.g. abcd01...). This is the
// consistent with crt.sh and SSLLabs.
// Note that Google Chrome does AB CD 01 format.
func sha256Fingerprint(_ context.Context, d *transform.TransformData) (interface{}, error) {
	ba := d.Value.([]byte)
	sum := sha256.Sum256(ba)
	sumSlice := sum[:]
	hexString := hex.EncodeToString(sumSlice)
	return hexString, nil
}

func serialNumberToHex(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(*big.Int)
	if i == nil {
		return nil, nil
	}
	hexString := fmt.Sprintf("%036x", i)
	re := regexp.MustCompile("..")
	return strings.TrimRight(re.ReplaceAllString(hexString, "$0:"), ":"), nil
}

func publicKeyToPem(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	var result interface{}
	switch d.Value.(type) {
	case *ecdsa.PublicKey:
		result, _ = ecdsaPublicKeyToPem(ctx, d)
	default:
		// Assume RSA
		result, _ = rsaPublicKeyToPem(ctx, d)
	}
	return result, nil
}

func rsaPublicKeyToPem(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	k := d.Value.(*rsa.PublicKey)
	if k == nil {
		return nil, nil
	}
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(k)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh.rsaPublicKeyToPem", "parse_error", err)
		return nil, err
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)
	return pubkeyPem, nil
}

func ecdsaPublicKeyToPem(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	k := d.Value.(*ecdsa.PublicKey)
	if k == nil {
		return nil, nil
	}
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(k)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh.ecdsaPublicKeyToPem", "parse_error", err)
		return nil, err
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)
	return pubkeyPem, nil
}

func connect(ctx context.Context, d *plugin.QueryData) (*sqlx.DB, error) {

	cacheKey := "crtsh"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		conn := cachedData.(*sqlx.DB)
		if err := conn.PingContext(ctx); err == nil {
			// Cached connection is good, return it
			return conn, nil
		}
	}

	connString := "postgres://guest@crt.sh:5432/certwatch?binary_parameters=yes"
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_ca_issuer.listCaIssuer", "connection_error", err)
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, db)

	return db, err
}

func queryWithQuals(ctx context.Context, d *plugin.QueryData, inputQuery string) (query string, args []interface{}) {

	query = inputQuery

	whereClauses := []string{}
	for _, col := range d.Table.Columns {
		if d.Quals[col.Name] == nil {
			continue
		}
		for _, q := range d.Quals[col.Name].Quals {
			switch col.Type.String() {
			case "STRING":
				args = append(args, q.Value.GetStringValue())
			case "TIMESTAMP":
				args = append(args, q.Value.GetTimestampValue().AsTime().Format(time.RFC3339))
			case "INT":
				args = append(args, q.Value.GetInt64Value())
			case "BOOL":
				args = append(args, q.Value.GetBoolValue())
			default:
				continue
			}
			whereClauses = append(whereClauses, fmt.Sprintf("%s %s $%d", col.Name, q.Operator, len(args)))
		}
	}
	if len(whereClauses) > 0 {
		query = query + "where " + strings.Join(whereClauses, " and ")
	}

	limit := d.QueryContext.Limit
	if limit != nil {
		args = append(args, *limit)
		query = fmt.Sprintf("%s limit $%d", query, len(args))
	}

	plugin.Logger(ctx).Debug("crtsh_ca_issuer.listCaIssuer", "query", regexp.MustCompile(`(?m)[\s\n]+`).ReplaceAllString(query, " "))
	plugin.Logger(ctx).Debug("crtsh_ca_issuer.listCaIssuer", "args", args)

	return
}
