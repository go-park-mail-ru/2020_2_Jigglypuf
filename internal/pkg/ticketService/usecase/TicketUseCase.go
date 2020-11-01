package usecase

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/hallService"
	"backend/internal/pkg/models"
	"backend/internal/pkg/ticketService"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"strconv"
)

type TicketUseCase struct{
	validator *validator.Validate
	sanitizer bluemonday.Policy
	repository ticketService.Repository
	userRepository authentication.AuthRepository
	hallRepository hallService.Repository
}

func NewTicketUseCase(repository ticketService.Repository, authRepository authentication.AuthRepository, hallRepository hallService.Repository) *TicketUseCase {
	return &TicketUseCase{
		validator: validator.New(),
		repository: repository,
		sanitizer: *bluemonday.UGCPolicy(),
		userRepository: authRepository,
		hallRepository: hallRepository,
	}
}

func (t *TicketUseCase) GetUserTickets(userID uint64) (*[]models.Ticket, error){
	user, getUserErr := t.userRepository.GetUserByID(userID)
	if getUserErr != nil{
		return nil,getUserErr
	}
	return t.repository.GetUserTickets(user.Username)
}

func (t *TicketUseCase) GetSimpleTicket(userID uint64, ticketID string)(*models.Ticket, error){
	castedTicketID, castErr := strconv.Atoi(ticketID)
	if castErr != nil {
		return nil, castErr
	}
	user,getUserErr := t.userRepository.GetUserByID(userID)
	if getUserErr != nil{
		return nil, getUserErr
	}
	return t.repository.GetSimpleTicket(uint64(castedTicketID),user.Username)
}

func (t *TicketUseCase) GetHallScheduleTickets(scheduleID string)(*[]models.TicketPlace,error){
	castedScheduleID, castErr := strconv.Atoi(scheduleID)
	if castErr != nil{
		return nil,models.ErrFooCastErr
	}
	return t.repository.GetHallTickets(uint64(castedScheduleID))
}

func (t *TicketUseCase) BuyTicket(ticket *models.TicketInput, userID uint64) error{
	// TODO check if place available in hall
	if ticket.Username == ""{
		user, getUserErr := t.userRepository.GetUserByID(userID)
		if getUserErr != nil{
			return models.ErrFooNoAuthorization
		}
		ticket.Username = user.Username
	}
	availability, avErr := t.hallRepository.CheckAvailability(ticket.HallID, &ticket.PlaceField)
	if avErr != nil || !availability{
		return models.ErrFooPlaceAlreadyBusy
	}

	validationError := t.validator.Struct(ticket)
	if validationError != nil{
		return validationError
	}

	ticket.Username = t.sanitizer.Sanitize(ticket.Username)
	return t.repository.CreateTicket(ticket)
}
