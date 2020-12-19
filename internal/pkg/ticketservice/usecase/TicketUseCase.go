package usecase

import (
	"context"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"strconv"
)

type TicketUseCase struct {
	validator          *validator.Validate
	sanitizer          bluemonday.Policy
	repository         ticketservice.Repository
	AuthServiceClient  authService.AuthenticationServiceClient
	hallRepository     hallservice.Repository
	scheduleRepository schedule.TimeTableRepository
}

func NewTicketUseCase(repository ticketservice.Repository, authRepository authService.AuthenticationServiceClient, hallRepository hallservice.Repository, scheduleRepository schedule.TimeTableRepository) *TicketUseCase {
	return &TicketUseCase{
		validator:          validator.New(),
		repository:         repository,
		sanitizer:          *bluemonday.UGCPolicy(),
		AuthServiceClient:  authRepository,
		hallRepository:     hallRepository,
		scheduleRepository: scheduleRepository,
	}
}

func (t *TicketUseCase) GetUserTickets(userID uint64) (*[]models.Ticket, error) {
	user, getUserErr := t.AuthServiceClient.GetUserByID(context.Background(), &authService.GetUserByIDRequest{UserID: userID})
	if getUserErr != nil {
		return nil, getUserErr
	}
	return t.repository.GetUserTickets(user.User.Login)
}

func (t *TicketUseCase) GetSimpleTicket(userID uint64, ticketID string) (*models.Ticket, error) {
	castedTicketID, castErr := strconv.Atoi(ticketID)
	if castErr != nil {
		return nil, castErr
	}
	user, getUserErr := t.AuthServiceClient.GetUserByID(context.Background(), &authService.GetUserByIDRequest{UserID: userID})
	if getUserErr != nil {
		return nil, getUserErr
	}
	return t.repository.GetSimpleTicket(uint64(castedTicketID), user.User.Login)
}

func (t *TicketUseCase) GetHallScheduleTickets(scheduleID string) (*[]models.TicketPlace, error) {
	castedScheduleID, castErr := strconv.Atoi(scheduleID)
	if castErr != nil {
		return nil, models.ErrFooCastErr
	}
	return t.repository.GetHallTickets(uint64(castedScheduleID))
}

func (t *TicketUseCase) BuyTicket(ticket *models.TicketInput, userID interface{}) error {
	if ticket.Login == "" {
		if userID == nil {
			return models.ErrFooNoAuthorization
		}
		if _, ok := userID.(uint64); !ok {
			return models.ErrFooNoAuthorization
		}
		user, getUserErr := t.AuthServiceClient.GetUserByID(context.Background(), &authService.GetUserByIDRequest{UserID: userID.(uint64)})
		if getUserErr != nil {
			return models.ErrFooNoAuthorization
		}
		ticket.Login = user.User.Login
	}

	HallID, HallErr := t.scheduleRepository.GetScheduleHallID(ticket.ScheduleID)
	if HallErr != nil {
		return models.ErrFooIncorrectInputInfo
	}

	for _, val := range ticket.PlaceField {
		availability, avErr := t.hallRepository.CheckAvailability(HallID, &val)
		if avErr != nil || !availability {
			return models.ErrFooPlaceAlreadyBusy
		}
	}

	validationError := t.validator.Struct(ticket)
	if validationError != nil {
		return validationError
	}

	ticket.Login = t.sanitizer.Sanitize(ticket.Login)
	return t.repository.CreateTicket(ticket)
}
