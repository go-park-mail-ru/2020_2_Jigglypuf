package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"testing"
)

type TestingCinemaUseCase struct {
	useCase          *CinemaUseCase
	DBMock           *mock.MockRepository
	goMockController *gomock.Controller
}

var (
	testingStruct *TestingCinemaUseCase = nil
)

func setup(t *testing.T) {
	testingStruct = new(TestingCinemaUseCase)

	testingStruct.goMockController = gomock.NewController(t)
	testingStruct.DBMock = mock.NewMockRepository(testingStruct.goMockController)
	testingStruct.useCase = NewCinemaUseCase(testingStruct.DBMock)
}

func tearDown() {
	testingStruct.goMockController.Finish()
}

func TestGetCinema(t *testing.T) {
	setup(t)
	testingStruct.DBMock.EXPECT().GetCinema(gomock.Any()).Return(nil, nil)
	_, _ = testingStruct.useCase.GetCinema(uint64(1))
	tearDown()
}

func TestCreateCinema(t *testing.T) {
	setup(t)

	testingStruct.DBMock.EXPECT().CreateCinema(gomock.Any()).Return(nil)
	_ = testingStruct.useCase.CreateCinema(new(models.Cinema))

	tearDown()
}

func TestGetCinemaList(t *testing.T) {
	setup(t)

	testingStruct.DBMock.EXPECT().GetCinemaList(gomock.Any(), gomock.Any()).Return(nil, nil)
	_, err := testingStruct.useCase.GetCinemaList(1, 1)
	if err != nil {
		t.Fatalf("INCORRECT TEST GetCinemaList")
	}
	tearDown()
}

func TestGetCinemaListErr(t *testing.T) {
	setup(t)

	_, err := testingStruct.useCase.GetCinemaList(0, 0)
	if err == nil {
		t.Fatalf("INCORRECT TEST GetCinemaListErr")
	}
	tearDown()
}
