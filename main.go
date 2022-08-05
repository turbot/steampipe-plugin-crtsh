package main

import (
	"github.com/turbot/steampipe-plugin-crtsh/crtsh"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: crtsh.Plugin})
}
