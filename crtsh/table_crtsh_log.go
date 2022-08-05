package crtsh

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableCrtshLog() *plugin.Table {
	return &plugin.Table{
		Name:        "crtsh_log",
		Description: "Certificate transparency log operators that track and record log entries.",
		List: &plugin.ListConfig{
			Hydrate: listLog,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "ID of the log."},
			{Name: "operator", Type: proto.ColumnType_STRING, Description: "Operator of the log."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the log."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the log."},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "True if the log is active."},
			{Name: "apple_inclusion_status", Type: proto.ColumnType_STRING, Description: "Status of this log with Apple."},
			{Name: "chrome_inclusion_status", Type: proto.ColumnType_STRING, Description: "Status of this log in Google Chrome."},
			// Other columns
			{Name: "apple_last_status_change", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the status of this log last changed with Apple."},
			{Name: "batch_size", Type: proto.ColumnType_INT, Description: "Batch size of the log."},
			{Name: "chrome_disqualified_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when Google Chrome disqualified the log."},
			{Name: "chrome_final_tree_size", Type: proto.ColumnType_INT, Description: "Final tree size of the log according to Google Chrome."},
			{Name: "chrome_issue_number", Type: proto.ColumnType_INT, Description: "Issue number discussing inclusion of the log in Google Chrome."},
			{Name: "chrome_version_added", Type: proto.ColumnType_INT, Description: "Version when the log was included in Google Chrome, if any."},
			{Name: "chunk_size", Type: proto.ColumnType_INT, Description: "Chunk size of the log."},
			{Name: "google_uptime", Type: proto.ColumnType_STRING, Description: "Uptime percentage of the log according to Google."},
			{Name: "latest_update", Type: proto.ColumnType_TIMESTAMP, Description: "Latest time when the log was contacted by crt.sh."},
			{Name: "latest_sth_timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Latest Signed Tree Head (STH) timestamp of the log."},
			{Name: "mmd_in_seconds", Type: proto.ColumnType_INT, Description: "Maximum Merge Delay of the log."},
			{Name: "public_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("PublicKey").Transform(byteArrayToString), Description: "Public key of the log."},
			{Name: "tree_size", Type: proto.ColumnType_INT, Description: "Tree size is the total number of nodes in the merkle tree for the log."},
		},
	}
}

type logRow struct {
	ID                    int        `db:"id"`
	Operator              string     `db:"operator"`
	URL                   string     `db:"url"`
	Name                  string     `db:"name"`
	PublicKey             []byte     `db:"public_key"`
	IsActive              bool       `db:"is_active"`
	LatestUpdate          *time.Time `db:"latest_update"`
	LatestSthTimestamp    *time.Time `db:"latest_sth_timestamp"`
	MmdInSeconds          *int       `db:"mmd_in_seconds"`
	TreeSize              *int       `db:"tree_size"`
	BatchSize             *int       `db:"batch_size"`
	ChunkSize             *int       `db:"chunk_size"`
	GoogleUptime          *string    `db:"google_uptime"`
	ChromeVersionAdded    *int       `db:"chrome_version_added"`
	ChromeInclusionStatus *string    `db:"chrome_inclusion_status"`
	ChromeIssueNumber     *int       `db:"chrome_issue_number"`
	ChromeFinalTreeSize   *int       `db:"chrome_final_tree_size"`
	ChromeDisqualifiedAt  *time.Time `db:"chrome_disqualified_at"`
	AppleInclusionStatus  *string    `db:"apple_inclusion_status"`
	AppleLastStatusChange *time.Time `db:"apple_last_status_change"`
}

func listLog(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_log.listLog", "connection_error", err)
		return nil, err
	}

	q := `
	select
		id,
		operator,
		url,
		name,
		public_key,
		is_active,
		latest_update,
		latest_sth_timestamp,
		mmd_in_seconds,
		tree_size,
		batch_size,
		chunk_size,
		google_uptime,
		chrome_version_added,
		chrome_inclusion_status,
		chrome_issue_number,
		chrome_final_tree_size,
		chrome_disqualified_at,
		apple_inclusion_status,
		apple_last_status_change
	from
	  ct_log`

	q, args := queryWithQuals(ctx, d, q)

	i := logRow{}
	rows, err := db.QueryxContext(ctx, q, args...)
	if err != nil {
		plugin.Logger(ctx).Error("crtsh_log.listLog", "query_error", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&i)
		if err != nil {
			plugin.Logger(ctx).Error("crtsh_log.listLog", "row_error", err)
			continue
		}
		d.StreamListItem(ctx, i)
	}

	return nil, err
}
