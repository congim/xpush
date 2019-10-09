package cluster

import (
	"context"
	"net"

	"github.com/congim/xpush/broker/internal/proto/peer"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) OnMessage(c context.Context, in *peer.Message) (*peer.Reply, error) {
	return &peer.Reply{Message: "Hello " + in.Name}, nil
}

func initServer(addr string, logger *zap.Logger) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("listen", zap.Error(err))
		return err
	}

	defer func() {
		_ = lis.Close()
	}()

	s := grpc.NewServer()
	peer.RegisterPeerServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to server", zap.Error(err))
		return err
	}

	defer func() {
		s.Stop()
	}()

	return nil
}
