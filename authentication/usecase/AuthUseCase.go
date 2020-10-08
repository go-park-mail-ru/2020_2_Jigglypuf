package usecase

import (
	"authentication"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"models"
	"net/http"
	"time"
)

type UserUseCase struct{
	memConn authentication.AuthRepository
	salt string
}

type IncorrectInputError struct{}

func (t IncorrectInputError) Error() string{
	return "Incorrect Login or Password!"
}

func NewUserUseCase(dbConn authentication.AuthRepository, Salt string) *UserUseCase{
	return &UserUseCase{
		memConn: dbConn,
		salt: Salt,
	}
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
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
		Name: "session_id",
		Value: RandStringRunes(32),
		Expires: time.Now().Add(96*time.Hour),
		Path: "/",
	}
}

func (t *UserUseCase) SignUp(input *models.RegistrationInput)(*http.Cookie,error){
	username := input.Login
	password := input.Password
	if username == "" || password == ""{
		return new(http.Cookie),IncorrectInputError{}
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
		return new(http.Cookie),IncorrectInputError{}
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

func (t *UserUseCase) SignOut() error{
	return nil
}