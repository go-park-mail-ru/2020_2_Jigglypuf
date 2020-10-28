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

func (t *CookieTarantoolRepository) GetCookie(cookie *http.Cookie) (uint64, error) {
	resp, DBErr := t.connectionDB.Eval("return check_session(...)", []interface{}{cookie.Value})

	if DBErr != nil {
		return 0, DBErr
	}

	if resp == nil {
		return 0, errors.New("incorrect session")
	}

	data := resp.Data
	if data == nil{
		return 0, errors.New("cookie not found")
	}
	tarantoolRes := new(models.TarantoolResponse)
	tarantoolRes.ID = data[0].(uint64)
	tarantoolRes.CookieValue = data[1].(string)
	tarantoolRes.UserID = data[2].(uint64)
	tarantoolRes.Cookie = data[3].(string)
	print(tarantoolRes)
	ok := false
	if  !ok {
		return tarantoolRes.UserID, nil
	}

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
