//go:generate mockgen -source ProfileManager.go -destination mock/ProfileMan_mock.go -package mock
package profile

import (
	"context"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"google.golang.org/grpc"
)

type ManagerInterface interface{
	CreateProfile(ctx context.Context, in *profileService.CreateProfileRequest, opts ...grpc.CallOption) (*profileService.Nil, error)
	GetProfile(ctx context.Context, in *profileService.GetProfileRequest, opts ...grpc.CallOption) (*profileService.Profile, error)
	GetProfileByID(ctx context.Context, in *profileService.GetProfileByUserIDRequest, opts ...grpc.CallOption) (*profileService.Profile, error)
	UpdateProfile(ctx context.Context, in *profileService.UpdateProfileRequest, opts ...grpc.CallOption) (*profileService.Nil, error)
}