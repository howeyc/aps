package cmd

import (
	"archive/tar"
	"bytes"
	"time"
)

func bufToHeader(name string, buf *bytes.Buffer) *tar.Header {
	return &tar.Header{
		Name:    name,
		Mode:    0400,
		ModTime: time.Now(),
		Size:    int64(buf.Len()),
	}
}
