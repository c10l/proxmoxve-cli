package cmd

import (
	"github.com/c10l/proxmoxve-client-go/api2"
	"github.com/spf13/viper"
)

func newClient() *api2.Client {
	return api2.NewClient(
		viper.GetString("url"),
		viper.GetString("token-id"),
		viper.GetString("secret"),
		viper.GetBool("insecure"))
}
