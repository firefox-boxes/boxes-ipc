package main

import "github.com/firefox-boxes/boxes"

func main() {
	p := boxes.Probe()
	attrDB := InitAttrDB(p.GetRelDir("attributes.db"))
	ipc := IPC{p:p,attrDB:attrDB}
	
	defaultBox := attrDB.GetDefault()
	if defaultBox != "<nobody>" {
		req := "exec " + defaultBox
		res := ""
		ipc.Handle(&req, &res)
	}

	StartIPC(&ipc)
}