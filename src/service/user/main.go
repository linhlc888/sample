package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	userpbgw "github.com/linhlc888/sample/src/proto/gateway/user"
	"github.com/linhlc888/sample/src/service/user/args"
	"github.com/linhlc888/sample/src/service/user/internal/rpc"
	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/gokums/core/log"
	"github.com/gokums/core/net/httpx"
)

func init() {
	flag.Parse()
}

func main() {

	httpLoc := fmt.Sprintf(":%s", *args.HTTPBind)
	rpcLoc := fmt.Sprintf(":%s", *args.RPCBind)
	//init grpc server
	srv := grpc.NewServer()
	userpbgw.RegisterUserServiceServer(srv, new(rpc.UserService))
	lis, err := net.Listen("tcp", rpcLoc)
	if err != nil {
		log.Fatalf("failed to listen rpc: %v", err)
	}
	defer lis.Close()

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve rpc: %v", err)
		}
	}()

	// Register the pb server to http gateway
	handler := httpx.NewHandler()

	muxSrv := runtime.NewServeMux(runtime.WithMarshalerOption("*", &runtime.JSONPb{OrigName: true}))
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	ctx := context.Background()
	err2 := userpbgw.RegisterUserServiceHandlerFromEndpoint(ctx, muxSrv, "localhost"+rpcLoc, opts)
	if err2 != nil {
		log.Fatalf("register endpoint failed: %v", err2)
	}
	handler.Handle("/v1/", muxSrv)
	go func() {
		if err := http.ListenAndServe(httpLoc, handler); err != nil {
			log.Fatalf("failed to listen and serve http: %v", err)
		}
	}()
	fmt.Printf("listening at %s port.", *args.HTTPBind)

	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGTERM)
	<-term
	srv.GracefulStop()
}
