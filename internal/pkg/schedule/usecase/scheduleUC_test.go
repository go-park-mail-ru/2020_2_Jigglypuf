package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

type ScheduleUCTesting struct{
	useCase *ScheduleUseCase
	DBMock *mock.MockTimeTableRepository
	controller *gomock.Controller
}
var(
	test *ScheduleUCTesting = nil
)

func setup(t *testing.T){
	test = new(ScheduleUCTesting)
	test.controller = gomock.NewController(t)
	test.DBMock = mock.NewMockTimeTableRepository(test.controller)
	test.useCase = NewTimeTableUseCase(test.DBMock)
}


func teardown(){
	test.controller.Finish()
}

func TestGetMovieScheduleSuccess( t *testing.T){
	setup(t)

	const(
		movieIDs = "1"
		cinemaIDs = "1"
		movieID = uint64(1)
		cinemaID = uint64(1)
		date = "2020-12-12"
	)
	test.DBMock.EXPECT().GetMovieCinemaSchedule(movieID, cinemaID, date).Return(nil, nil)
	_, _ = test.useCase.GetMovieSchedule(movieIDs, cinemaIDs, date)

	teardown()
}


func TestGetScheduleSuccess( t *testing.T){
	setup(t)

	const(
		movieIDs = "1"
		cinemaIDs = "1asd"
		movieID = uint64(1)
		cinemaID = uint64(1)
		date = "2020-12-12"
	)
	test.DBMock.EXPECT().GetMovieSchedule(movieID, date).Return(nil, nil)
	_, _ = test.useCase.GetMovieSchedule(movieIDs, cinemaIDs, date)

	teardown()
}

func TestGetScheduleFailMovie( t *testing.T){
	setup(t)

	const(
		movieIDs = "1asdsad"
		cinemaIDs = "1asd"
		movieID = uint64(1)
		cinemaID = uint64(1)
		date = "2020-12-12"
	)
	//test.DBMock.EXPECT().GetMovieSchedule(movieID, date).Return(nil, nil)
	_, _ = test.useCase.GetMovieSchedule(movieIDs, cinemaIDs, date)

	teardown()
}

func TestGetScheduleFailDate( t *testing.T){
	setup(t)

	const(
		movieIDs = "1"
		cinemaIDs = "1asd"
		movieID = uint64(1)
		cinemaID = uint64(1)
		date = "asd"
	)
	test.DBMock.EXPECT().GetMovieSchedule(movieID, gomock.Any()).Return(nil, nil)
	_, _ = test.useCase.GetMovieSchedule(movieIDs, cinemaIDs, date)

	teardown()
}

func TestGetSimpleSchedule( t *testing.T){
	setup(t)

	const(
		movieIDs = "1"
		cinemaIDs = "1asd"
		movieID = uint64(1)
		cinemaID = uint64(1)
		date = "asd"
	)
	test.DBMock.EXPECT().GetSchedule( gomock.Any()).Return(nil, nil)
	_, _ = test.useCase.GetSchedule(movieIDs)

	teardown()
}

func TestGetSimpleScheduleFail( t *testing.T){
	setup(t)

	const(
		movieIDs = "1"
		cinemaIDs = "1asd"
		movieID = uint64(1)
		cinemaID = uint64(1)
		date = "asd"
	)
	//test.DBMock.EXPECT().GetSchedule( gomock.Any()).Return(nil, nil)
	_, _ = test.useCase.GetSchedule(cinemaIDs)

	teardown()
}