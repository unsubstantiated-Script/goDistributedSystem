package worker

import "google.golang.org/grpc"

func Run(masterAddr string) error {
	conn, err := grpc.Dial(masterAddr, grpc.WithInsecure())
}
