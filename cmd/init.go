/*
Copyright © 2021 Chris Howey <chris@howey.me>
Use of source code is governed by a ISC-style
license that can be found in the LICENSE file.
*/
package cmd

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"filippo.io/age"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password store",
	Run: func(cmd *cobra.Command, args []string) {
		if _, serr := os.Stat(passFile); !os.IsNotExist(serr) {
			fmt.Fprintf(os.Stderr, "Password file exists!? Overwrite? [y/N]")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" {
				os.Exit(0)
			}
		}

		var pass1, pass2 string
		for pass1 != pass2 || (pass1 == "" && pass2 == "") {
			fmt.Fprint(os.Stderr, "Enter password: ")
			b1, e1 := term.ReadPassword(int(os.Stdin.Fd()))
			if e1 != nil {
				log.Fatalln(e1)
			}
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprint(os.Stderr, "Repeat password: ")
			b2, e2 := term.ReadPassword(int(os.Stdin.Fd()))
			if e2 != nil {
				log.Fatalln(e2)
			}
			fmt.Fprintln(os.Stderr, "")

			pass1, pass2 = string(b1), string(b2)
		}

		ofile, oerr := os.Create(passFile)
		if oerr != nil {
			log.Fatalln(oerr)
		}
		defer ofile.Close()
		ageIdentity, genErr := age.GenerateX25519Identity()
		if genErr != nil {
			log.Fatalln(genErr)
		}

		scryptRecipient, serr := age.NewScryptRecipient(pass1)
		if serr != nil {
			log.Fatalln(serr)
		}

		pfile := tar.NewWriter(ofile)

		sout := bytes.Buffer{}
		w, _ := age.Encrypt(&sout, scryptRecipient)
		io.WriteString(w, ageIdentity.String())
		w.Close()

		pfile.WriteHeader(bufToHeader(".keys/prv.enc", &sout))
		io.Copy(pfile, &sout)

		agePub := bytes.NewBufferString(ageIdentity.Recipient().String())
		pfile.WriteHeader(bufToHeader(".keys/pub.txt", agePub))
		io.Copy(pfile, agePub)

		pfile.Close()

		fmt.Println("password store created in file: ", passFile)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
