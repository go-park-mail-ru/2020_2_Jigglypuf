//go:generate mockgen -source AuthManager.go -destination mock/AuthenticationManager_mock.go -package mock
package interfaces

import (
	"context"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"google.golang.org/grpc"
)

type AuthManager interface {
	GetUserByID(ctx context.Context, in *authService.GetUserByIDRequest, opts ...grpc.CallOption) (*authService.UserResponse, error)
	SignUp(ctx context.Context, in *authService.SignUpRequest, opts ...grpc.CallOption) (*authService.Response, error)
	SignIn(ctx context.Context, in *authService.SignInRequest, opts ...grpc.CallOption) (*authService.Response, error)
}
