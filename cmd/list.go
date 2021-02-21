/*
Copyright © 2021 Chris Howey <chris@howey.me>
Use of source code is governed by a ISC-style
license that can be found in the LICENSE file.

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all entry names",
	Run: func(cmd *cobra.Command, args []string) {
		View(ViewOptions{ShowNames: true})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
