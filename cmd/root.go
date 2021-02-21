/*
Copyright © 2021 Chris Howey <chris@howey.me>
Use of source code is governed by a ISC-style
license that can be found in the LICENSE file.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var passFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aps",
	Short: "A Password Store",
	Long: `A Password Store keeps your passwords encrypted with age 
inside a tar file.

Only the content of each password note is encrypted with age.
Each entry is a "file" inside tar. This allows for listing
to occur without the need to "unlock" or decrypt any entries.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&passFile, "file", "", "password file (default is aps.tar)")

	viper.SetEnvPrefix("aps")
	viper.BindEnv("file")

	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
	viper.SetDefault("file", "aps.tar")

	viper.SetConfigType("toml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	passFile = viper.GetString("file")
}
