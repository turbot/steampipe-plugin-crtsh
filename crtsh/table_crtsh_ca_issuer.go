package crtsh

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableCrtshCaIssuer() *plugin.Table {
	return &plugin.Table{
		Name:        "crtsh_ca_issuer",
		Description: "Certificate Authority Issuers known to crt.sh, including the status of their last check.",
		List: &plugin.ListConfig{
			Hydrate: listCaIssuer,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "ca_id", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "next_check_due", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "last_checked", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "url", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "result", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "first_certificate_id", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "is_active", Operators: []string{"=", "<>"}, Require: plugin.Optional},
				{Name: "content_type", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ca_id", Type: proto.ColumnType_INT, Description: "Unique ID of the CA represented by this issuer record."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "result", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "ca_certificate_ids", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "first_certificate_id", Type: proto.ColumnType_INT, Description: ""},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "content_type", Type: proto.ColumnType_STRING, Description: ""},
			// Other columns
			{Name: "next_check_due", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "last_checked", Type: proto.ColumnType_TIMESTAMP, Description: ""},
		},
	}
}

type caIssuerRow struct {
	CaID               int             `db:"ca_id"`
	URL                string          `db:"url"`
	Result             *string         `db:"result"`
	CaCertificateIds   *types.JSONText `db:"ca_certificate_ids"`
	FirstCertificateID *int64          `db:"first_certificate_id"`
	IsActive           *bool           `db:"is_active"`
	ContentType        *string         `db:"content_type"`
	NextCheckDue       *time.Time      `db:"next_check_due"`
	LastChecked        *time.Time      `db:"last_checked"`
}

func listCaIssuer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_ca_issuer.listCaIssuer", "connection_error", err)
		return nil, err
	}

	q := `
		select
			ca_id,
			url,
			result,
			to_jsonb(ca_certificate_ids) as ca_certificate_ids,
			first_certificate_id,
			is_active,
			content_type,
			next_check_due,
			last_checked
		from ca_issuer
	`

	q, args := queryWithQuals(ctx, d, q)

	i := caIssuerRow{}
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_ca_issuer.listCaIssuer", "query_error", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&i)
		if err != nil {
			plugin.Logger(ctx).Error("crtsh_ca_issuer.listCaIssuer", "row_error", err)
			continue
		}
		d.StreamListItem(ctx, i)
	}

	return nil, err
}
