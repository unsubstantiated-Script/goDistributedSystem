package master

import (
	"context"
	"goDistributedSystem/pkg/pb"
	"log"
)

type NodeServer struct {
	pb.UnimplementedNodeServiceServer
	CmdChannel chan string
}

func NewNodeServer() *NodeServer {
	return &NodeServer{
		CmdChannel: make(chan string, 100),
	}
}

func (s *NodeServer) ReportStatus(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("worker status check: %s", req.Data)
	return &pb.Response{Data: "ok"}, nil
}

func (s *NodeServer) AssignTask(req *pb.Request, stream pb.NodeService_AssignTaskServer) error {
	log.Printf("worker connected for tasks: %s", req.Data)

	for cmd := range s.CmdChannel {
		log.Printf("sending task to worker: %s", cmd)

		if err := stream.Send(&pb.Response{Data: cmd}); err != nil {
			return err
		}
	}

	return nil
}
