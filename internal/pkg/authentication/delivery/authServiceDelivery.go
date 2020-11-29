package delivery

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces"
	auth "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen/"
)

type AuthServiceDelivery struct{
	useCase interfaces.UserUseCase
}

func NewAuthServiceDelivery(useCase interfaces.UserUseCase) *AuthServiceDelivery{
	return &AuthServiceDelivery{
		useCase: useCase,
	}
}


func (t *AuthServiceDelivery) SignIn(ctx context.Context, request *auth.SignInRequest)(*auth.SignInResponse, error){

}