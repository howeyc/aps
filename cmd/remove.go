/*
Copyright © 2021 Chris Howey <chris@howey.me>
Use of source code is governed by a ISC-style
license that can be found in the LICENSE file.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [location]",
	Short: "remove an entry at specified location",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("specify location")
			return
		}
		location := args[0] + ".age"

		Edit(EditOptions{
			RemoveLocation: location,
		})
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
