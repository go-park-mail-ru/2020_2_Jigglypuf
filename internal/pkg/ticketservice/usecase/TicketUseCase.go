package usecase

import (
	"backend/internal/pkg/authentication/interfaces"
	"backend/internal/pkg/hallservice"
	"backend/internal/pkg/models"
	"backend/internal/pkg/schedule"
	"backend/internal/pkg/ticketservice"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"strconv"
)

type TicketUseCase struct {
	validator      *validator.Validate
	sanitizer      bluemonday.Policy
	repository     ticketservice.Repository
	userRepository interfaces.AuthRepository
	hallRepository hallservice.Repository
	scheduleRepository schedule.TimeTableRepository
}

func NewTicketUseCase(repository ticketservice.Repository, authRepository interfaces.AuthRepository, hallRepository hallservice.Repository,scheduleRepository schedule.TimeTableRepository) *TicketUseCase {
	return &TicketUseCase{
		validator:      validator.New(),
		repository:     repository,
		sanitizer:      *bluemonday.UGCPolicy(),
		userRepository: authRepository,
		hallRepository: hallRepository,
		scheduleRepository: scheduleRepository,
	}
}

func (t *TicketUseCase) GetUserTickets(userID uint64) (*[]models.Ticket, error) {
	user, getUserErr := t.userRepository.GetUserByID(userID)
	if getUserErr != nil {
		return nil, getUserErr
	}
	return t.repository.GetUserTickets(user.Login)
}

func (t *TicketUseCase) GetSimpleTicket(userID uint64, ticketID string) (*models.Ticket, error) {
	castedTicketID, castErr := strconv.Atoi(ticketID)
	if castErr != nil {
		return nil, castErr
	}
	user, getUserErr := t.userRepository.GetUserByID(userID)
	if getUserErr != nil {
		return nil, getUserErr
	}
	return t.repository.GetSimpleTicket(uint64(castedTicketID), user.Login)
}

func (t *TicketUseCase) GetHallScheduleTickets(scheduleID string) (*[]models.TicketPlace, error) {
	castedScheduleID, castErr := strconv.Atoi(scheduleID)
	if castErr != nil {
		return nil, models.ErrFooCastErr
	}
	return t.repository.GetHallTickets(uint64(castedScheduleID))
}

func (t *TicketUseCase) BuyTicket(ticket *models.TicketInput, userID uint64) error {
	if ticket.Login == "" {
		user, getUserErr := t.userRepository.GetUserByID(userID)
		if getUserErr != nil {
			return models.ErrFooNoAuthorization
		}
		ticket.Login = user.Login
	}

	HallID, HallErr := t.scheduleRepository.GetScheduleHallID(ticket.ScheduleID)
	if HallErr != nil{
		return models.ErrFooIncorrectInputInfo
	}

	availability, avErr := t.hallRepository.CheckAvailability(HallID, &ticket.PlaceField)
	if avErr != nil || !availability {
		return models.ErrFooPlaceAlreadyBusy
	}

	validationError := t.validator.Struct(ticket)
	if validationError != nil {
		return validationError
	}

	ticket.Login = t.sanitizer.Sanitize(ticket.Login)
	return t.repository.CreateTicket(ticket)
}
