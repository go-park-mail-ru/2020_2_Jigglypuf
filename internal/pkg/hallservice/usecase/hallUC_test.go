package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

type HallTesting struct {
	goMockController *gomock.Controller
	DBMock           *mock.MockRepository
	useCase          *HallUseCase
}

var (
	testingStruct *HallTesting = nil
)

func setup(t *testing.T) {
	testingStruct = new(HallTesting)

	testingStruct.goMockController = gomock.NewController(t)
	testingStruct.DBMock = mock.NewMockRepository(testingStruct.goMockController)
	testingStruct.useCase = NewHallUseCase(testingStruct.DBMock)
}

func tearDown() {
	testingStruct.goMockController.Finish()
}

func TestCheckAvailabilitySuccess(t *testing.T) {
	setup(t)

	testingStruct.DBMock.EXPECT().CheckAvailability(gomock.Any(), gomock.Any()).Return(true, nil)
	boolean, _ := testingStruct.useCase.CheckAvailability("1", new(models.TicketPlace))
	assert.Equal(t, boolean, true)
	tearDown()
}
func TestCheckAvailabilityNotSuccess(t *testing.T) {
	setup(t)

	// testingStruct.DBMock.EXPECT().CheckAvailability(gomock.Any(), gomock.Any()).Return(true, nil)
	boolean, _ := testingStruct.useCase.CheckAvailability("asd", new(models.TicketPlace))
	assert.Equal(t, boolean, false)
	tearDown()
}

func TestGetHallSuccess(t *testing.T) {
	setup(t)

	testingStruct.DBMock.EXPECT().GetHallStructure(gomock.Any()).Return(nil, nil)
	_, _ = testingStruct.useCase.GetHallStructure("1")
	tearDown()
}

func TestGetHallNotSuccess(t *testing.T) {
	setup(t)

	// testingStruct.DBMock.EXPECT().GetHallStructure(gomock.Any()).Return(nil, nil)
	_, _ = testingStruct.useCase.GetHallStructure("asd")
	tearDown()
}
