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
)

type EditOptions struct {
	InsertLocation string
	InsertData     string
	RemoveLocation string
}

func Edit(options EditOptions) {
	ifile, ierr := os.Open(passFile)
	if ierr != nil {
		log.Fatalln(ierr)
	}

	ofile, oerr := os.Create(passFile + ".new")
	if oerr != nil {
		log.Fatalln(oerr)
	}
	tw := tar.NewWriter(ofile)

	var recipient *age.X25519Recipient
	var foundRecipient bool

	tr := tar.NewReader(ifile)
	for {
		hdr, herr := tr.Next()
		if herr == io.EOF {
			break
		}
		if herr != nil {
			log.Fatalln(herr)
		}

		var buf bytes.Buffer
		io.Copy(&buf, tr)
		if hdr.Name == ".keys/pub.txt" {
			var rerr error
			recipient, rerr = age.ParseX25519Recipient(buf.String())
			if rerr != nil {
				foundRecipient = false
				log.Println(rerr)
			}
			foundRecipient = true
		}

		// Replacing / Removing
		if hdr.Name == options.InsertLocation || hdr.Name == options.RemoveLocation {
			continue
		}

		tw.WriteHeader(hdr)
		tw.Write(buf.Bytes())
	}
	ifile.Close()

	if foundRecipient && options.InsertLocation != "" {
		pdata := strings.NewReader(options.InsertData)

		pout := &bytes.Buffer{}
		w, _ := age.Encrypt(pout, recipient)
		io.Copy(w, pdata)
		w.Close()

		tw.WriteHeader(bufToHeader(options.InsertLocation, pout))
		io.Copy(tw, pout)

		fmt.Println("entry added at", options.InsertLocation)
	}

	tw.Close()
	ofile.Close()

	os.Rename(passFile+".new", passFile)

	if options.RemoveLocation != "" {
		fmt.Println("entry removed at", options.RemoveLocation)
	}
}
