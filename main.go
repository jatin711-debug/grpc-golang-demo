package main

import (
	"context"
	"log"
	"net"
	"github.com.jatin711-debug/grpc-demo/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoiceServer
}

func (s myInvoicerServer) Create(ctx context.Context,req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoiceServer(serverRegistrar, service)
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
