package worker

import (
	"context"
	"goDistributedSystem/pkg/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(masterAddr string) error {
	conn, err := grpc.NewClient(
		masterAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewNodeServiceClient(conn)

	statusResp, err := client.ReportStatus(context.Background(), &pb.Request{
		Data: "worker online",
	})

	if err != nil {
		return err
	}

	log.Printf("master status response: %s", statusResp.Data)

	stream, err := client.AssignTask(context.Background(), &pb.Request{
		Data: "ready for work",
	})

	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != io.EOF {
			log.Println("task stream closed")
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("recieved task: %s", resp.Data)

		time.Sleep(2 * time.Second)

		log.Printf("completed task: %s", resp.Data)
	}

}
