package main

import (
	"net"
	"net/rpc"

	"github.com/firefox-boxes/boxes"
)

type IPC struct {
	attrDB AttrDB
	p boxes.ProbeResult
}

func (ipc *IPC) Handle(req *string, res *string) error {
	*res = Handle(*req, ipc.attrDB, ipc.p)
	return nil
}

func StartIPC(ipc *IPC) {
	rpc.Register(ipc)
	listener, _ := net.Listen("tcp", ":6688")
	defer listener.Close()
	rpc.Accept(listener)
}