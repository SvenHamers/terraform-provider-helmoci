package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/svenhamers/terraform-provider-helmoci/helmoci"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: helmoci.Provider,
	})
}
