package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-avi/avi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: avi.Provider})
}
