package crtsh

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableCrtshCa() *plugin.Table {
	return &plugin.Table{
		Name:        "crtsh_ca",
		Description: "TODO",
		List: &plugin.ListConfig{
			Hydrate: listCa,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "name", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "num_certs_issued", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "num_precerts_issued", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "num_certs_expired", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "num_precerts_expired", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "linting_applies", Operators: []string{"=", "<>"}, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier of the CA."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the CA."},
			{Name: "num_certs_issued", Type: proto.ColumnType_INT, Description: "Number of certificates issued by the CA."},
			{Name: "num_precerts_issued", Type: proto.ColumnType_INT, Description: "Number of pre-certificates issued by the CA."},
			{Name: "num_certs_expired", Type: proto.ColumnType_INT, Description: "Number of certificates from the CA that have expired."},
			{Name: "num_precerts_expired", Type: proto.ColumnType_INT, Description: "Number of pre-certificates from the CA that have expired."},
			{Name: "linting_applies", Type: proto.ColumnType_BOOL, Description: ""},
			// The public key format is a mystery to me, and I'm not sure it's even
			// valuable, so I'm leaving it out for now.
			//{Name: "public_key", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

type caRow struct {
	ID                 int        `db:"id"`
	Name               string     `db:"name"`
	NumCertsIssued     *int64     `db:"num_certs_issued"`
	NumPrecertsIssued  *int64     `db:"num_precerts_issued"`
	NumCertsExpired    *int64     `db:"num_certs_expired"`
	NumPrecertsExpired *int64     `db:"num_precerts_expired"`
	LastNotAfter       *time.Time `db:"last_not_after"`
	NextNotAfter       *time.Time `db:"next_not_after"`
	LintingApplies     bool       `db:"linting_applies"`
	//PublicKey          []byte  `db:"public_key"`
}

func listCa(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_ca.listCa", "connection_error", err)
		return nil, err
	}

	// Use a CTE query to setup the column names we need for use with
	// queryWithQuals.
	//
	// Note: the last_not_after and next_not_after fields appear to be related to
	// crt.sh backend functionality to maintain the expired certs count. I
	// believe they are tracking timestamps to use when searching the database
	// and updating those counts. So, they are not included in our results.
	//
	// If using the *_not_after columns then Some rows (2 of 236k on 2-Jun-2022)
	// have `infinity` as the value here, which does not compile properly into a
	// time.Time.  They will be logged as errors, but are ignored because of this
	// decision. A string mapping would just be too inconvenient relative to the
	// value.
	//
	q := `
		with ca_expanded as (
			select
				id,
				num_issued[1] as num_certs_issued,
				num_issued[2] as num_precerts_issued,
				num_expired[1] as num_certs_expired,
				num_expired[2] as num_precerts_expired,
				-- last_not_after,
				-- next_not_after,
				linting_applies,
				name
			from ca
		)
		select * from ca_expanded
	`

	q, args := queryWithQuals(ctx, d, q)

	i := caRow{}
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_ca.listCa", "query_error", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&i)
		if err != nil {
			plugin.Logger(ctx).Error("crtsh_ca.listCa", "row_error", err)
			continue
		}
		d.StreamListItem(ctx, i)
	}

	return nil, err
}
