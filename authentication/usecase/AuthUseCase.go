package usecase

import (
	"authentication"
	"fmt"
	"math/rand"
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
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func createHashPassword(password, salt string) string{
	return password + salt
}

func createUserCookie() http.Cookie{
	return http.Cookie{
		Name: "session_id",
		Value: RandStringRunes(32),
		Expires: time.Now().Add(24*time.Hour),
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
		user.Cookie = createUserCookie()
	}
	fmt.Println("kek")

	return &user.Cookie, err
}

func (t *UserUseCase) SignOut() error{
	return nil
}