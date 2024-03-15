package transactor

import (
	"context"

	pb "github.com/fair-n-square-co/apis/gen/pkg/fairnsquare/transaction_manager/v1alpha1"
)

type transactionsServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *transactionsServer) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}

func NewTransactionsServer() *transactionsServer {
	return &transactionsServer{}
}
