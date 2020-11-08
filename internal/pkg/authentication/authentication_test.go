package authentication

import (
	"backend/internal/pkg/authentication/delivery"
	"database/sql"
	"github.com/golang/mock/gomock"
	"testing"
)

type TestingAuthenticationStruct struct{
	Handler delivery.UserHandler
	UseCase UserUseCase
	Repository AuthRepository
	DBConn *sql.DB
	GoMockController *gomock.Controller
}
var(
	TestingStruct = new(TestingAuthenticationStruct)
)

func TestSignUp(t *testing.T){
	TestingStruct.GoMockController = gomock.NewController(t)
}