package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/support-sphere/support-sphere/internal/adapters/ticket/handler/grpc/pb"
	repo "github.com/support-sphere/support-sphere/internal/adapters/ticket/repository"
	"github.com/support-sphere/support-sphere/internal/core/ticket/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TicketServer struct {
	pb.UnimplementedTicketServiceServer
	TicketService *services.TicketService
}

// NewTicketServer creates a new TicketServer.

func NewTicketServer(db *sql.DB) *TicketServer {
	repo := repo.NewTicketRepository(db)
	return &TicketServer{
		TicketService: services.NewTicketService(repo),
	}
}

func (s *TicketServer) GetTicketByID(ctx context.Context, req *pb.GetTicketByIDRequest) (*pb.GetTicketByIDResponse, error) {
	t, err := s.TicketService.GetTicketByID(int(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Ticket not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get ticket: %v", err)
	}

	return &pb.GetTicketByIDResponse{
		Ticket: &pb.Ticket{
			Id:          int64(t.ID),
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt.Format(time.RFC3339),
			CreatedBy:   t.CreatedBy,
		},
	}, nil
}

func (s *TicketServer) GetTickets(ctx context.Context, req *pb.GetTicketsRequest) (*pb.GetTicketsResponse, error) {
	tickets, err := s.TicketService.GetTickets(int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get tickets: %v", err)
	}

	total, err := s.TicketService.GetTotalTickets()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get total tickets: %v", err)
	}

	var pbTickets []*pb.Ticket
	for _, t := range tickets {
		pbTickets = append(pbTickets, &pb.Ticket{
			Id:          int64(t.ID),
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt.Format(time.RFC3339),
			CreatedBy:   t.CreatedBy,
		})
	}

	return &pb.GetTicketsResponse{
		Tickets: pbTickets,
		Total:   int32(total),
	}, nil
}

func (s *TicketServer) CreateTicket(ctx context.Context, req *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	id, err := s.TicketService.CreateTicket(req.Title, req.Description, req.Status, req.CreatedBy)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create ticket: %v", err)
	}

	return &pb.CreateTicketResponse{
		Ticket: &pb.Ticket{
			Id:          int64(id.ID),
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
			CreatedBy:   req.CreatedBy,
		},
	}, nil
}
