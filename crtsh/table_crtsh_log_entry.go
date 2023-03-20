package crtsh

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCrtshLogEntry() *plugin.Table {
	return &plugin.Table{
		Name:        "crtsh_log_entry",
		Description: "Certificate transparency log entries recorded for each certificate.",
		List: &plugin.ListConfig{
			Hydrate: listLogEntry,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "certificate_id", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.AnyOf},
				{Name: "entry_id", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.AnyOf},
				{Name: "entry_timestamp", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
				{Name: "ct_log_id", Operators: []string{">", ">=", "=", "<", "<=", "<>"}, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ct_log_id", Type: proto.ColumnType_INT, Description: "The log this entry is defined in."},
			{Name: "entry_id", Type: proto.ColumnType_INT, Description: "Unique ID of the entry."},
			{Name: "entry_timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the entry."},
			{Name: "certificate_id", Type: proto.ColumnType_INT, Description: "Certificate the entry represents."},
		},
	}
}

type logEntryRow struct {
	CertificateID  int        `db:"certificate_id"`
	EntryID        int        `db:"entry_id"`
	EntryTimestamp *time.Time `db:"entry_timestamp"`
	CtLogID        int        `db:"ct_log_id"`
}

func listLogEntry(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_log.listLogEntry", "connection_error", err)
		return nil, err
	}

	q := `
		select
			certificate_id,
			entry_id,
			entry_timestamp,
			ct_log_id
		from
			ct_log_entry
	`

	q, args := queryWithQuals(ctx, d, q)

	i := logEntryRow{}
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_log.listLogEntry", "query_error", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&i)
		if err != nil {
			plugin.Logger(ctx).Error("crtsh_log.listLogEntry", "row_error", err)
			continue
		}
		d.StreamListItem(ctx, i)
	}

	return nil, err
}
