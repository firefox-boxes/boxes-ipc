package main

import "net/rpc"

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6688")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	req := "default:get"
	res := ""
	client.Call("IPC.Handle", &req, &res)
	req = "exec " + res
	res = ""
	client.Call("IPC.Handle", &req, &res)
}