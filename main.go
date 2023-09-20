package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/tonyfairbanks/terraform-provider-graylog/graylog"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/tonyfairbanks/graylog",
		ProviderFunc: func() *schema.Provider {
			return graylog.Provider()
		},
	}

	plugin.Serve(opts)
}
