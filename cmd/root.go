/*
Copyright © 2022 Cassiano Leal cassiano@c10l.cc

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ProxMoxURL string
var TokenID string
var Secret string
var TLSInsecure bool

var rootCmd = &cobra.Command{
	Use:   "proxmoxve-cli",
	Short: "ProxMox VE CLI",
	Long: `CLI to interact with the ProxMox VE JSON API.

It uses the API client library proxmoxve-client-go from github.com/c10l/proxmoxve-client-go

Official API docs: https://pve.proxmox.com/pve-docs/api-viewer/`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(markRequiredFlags)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.proxmoxve-cli.yaml)")

	rootCmd.PersistentFlags().StringP("url", "u", "", "(required) Root URL of the ProxMox VE server. e.g. https://proxmox.example.com:8006")
	rootCmd.PersistentFlags().String("token-id", "", "(required) ProxMox API Token ID")
	rootCmd.PersistentFlags().String("secret", "", "(required) ProxMox API Secret")
	rootCmd.PersistentFlags().BoolVarP(&TLSInsecure, "insecure", "k", false, "Allow untrusted connections to the API.")
	cobra.CheckErr(viper.BindPFlags(rootCmd.PersistentFlags()))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".proxmoxve-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".proxmoxve-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func markRequiredFlags() {
	if !viper.IsSet("url") {
		rootCmd.MarkPersistentFlagRequired("url")
	}
	if !viper.IsSet("token-id") {
		rootCmd.MarkPersistentFlagRequired("token-id")
	}
	if !viper.IsSet("secret") {
		rootCmd.MarkPersistentFlagRequired("secret")
	}
}
