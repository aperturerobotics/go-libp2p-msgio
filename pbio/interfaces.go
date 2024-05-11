// Package pbio reads and writes varint-prefix protobufs, using Google's Protobuf package.
package pbio

import (
	"io"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
)

type Writer interface {
	WriteMsg(protobuf_go_lite.Message) error
}

type WriteCloser interface {
	Writer
	io.Closer
}

type Reader interface {
	ReadMsg(msg protobuf_go_lite.Message) error
}

type ReadCloser interface {
	Reader
	io.Closer
}
