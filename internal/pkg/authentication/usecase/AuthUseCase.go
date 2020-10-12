package usecase

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/cookie"
	"backend/internal/pkg/models"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"net/http"
	"time"
)
var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type IncorrectInputError struct{}

func (t IncorrectInputError) Error() string{
	return "Incorrect Login or Password!"
}

type UserUseCase struct{
	memConn authentication.AuthRepository
	salt    string
}

func NewUserUseCase(dbConn authentication.AuthRepository, Salt string) *UserUseCase {
	return &UserUseCase{
		memConn: dbConn,
		salt: Salt,
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		b[i] = letterRunes[randInt.Int64()]
	}
	return string(b)
}

func createHashPassword(password, salt string) string{
	reqString := password + salt
	decoder := sha256.New()
	decoder.Write([]byte(reqString))
	resultString := hex.EncodeToString(decoder.Sum(nil))
	return resultString
}

func createUserCookie() http.Cookie{
	return http.Cookie{
		Name:    cookie.SessionCookieName,
		Value:   randStringRunes(32),
		Expires: time.Now().Add(96*time.Hour),
		Path:    "/",
		SameSite: http.SameSiteNoneMode,
		Secure: true,
	}
}

func (t *UserUseCase) SignUp(input *models.RegistrationInput)(*http.Cookie,error){
	username := input.Login
	password := input.Password

	if username == "" || password == ""{
		return new(http.Cookie), IncorrectInputError{}
	}

	hashPassword := createHashPassword(password, t.salt)
	cookieValue := createUserCookie()

	user := models.User{
		Username: username,
		Password: hashPassword,
		Cookie: cookieValue,
	}

	err := t.memConn.CreateUser(&user)

	return &cookieValue,err
}

func (t *UserUseCase) SignIn (input *models.AuthInput)(*http.Cookie,error){
	username := input.Login
	password := input.Password
	if username == "" || password == ""{
		return new(http.Cookie), IncorrectInputError{}
	}

	hashPassword := createHashPassword(password, t.salt)

	user, err := t.memConn.GetUser(username,hashPassword)
	if err != nil{
		return &http.Cookie{}, err
	}

	if time.Now().After(user.Cookie.Expires){
		cookieValue := createUserCookie()
		user.Cookie = cookieValue
		t.memConn.SetCookie(user,&cookieValue)
	}

	return &user.Cookie, err
}

func (t *UserUseCase) SignOut(cookie *http.Cookie) (*http.Cookie,error){
	cookie.Expires = time.Now().Add(-time.Hour)
	return cookie, nil
}