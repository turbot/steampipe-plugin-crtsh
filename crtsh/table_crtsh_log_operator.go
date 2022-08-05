package crtsh

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableCrtshLogOperator() *plugin.Table {
	return &plugin.Table{
		Name:        "crtsh_log_operator",
		Description: "Log operators used by crt.sh.",
		List: &plugin.ListConfig{
			Hydrate: listLogOperator,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "operator", Type: proto.ColumnType_STRING, Description: "Name of the operator."},
		},
	}
}

type operatorRow struct {
	Operator string `db:"operator"`
}

func listLogOperator(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_log_operator.listLogOperator", "connection_error", err)
		return nil, err
	}

	q := `select operator from ct_log_operator`

	q, args := queryWithQuals(ctx, d, q)

	i := operatorRow{}
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_log_operator.listLogOperator", "query_error", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&i)
		if err != nil {
			plugin.Logger(ctx).Error("crtsh_log_operator.listLogOperator", "row_error", err)
			continue
		}
		d.StreamListItem(ctx, i)
	}

	return nil, err
}
