package crtsh

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-crtsh",
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"crtsh_ca":          tableCrtshCa(),
			"crtsh_ca_issuer":   tableCrtshCaIssuer(),
			"crtsh_certificate": tableCrtshCertificate(),
			"crtsh_log":         tableCrtshLog(),
			"crtsh_log_entry":   tableCrtshLogEntry(),
		},
	}
	return p
}
