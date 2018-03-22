package main

import (
	"google.golang.org/grpc"
)

// NewGRPCConn is a helper wrapper around grpc.Dial.
func NewGRPCConn(
	address string,
) (*grpc.ClientConn, error) {
		return grpc.Dial(address,
			grpc.WithInsecure())

}