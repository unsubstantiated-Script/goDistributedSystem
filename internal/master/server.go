package master

import (
	"context"
	"goDistributedSystem/pkg/pb"
)

type NodeServer struct {
	pb.UnimplementedNodeServiceServer
	CmdChannel chan string
}

func NewNodeServer() *NodeServer {
	return &NodeServer{
		CmdChannel: make(chan string),
	}
}

func (s *NodeServer) ReportStatus(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Data: "ok"}, nil
}

func (s *NodeServer) AssignTask(req *pb.Request, stream pb.NodeService_AssignTaskServer) error {
	for cmd := range s.CmdChannel {
		if err := stream.Send(&pb.Response{Data: cmd}); err != nil {
			return err
		}
	}
	return nil
}
