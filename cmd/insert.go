/*
Copyright © 2021 Chris Howey <chris@howey.me>
Use of source code is governed by a ISC-style
license that can be found in the LICENSE file.
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert [location]",
	Short: "insert a password/note at location",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("specify location")
			return
		}
		location := args[0] + ".age"

		fmt.Println("reading entry from stdin:")
		pout := &bytes.Buffer{}
		io.Copy(pout, os.Stdin)

		Edit(EditOptions{
			InsertLocation: location,
			InsertData:     pout.String(),
		})
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}
