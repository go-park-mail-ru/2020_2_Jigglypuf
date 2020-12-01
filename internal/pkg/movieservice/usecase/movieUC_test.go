package usecase
//
//import (
//	"errors"
//	mock2 "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/mock"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
//	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/mock"
//	"github.com/golang/mock/gomock"
//	"testing"
//)
//
//type MovieUseCaseTesting struct {
//	useCase          *MovieUseCase
//	DBMock           *mock.MockMovieRepository
//	UserMock         *mock2.MockAuthRepository
//	goMockController *gomock.Controller
//}
//
//var (
//	testingStruct *MovieUseCaseTesting = nil
//)
//
//func setup(t *testing.T) {
//	testingStruct = new(MovieUseCaseTesting)
//
//	testingStruct.goMockController = gomock.NewController(t)
//	testingStruct.DBMock = mock.NewMockMovieRepository(testingStruct.goMockController)
//	testingStruct.UserMock = mock2.NewMockAuthRepository(testingStruct.goMockController)
//	testingStruct.useCase = NewMovieUseCase(testingStruct.DBMock, testingStruct.UserMock)
//}
//
//func teardown() {
//	testingStruct.goMockController.Finish()
//}
//
//func TestGetMovieSuccess(t *testing.T) {
//	setup(t)
//	testingStruct.DBMock.EXPECT().GetRating(gomock.Any(), gomock.Any()).Return(int64(0), nil)
//	testingStruct.DBMock.EXPECT().GetMovie(gomock.Any()).Return(new(models.Movie), nil)
//
//	_, _ = testingStruct.useCase.GetMovie(uint64(1), true, uint64(1))
//	teardown()
//}
//
//func TestGetMovieFail(t *testing.T) {
//	setup(t)
//	testingStruct.DBMock.EXPECT().GetMovie(gomock.Any()).Return(nil, errors.New("some error"))
//
//	_, _ = testingStruct.useCase.GetMovie(uint64(1), true, uint64(1))
//	teardown()
//}
