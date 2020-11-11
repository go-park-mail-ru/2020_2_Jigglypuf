package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HashCSRFToken struct{
	Secret []byte
	Duration time.Duration
}

type Response struct{
	Token string
}

func NewHashCSRFToken(secret string, duration time.Duration)(*HashCSRFToken, error){
	return &HashCSRFToken{
		[]byte(secret),
		duration,
	}, nil
}


// CSRF godoc
// @Summary Get CSRF by cookie
// @Description Returns movie schedule by ID
// @ID csrf-id
// @Success 200 {object} Response
// @Failure 400 {object} models.ServerResponse "Bad body"
// @Failure 405 {object} models.ServerResponse "Method not allowed"
// @Failure 500 {object} models.ServerResponse "internal error"
// @Router /csrf/ [get]
func (t *HashCSRFToken) GenerateCSRFToken(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}
	isAuth := r.Context().Value(cookie.ContextIsAuthName)
	userID := r.Context().Value(cookie.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) || userID == nil{
		models.UnauthorizedHTTPResponse(&w)
		return
	}

	h := hmac.New(sha256.New, t.Secret)
	tokenExpTime := time.Now().Add(t.Duration).Unix()
	csrfToken := fmt.Sprintf("%d:%d",userID.(uint64), tokenExpTime)
	h.Write([]byte(csrfToken))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)

	csrfResponse := Response{
		token,
	}

	outputBuf, _ := json.Marshal(csrfResponse)
	_, _ = w.Write(outputBuf)
}

func (t *HashCSRFToken) CSRFMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost && r.Method != http.MethodPut && r.Method != http.MethodDelete{
			next.ServeHTTP(w,r)
			return
		}

		isAuth := r.Context().Value(cookie.ContextIsAuthName)
		userID := r.Context().Value(cookie.ContextUserIDName)
		if isAuth == nil || userID == nil || !isAuth.(bool){
			next.ServeHTTP(w,r)
			return
		}
		inputCSRFToken := r.Header.Get("X-CSRF-Token")
		if inputCSRFToken != "" && t.CheckCSRFToken(userID.(uint64), inputCSRFToken){
			next.ServeHTTP(w,r)
			return
		}
		w.WriteHeader(http.StatusForbidden)
	})
}


func (t *HashCSRFToken) CheckCSRFToken(userID uint64, token string) bool{
	tokenSplit := strings.Split(token, ":")
	if len(tokenSplit) != 2{
		log.Println("incorrect token split length")
		return false
	}

	tokenExpTime, timeErr := strconv.ParseInt(tokenSplit[1], 10, 64)
	if timeErr != nil{
		log.Println("incorrect time in token")
		return false
	}

	if tokenExpTime < time.Now().Unix(){
		log.Println("token expired")
		return false
	}

	inputString := fmt.Sprintf("%d:%d",userID, tokenExpTime)
	h := hmac.New(sha256.New,t.Secret)
	h.Write([]byte(inputString))

	expectedMAC := h.Sum(nil)
	requestMAC, decodeErr := hex.DecodeString(tokenSplit[0])
	if decodeErr != nil{
		log.Println("cannot decode token hex string")
		return false
	}
	log.Println(hmac.Equal(requestMAC,expectedMAC))
	return hmac.Equal(requestMAC,expectedMAC)
}
