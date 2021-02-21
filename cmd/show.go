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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [location]",
	Short: "show entry at location",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("specify location")
			return
		}
		location := args[0] + ".age"

		View(ViewOptions{ShowLocation: location})
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
