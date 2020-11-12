package usecase

import(
	"errors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

type TestingCookieStruct struct {
	repositoryMock   *mock.MockRepository
	goMockController *gomock.Controller
	mainUseCase      *CookieUseCase
}

var(
	testingStruct *TestingCookieStruct = nil
)

func setUp(t *testing.T){
	testingStruct = new(TestingCookieStruct)
	testingStruct.goMockController = gomock.NewController(t)
	testingStruct.repositoryMock = mock.NewMockRepository(testingStruct.goMockController)
	testingStruct.mainUseCase = NewCookieUseCase(testingStruct.repositoryMock)
}

func tearDown(){
	testingStruct.goMockController.Finish()
}


func TestCheckCookieSuccess(t *testing.T){
	setUp(t)
	returnValue := models.DBResponse {
		1,
		"some cookie",
		1,
		http.Cookie{Expires: time.Now().Add(time.Hour)},
	}
	testingStruct.repositoryMock.EXPECT().GetCookie(gomock.Any()).Return(&returnValue, nil)
	val, ok := testingStruct.mainUseCase.CheckCookie(new(http.Cookie))
	assert.Equal(t, ok, true)
	assert.Equal(t,val,returnValue.UserID)
	tearDown()
}

func TestCheckCookieTimeAfter(t *testing.T){
	setUp(t)
	returnValue := models.DBResponse {
		1,
		"some cookie",
		1,
		http.Cookie{Expires: time.Now().Add(-time.Hour)},
	}
	testingStruct.repositoryMock.EXPECT().GetCookie(gomock.Any()).Return(&returnValue, nil)
	val, ok := testingStruct.mainUseCase.CheckCookie(new(http.Cookie))
	assert.Equal(t, ok, false)
	assert.Equal(t,val,uint64(0))
	tearDown()
}

func TestCheckCookieErr(t *testing.T){
	setUp(t)
	testingStruct.repositoryMock.EXPECT().GetCookie(gomock.Any()).Return(nil, errors.New("some error"))
	val, ok := testingStruct.mainUseCase.CheckCookie(new(http.Cookie))
	assert.Equal(t, ok, false)
	assert.Equal(t,val,uint64(0))
	tearDown()
}