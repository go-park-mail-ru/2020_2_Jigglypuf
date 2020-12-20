package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ReplyDelivery struct {
	useCase replies.UseCase
}

func NewReplyDelivery(useCase replies.UseCase) *ReplyDelivery{
	return &ReplyDelivery{
		useCase: useCase,
	}
}

func (t *ReplyDelivery) CreateReply(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		models.BadMethodHTTPResponse(&w)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || isAuth.(bool) == false{
		models.UnauthorizedHTTPResponse(&w)
		return
	}
	input := new(models.ReplyInput)
	defer func(){
		_= r.Body.Close()
	}()
	translationErr := json.NewDecoder(r.Body).Decode(&input)
	if translationErr != nil{
		models.BadBodyHTTPResponse(&w, models.ErrFooIncorrectInputInfo)
	}

	err := t.useCase.CreateReply(input, userID.(uint64))
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}
}

func (t *ReplyDelivery) GetMovieReplies(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		models.BadMethodHTTPResponse(&w)
		return
	}

	movieID, movieOk := mux.Vars(r)[replies.MovieIdQuery]
	limit, limitOk := mux.Vars(r)[movieservice.LimitQuery]
	page, pageOk := mux.Vars(r)[movieservice.PageQuery]
	if !movieOk || !limitOk || !pageOk{
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	castedMovieID, err := strconv.Atoi(movieID)
	if err != nil{
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	castedLimit, err := strconv.Atoi(limit)
	if err != nil{
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	castedPage, err := strconv.Atoi(page)
	if err != nil{
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	resp, err := t.useCase.GetMovieReplies(castedMovieID, castedLimit, castedPage)
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	outputBuf, _ := json.Marshal(resp)
	w.Write(outputBuf)
}