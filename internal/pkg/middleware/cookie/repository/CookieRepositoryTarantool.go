package repository

import (
	tarantoolConfig "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"encoding/json"
	"errors"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
)

type CookieTarantoolRepository struct {
	connectionDB *tarantool.Connection
}

func NewCookieTarantoolRepository(conn *tarantool.Connection) (*CookieTarantoolRepository, error) {
	return &CookieTarantoolRepository{
		connectionDB: conn,
	}, nil
}

func (t *CookieTarantoolRepository) GetCookie(cookie *http.Cookie) (*models.TarantoolResponse, error) {
	resp, DBErr := t.connectionDB.Eval("return check_session(...)", []interface{}{cookie.Value})
	if DBErr != nil {
		return nil, DBErr
	}
	if resp == nil {
		return nil, errors.New("incorrect session")
	}
	if resp.Data[0] == nil{
		return nil, errors.New("cookie not found")
	}

	data := resp.Data[0].([]interface{})
	tarantoolRes := new(models.TarantoolResponse)
	if data != nil && len(data) > 2{
		tarantoolRes.CookieValue = data[0].(string)
		rawUserID, ok := data[1].(uint64)
		if !ok{
			log.Println("cast err")
			return nil, errors.New("bad cookie")
		}
		tarantoolRes.UserID = rawUserID
		rawCookie := data[2].(string)
		translationErr := json.Unmarshal([]byte(rawCookie), &tarantoolRes.Cookie)
		if translationErr != nil{
			log.Println("translation err")
			return nil, errors.New("bad cookie")
		}
		log.Println(tarantoolRes)
		return tarantoolRes, nil
	}
	return nil, errors.New("cookie not found")
}

func (t *CookieTarantoolRepository) SetCookie(cookie *http.Cookie, userID uint64) error {
	stringCookie, cookieErr := json.Marshal(cookie)
	if cookieErr != nil {
		return errors.New("incorrect cookie structure")
	}
	resp, err := t.connectionDB.Eval("return create_session(...)", []interface{}{cookie.Value, string(stringCookie), userID})

	if err != nil {
		return err
	}
	log.Println(resp)

	return nil
}

func (t *CookieTarantoolRepository) RemoveCookie(cookie *http.Cookie) error {
	_, DBErr := t.connectionDB.Delete(tarantoolConfig.DBSpaceName, "primary", []interface{}{cookie.Value})

	if DBErr != nil {
		return DBErr
	}

	return nil
}
