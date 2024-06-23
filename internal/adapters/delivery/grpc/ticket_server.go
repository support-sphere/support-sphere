package grpc

import (
	"context"

	"github.com/support-sphere/support-sphere/internal/core/domain"
	"github.com/support-sphere/support-sphere/internal/core/services"
	pb "github.com/support-sphere/support-sphere/proto"
)

type TicketServer struct {
	pb.UnimplementedTicketServiceServer
	service *services.TicketService
}

func NewTicketServer(service *services.TicketService) *TicketServer {
	return &TicketServer{service: service}
}

func (s *TicketServer) CreateTicket(ctx context.Context, req *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	ticket := domain.Ticket{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Status:      "New",
	}
	id, err := s.service.CreateTicket(ticket)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTicketResponse{Id: id}, nil
}
