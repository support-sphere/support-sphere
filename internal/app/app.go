package app

import (
	"github.com/support-sphere/support-sphere/internal/core/ticket"
	"github.com/support-sphere/support-sphere/internal/core/user"
)

type App struct {
	userService   user.Service
	ticketService ticket.Service
}

func NewApp(userService user.Service, ticketService ticket.Service) *App {
	return &App{
		userService:   userService,
		ticketService: ticketService,
	}
}

func (a *App) UserService() *user.Service {
	return &a.userService
}

func (a *App) TicketService() *ticket.Service {
	return &a.ticketService
}
