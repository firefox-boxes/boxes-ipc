package main

import (
	"net"
	"net/rpc"

	"github.com/firefox-boxes/boxes-ipc/logging"

	"github.com/firefox-boxes/boxes"
)

const IPC_PORT = ":6688"

type IPC struct {
	attrDB AttrDB
	p boxes.ProbeResult
}

func (ipc *IPC) Handle(req *string, res *string) error {
	i := logging.IPCCmdID()
	logging.Info("%v !%v", i, *req)
	*res = Handle(*req, ipc.attrDB, ipc.p)
	logging.Info("%v =%v", i, logging.ProcessStr(*res, 40))
	return nil
}

func StartIPC(ipc *IPC) {
	rpc.Register(ipc)
	listener, _ := net.Listen("tcp", IPC_PORT)
	defer listener.Close()
	logging.Info("Listening on 127.0.0.1%v", IPC_PORT)
	rpc.Accept(listener)
}