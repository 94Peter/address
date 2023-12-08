package client

import (
	"address/mygrpc/pb"

	"github.com/94peter/sterna/mygrpc"
	"golang.org/x/net/context"
)

type AddressParserClient interface {
	Parser(address string) (*pb.ParserResponse, error)
}

func NewAddressParserClient(address string) AddressParserClient {
	return &addressParserImpl{
		address: address,
	}
}

type addressParserImpl struct {
	address string
}

func (grpc *addressParserImpl) Parser(address string) (*pb.ParserResponse, error) {
	var err error

	grpcClt, err := mygrpc.New(grpc.address)
	if err != nil {
		return nil, err
	}
	defer grpcClt.Close()
	clt := pb.NewAddressParserServiceClient(grpcClt)
	resp, err := clt.SimpleParser(context.Background(), &pb.ParserRequest{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
