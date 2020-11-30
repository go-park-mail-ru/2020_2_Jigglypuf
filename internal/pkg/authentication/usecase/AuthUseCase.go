package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	sanitizer      *bluemonday.Policy
	validator      *validator.Validate
	repository     interfaces.AuthRepository
	profileService profileService.ProfileServiceClient
	salt           string
}

func NewUserUseCase(repository interfaces.AuthRepository, profileService profileService.ProfileServiceClient, salt string) *UserUseCase {
	return &UserUseCase{
		sanitizer:      bluemonday.UGCPolicy(),
		validator:      validator.New(),
		repository:     repository,
		profileService: profileService,
		salt:           salt,
	}
}

func createHashPassword(password, salt string) (string, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), 7)
	return string(hashedPassword), err == nil
}

func compareHashAndPassword(password, hash, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

func (t *UserUseCase) validateInput(input interface{}) error {
	return t.validator.Struct(input)
}

func (t *UserUseCase) SignUp(input *models.RegistrationInput) (uint64, error) {
	// input validation
	if input.Login == "" || input.Password == "" || input.Name == "" || input.Surname == "" {
		return 0, models.ErrFooIncorrectInputInfo
	}
	utils.SanitizeInput(t.sanitizer, &input.Login, &input.Password, &input.Name, &input.Surname)
	validationErr := t.validateInput(input)
	if validationErr != nil {
		return 0, models.ErrFooIncorrectInputInfo
	}
	// creating user credentials
	hashPassword, ok := createHashPassword(input.Password, t.salt)
	if !ok {
		return 0, models.ErrFooInternalServerError
	}
	user := models.User{
		Login:    input.Login,
		Password: hashPassword,
	}
	err := t.repository.CreateUser(&user)
	if err != nil {
		return 0, err
	}

	// creating profile
	_, profileErr := t.profileService.CreateProfile(context.Background(),&profileService.CreateProfileRequest{
		Profile: &profileService.Profile{
			Name: input.Name,
			Surname: input.Surname,
			UserID: user.ID,
			AvatarPath: profile.NoAvatarImage,
		},
	})
	if profileErr != nil {
		return 0, profileErr
	}

	return user.ID, nil
}

func (t *UserUseCase) SignIn(input *models.AuthInput) (uint64, error) {
	if input.Login == "" || input.Password == "" {
		return 0, models.ErrFooIncorrectInputInfo
	}
	utils.SanitizeInput(t.sanitizer, &input.Login, &input.Password)
	validationErr := t.validateInput(input)
	if validationErr != nil {
		return 0, models.ErrFooIncorrectInputInfo
	}

	user, err := t.repository.GetUser(input.Login)
	if err != nil {
		return 0, err
	}

	isAuth := compareHashAndPassword(input.Password, user.Password, t.salt)
	if !isAuth {
		return 0, models.ErrFooIncorrectInputInfo
	}
	return user.ID, nil
}


func (t *UserUseCase) GetUserByID(userID uint64) (*models.User, error){
	// TODO tracing
	return t.repository.GetUserByID(userID)
}