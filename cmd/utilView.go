package cmd

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"filippo.io/age"
	"golang.org/x/term"
)

type ViewOptions struct {
	ShowLocation string
	ShowNames    bool
}

func View(options ViewOptions) {
	ifile, ierr := os.Open(passFile)
	if ierr != nil {
		log.Fatalln(ierr)
	}

	var identity *age.X25519Identity
	var foundIdentity bool
	tr := tar.NewReader(ifile)
	for {
		hdr, herr := tr.Next()
		if herr == io.EOF {
			break
		}
		if herr != nil {
			log.Fatalln(herr)
		}

		if options.ShowLocation != "" && hdr.Name == ".keys/prv.enc" {
			fmt.Fprint(os.Stderr, "Enter password: ")
			b1, _ := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Fprintln(os.Stderr, "")
			pass := string(b1)
			sIdent, iderr := age.NewScryptIdentity(pass)
			if iderr != nil {
				log.Fatalln(iderr)
			}
			sread, derr := age.Decrypt(tr, sIdent)
			if derr != nil {
				log.Fatalln(derr)
			}
			pkeyBuf := &bytes.Buffer{}
			io.Copy(pkeyBuf, sread)
			var ierr error
			identity, ierr = age.ParseX25519Identity(pkeyBuf.String())
			if ierr != nil {
				log.Fatalln("unable to decrypt private key")
			}
			foundIdentity = true
		}

		if hdr.Name == options.ShowLocation {
			if !foundIdentity {
				log.Fatalln("unable to find private key")
			}
			r, derr := age.Decrypt(tr, identity)
			if derr != nil {
				log.Fatalln(derr)
			}
			io.Copy(os.Stdout, r)
		}

		if options.ShowNames && !strings.HasPrefix(hdr.Name, ".keys") {
			fmt.Println(strings.TrimSuffix(hdr.Name, ".age"))
		}
	}
}
