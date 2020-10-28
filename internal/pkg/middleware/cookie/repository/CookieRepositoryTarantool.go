package repository

import (
	tarantoolConfig "backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"strconv"
)

type CookieTarantoolRepository struct {
	connectionDB *tarantool.Connection
}

func NewCookieTarantoolRepository(conn *tarantool.Connection) (*CookieTarantoolRepository, error) {
	return &CookieTarantoolRepository{
		connectionDB: conn,
	}, nil
}

func (t *CookieTarantoolRepository) GetCookie(cookie *http.Cookie) (uint64, error) {
	resp, DBErr := t.connectionDB.Eval("return check_session(...)", []interface{}{cookie.Value})
	if DBErr != nil {
		return 0, DBErr
	}
	if resp == nil {
		return 0, errors.New("incorrect session")
	}
	data := resp.Data[0].([]interface{})
	tarantoolRes := new(models.TarantoolResponse)
	if data != nil{
		if len(data) > 0{
			tarantoolRes.CookieValue = data[0].(string)
		}
		if len(data) > 1{
			rawUserID := data[0].(string)
			userIDInt, translationErr := strconv.Atoi(rawUserID)
			if translationErr != nil{
				return 0, errors.New("bad cookie ")
			}
			tarantoolRes.UserID = uint64(userIDInt)
		}
		if len(data) > 2{
			rawCookie := data[0].(string)
			translationErr := json.Unmarshal([]byte(rawCookie), &tarantoolRes.Cookie)
			if translationErr != nil{
				return 0, errors.New("bad cookie")
			}
		}
	}
	fmt.Println(data)
	return 0, errors.New("cookie not found")
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
