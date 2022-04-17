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
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// poolsCmd represents the pools command
var poolsCmd = &cobra.Command{
	Use: "pools",
}

// poolsGetCmd represents the poolsGet command
var poolsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get pool configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		data, err := client.GetPools()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		jsonData, err := json.Marshal(data.Data)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(poolsCmd)
	poolsCmd.AddCommand(poolsGetCmd)
}
