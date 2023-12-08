package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/94peter/address/model"
	"github.com/94peter/address/mygrpc/pb"
	"github.com/94peter/address/mygrpc/server"

	"github.com/94peter/sterna/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
)

var (
	v = flag.Bool("v", false, "version")

	Version   = "1.0.0"
	BuildTime = "2000-01-01T00:00:00+0800"
)

func main() {
	flag.Parse()

	if *v {
		fmt.Println("Version: " + Version)
		fmt.Println("Build Time: " + BuildTime)
		return
	}

	if util.FileExists(".env") {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("No .env file")
		}
	}
	model.InitData(os.Getenv("DATA_FILE"))

	gRPC()
}

func gRPC() {
	port := ":" + os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterAddressParserServiceServer(s, server.NewAddressParserServer())
	log.Printf("core gRPC server is running [%s].", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
