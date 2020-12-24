package usecase

import (
	"context"
	"fmt"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/globalconfig"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils/mailer"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type TicketUseCase struct {
	validator          *validator.Validate
	sanitizer          bluemonday.Policy
	repository         ticketservice.Repository
	AuthServiceClient  authService.AuthenticationServiceClient
	hallRepository     hallservice.Repository
	scheduleRepository schedule.TimeTableRepository
	mailer             *mailer.Mailer
}

func NewTicketUseCase(repository ticketservice.Repository, authRepository authService.AuthenticationServiceClient, hallRepository hallservice.Repository, scheduleRepository schedule.TimeTableRepository, email, password, host string, port int) *TicketUseCase {
	if _, err := os.Stat(globalconfig.QRCodesPath); os.IsNotExist(err) {
		err := os.MkdirAll(globalconfig.QRCodesPath, os.ModePerm)
		if err != nil {
			log.Println("ERROR FILE", err)
			return nil
		}
	}
	return &TicketUseCase{
		validator:          validator.New(),
		repository:         repository,
		sanitizer:          *bluemonday.UGCPolicy(),
		AuthServiceClient:  authRepository,
		hallRepository:     hallRepository,
		scheduleRepository: scheduleRepository,
		mailer:             mailer.NewMailer(email, password, host, port),
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

func (t *TicketUseCase) sendQrTicketMail(unique string, to string) error {
	filename := filepath.Join(globalconfig.QRCodesPath, unique) + ".png"
	f, err := os.Create(filename)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		fmt.Println("QR FIL", err)
		return err
	}
	err = qrcode.WriteFile(globalconfig.QRURL+unique+"/", qrcode.Medium, 256, filename)
	if err != nil {
		fmt.Println("QRE", err)
		return err
	}
	Subject := "Cinemascope ticket"
	BodyType := "text/html"
	Body := `<h1>QR CODE FOR TICKET</h1><img src="cid:image.png" alt="My image" />`
	return t.mailer.SendFiledMail(filename, to, Subject, BodyType, Body)
}

func (t *TicketUseCase) sendMails(ticket *models.TicketInput) {
	for _, val := range ticket.Transaction {
		_ = t.sendQrTicketMail(val, ticket.Login)
	}
}

func (t *TicketUseCase) GetTicketByTransaction(transaction string) (*models.TicketInfo, error) {
	if transaction == "" {
		return nil, models.ErrFooIncorrectInputInfo
	}
	return t.repository.GetTicketByTransaction(transaction)
}

func (t *TicketUseCase) signTicket(ticket *models.TicketInput, hallStructure *models.CinemaHall) bool {
	for _, val := range ticket.PlaceField {
		exists := false
		ticket.Transaction = append(ticket.Transaction, models.RandStringRunes(32))
		for _, hall := range hallStructure.PlaceConfig.Levels {
			if val.Place == hall.Place && val.Row == hall.Row {
				exists = true
				break
			}
		}
		if !exists {
			return exists
		}
	}
	return true
}

func (t *TicketUseCase) BuyTicket(ticket *models.TicketInput, userID interface{}) error {
	if len(ticket.PlaceField) > ticketservice.MaxPlaceCollection {
		return models.ErrFooIncorrectInputInfo
	}
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
	hallStructure, hallErr := t.hallRepository.GetHallStructure(HallID)
	if hallErr != nil {
		return models.ErrFooIncorrectInputInfo
	}

	ok := t.signTicket(ticket, hallStructure)
	if !ok {
		return models.ErrFooPlaceDoesntExists
	}

	validationError := t.validator.Struct(ticket)
	if validationError != nil {
		return validationError
	}

	ticket.Login = t.sanitizer.Sanitize(ticket.Login)
	if ticket.Login == "" {
		return models.ErrFooIncorrectInputInfo
	}
	err := t.repository.CreateTicket(ticket)
	if err != nil {
		return err
	}

	t.sendMails(ticket)
	return nil
}
