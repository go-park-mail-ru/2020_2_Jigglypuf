package usecase

import (
	"backend/internal/pkg/authentication/interfaces"
	"backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"backend/internal/pkg/profile"
	"backend/internal/pkg/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type UserUseCase struct {
	sanitizer         *bluemonday.Policy
	validator         *validator.Validate
	repository        interfaces.AuthRepository
	cookieDBConn      cookie.Repository
	profileRepository profile.Repository
	salt              string
}

func NewUserUseCase(repository interfaces.AuthRepository, profileRepository profile.Repository, cookieConn cookie.Repository, salt string) *UserUseCase {
	return &UserUseCase{
		sanitizer:         bluemonday.UGCPolicy(),
		validator:         validator.New(),
		repository:        repository,
		cookieDBConn:      cookieConn,
		profileRepository: profileRepository,
		salt:              salt,
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

func createUserCookie() http.Cookie {
	return http.Cookie{
		Name:     cookie.SessionCookieName,
		Value:    models.RandStringRunes(32),
		Expires:  time.Now().Add(96 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	}
}

func (t *UserUseCase) validateInput(input interface{}) error {
	return t.validator.Struct(input)
}

func (t *UserUseCase) SignUp(input *models.RegistrationInput) (*http.Cookie, error) {
	// input validation
	if input.Login == "" || input.Password == "" || input.Name == "" || input.Surname == "" {
		return new(http.Cookie), models.ErrFooIncorrectInputInfo
	}
	utils.SanitizeInput(t.sanitizer, &input.Login, &input.Password, &input.Name, &input.Surname)
	validationErr := t.validateInput(input)
	if validationErr != nil {
		return nil, models.ErrFooIncorrectInputInfo
	}
	// creating user credentials
	hashPassword, ok := createHashPassword(input.Password, t.salt)
	if !ok {
		return nil, models.ErrFooInternalServerError
	}
	user := models.User{
		Login:    input.Login,
		Password: hashPassword,
	}
	err := t.repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	// creating profile
	prof := new(models.Profile)
	prof.Name = input.Name
	prof.Surname = input.Surname
	prof.UserCredentials = &user
	prof.AvatarPath = profile.NoAvatarImage
	profileErr := t.profileRepository.CreateProfile(prof)
	if profileErr != nil {
		return nil, profileErr
	}

	// creating cookie for user
	cookieValue := createUserCookie()
	cookieErr := t.cookieDBConn.SetCookie(&cookieValue, user.ID)
	if cookieErr != nil {
		return nil, cookieErr
	}
	return &cookieValue, nil
}

func (t *UserUseCase) SignIn(input *models.AuthInput) (*http.Cookie, error) {
	if input.Login == "" || input.Password == "" {
		return nil, models.ErrFooIncorrectInputInfo
	}
	utils.SanitizeInput(t.sanitizer, &input.Login, &input.Password)
	validationErr := t.validateInput(input)
	if validationErr != nil {
		return nil, models.ErrFooIncorrectInputInfo
	}

	user, err := t.repository.GetUser(input.Login)
	if err != nil {
		return nil, err
	}

	isAuth := compareHashAndPassword(input.Password, user.Password, t.salt)
	if !isAuth {
		return nil, models.ErrFooIncorrectInputInfo
	}

	cookieValue := createUserCookie()
	cookieErr := t.cookieDBConn.SetCookie(&cookieValue, user.ID)
	if cookieErr != nil {
		return nil, cookieErr
	}

	log.Println(cookieValue)
	return &cookieValue, nil
}

func (t *UserUseCase) SignOut(cookie *http.Cookie) (*http.Cookie, error) {
	cookie.Expires = time.Now().Add(-time.Hour)
	fmt.Print(cookie)
	cookieErr := t.cookieDBConn.RemoveCookie(cookie)
	return cookie, cookieErr
}
