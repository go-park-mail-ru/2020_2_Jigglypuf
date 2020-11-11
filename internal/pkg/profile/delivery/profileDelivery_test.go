package delivery

import(
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/mock"
	"github.com/golang/mock/gomock"
	"testing"
)


type ProfileTesting struct{
	handler *ProfileHandler
	useCaseMock *mock.MockUseCase
	goMockController *gomock.Controller
}

var(
	testingStruct *ProfileTesting = nil
)

func setup(t *testing.T){

}