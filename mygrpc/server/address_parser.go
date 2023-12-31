package server

import (
	"context"

	"github.com/94peter/address/model"
	"github.com/94peter/address/mygrpc/pb"
)

func NewAddressParserServer() pb.AddressParserServiceServer {
	return &AddressParserServer{}
}

type AddressParserServer struct {
	pb.UnimplementedAddressParserServiceServer
}

func (m *AddressParserServer) SimpleParser(ctx context.Context, req *pb.ParserRequest) (*pb.ParserResponse, error) {
	addressParser := model.NewSimpleAddressParser()
	myaddress := addressParser.Parser(req.Address)
	return &pb.ParserResponse{
		Code:    myaddress.ZipCode,
		State:   myaddress.State,
		City:    myaddress.City,
		Address: myaddress.Street,
	}, nil
}
