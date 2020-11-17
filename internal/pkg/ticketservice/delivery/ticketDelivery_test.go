package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/ticketservice/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TicketTesting struct {
	handler          *TicketDelivery
	useCaseMock      *mock.MockUseCase
	goMockController *gomock.Controller
}

var (
	testingStruct *TicketTesting = nil
)

func setup(t *testing.T) {
	testingStruct = new(TicketTesting)
	testingStruct.goMockController = gomock.NewController(t)

	testingStruct.useCaseMock = mock.NewMockUseCase(testingStruct.goMockController)
	testingStruct.handler = NewTicketDelivery(testingStruct.useCaseMock)
}

func teardown() {
	testingStruct.goMockController.Finish()
}

func TestBuyTicketSuccess(t *testing.T) {
	setup(t)

	ticketInput := new(models.TicketInput)
	ticketInput.Login = "somelogin"
	ticketBody, _ := json.Marshal(ticketInput)
	testReq := httptest.NewRequest(http.MethodPost, "/ticket/buy", strings.NewReader(string(ticketBody)))
	testingStruct.useCaseMock.EXPECT().BuyTicket(gomock.Any(), gomock.Any()).Return(nil)

	testRec := httptest.NewRecorder()
	testingStruct.handler.BuyTicket(testRec, testReq)
	assert.Equal(t, testRec.Code, http.StatusOK)
	teardown()
}

func TestGetUserTicketsSuccess(t *testing.T) {
	setup(t)

	testReq := httptest.NewRequest(http.MethodGet, "/ticket/user", nil)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(0))
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)

	ticketList := new([]models.Ticket)
	testingStruct.useCaseMock.EXPECT().GetUserTickets(gomock.Any()).Return(ticketList, nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.GetUserTickets(testRec, testReq.WithContext(ctx))
	assert.Equal(t, testRec.Code, http.StatusOK)
	teardown()
}

func TestGetUserSimpleTicketSuccess(t *testing.T) {
	setup(t)

	testReq := httptest.NewRequest(http.MethodGet, "/ticket/1", nil)
	mux.SetURLVars(testReq, map[string]string{
		ticketservice.TicketIDQuery: "1",
	})
	ctx := testReq.Context()
	ticketList := new(models.Ticket)
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(0))
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	testingStruct.useCaseMock.EXPECT().GetSimpleTicket(gomock.Any(), gomock.Any()).Return(ticketList, nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.GetUsersSimpleTicket(testRec, testReq.WithContext(ctx))
	assert.Equal(t, testRec.Code, http.StatusOK)

	teardown()
}

func TestGetHallScheduleTickets(t *testing.T) {
	setup(t)

	testReq := httptest.NewRequest(http.MethodGet, "/ticket/1", nil)
	mux.SetURLVars(testReq, map[string]string{
		ticketservice.ScheduleIDName: "1",
	})
	testingStruct.useCaseMock.EXPECT().GetHallScheduleTickets(gomock.Any()).Return(new([]models.TicketPlace), nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.GetHallScheduleTickets(testRec, testReq)
	assert.Equal(t, http.StatusOK, testRec.Code)

	teardown()
}

func TestTicketHandlersInvalidMethod(t *testing.T) {
	setup(t)

	var testCases = []struct {
		handler func(w http.ResponseWriter, r *http.Request)
		request *http.Request
	}{
		{
			testingStruct.handler.GetHallScheduleTickets,
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
		{
			testingStruct.handler.GetUsersSimpleTicket,
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
		{
			testingStruct.handler.GetUserTickets,
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
		{
			testingStruct.handler.BuyTicket,
			httptest.NewRequest(http.MethodGet, "/somepath/", nil),
		},
	}

	for _, val := range testCases {
		testRec := httptest.NewRecorder()
		val.handler(testRec, val.request)
		assert.Equal(t, http.StatusMethodNotAllowed, testRec.Code)
	}

	teardown()
}

func TestUnAuthorizedTicketHandlers(t *testing.T) {
	setup(t)
	ticketInput := new(models.TicketInput)
	ticketBody, _ := json.Marshal(ticketInput)
	var testCases = []struct {
		handler    func(w http.ResponseWriter, r *http.Request)
		request    *http.Request
		statusCode int
	}{
		{
			testingStruct.handler.BuyTicket,
			httptest.NewRequest(http.MethodPost, "/ticket/buy", strings.NewReader(string(ticketBody))),
			400,
		},
		{
			testingStruct.handler.GetUserTickets,
			httptest.NewRequest(http.MethodGet, "/ticket/user", nil),
			401,
		},
		{
			testingStruct.handler.GetUsersSimpleTicket,
			httptest.NewRequest(http.MethodGet, "/ticket/user", nil),
			401,
		},
	}

	for _, val := range testCases {
		testRec := httptest.NewRecorder()
		val.handler(testRec, val.request)
		assert.Equal(t, testRec.Code, val.statusCode)
	}

	teardown()
}
