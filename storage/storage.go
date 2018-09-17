package storage

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/b4t3ou/csv-ingester/config"
	pb "github.com/b4t3ou/csv-ingester/proto"
)

// Server representing the storage server
type Server struct {
	config *config.Config
	db     DBInterface
}

// NewServer returning with a new server
func NewServer(cfg *config.Config, db DBInterface) *Server {
	return &Server{
		config: cfg,
		db:     db,
	}
}

// ListenAndServe is running the server
func (s *Server) ListenAndServe() {
	s.runGRPServer()
}

func (s *Server) runGRPServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.config.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	server := grpc.NewServer()
	pb.RegisterStorageServer(server, s)
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

// Save is saving a list of records
func (s *Server) Save(ctx context.Context, req *pb.StorageRequest) (*pb.StorageReply, error) {
	for _, record := range req.Items {
		err := s.db.Save(Record{
			ID:           record.Id,
			Email:        record.Email,
			MobileNumber: record.MobileNumber,
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.StorageReply{}, nil
}
