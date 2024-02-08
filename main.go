package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"task-rpc-server/service"
)


func main() {
   rpc.RegisterName("MyService", &service.TaskService{})

   listener, err := net.Listen("tcp", ":1234")
   if err != nil {
	   fmt.Println("Error starting RPC server:", err)
	   return
   }
   defer listener.Close()

   fmt.Println("RPC server started, listening on port 1234...")

   for {
	   conn, err := listener.Accept()
	   if err != nil {
		   fmt.Println("Error accepting connection:", err)
		   continue
	   }
	   go jsonrpc.ServeConn(conn)
   }
}