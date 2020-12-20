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

func NewReplyDelivery(useCase replies.UseCase) *ReplyDelivery {
	return &ReplyDelivery{
		useCase: useCase,
	}
}

// Reply godoc
// @Summary CreateReply
// @Description Create reply to movie
// @ID create-reply-id
// @Param Reply_info body models.ReplyInput true "Login information"
// @Success 200
// @Failure 400 {object} models.ServerResponse
// @Failure 401 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/reply/ [post]
func (t *ReplyDelivery) CreateReply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		models.BadMethodHTTPResponse(&w)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}
	input := new(models.ReplyInput)
	defer func() {
		_ = r.Body.Close()
	}()
	translationErr := json.NewDecoder(r.Body).Decode(&input)
	if translationErr != nil {
		models.BadBodyHTTPResponse(&w, models.ErrFooIncorrectInputInfo)
	}

	err := t.useCase.CreateReply(input, userID.(uint64))
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}
}

// Reply godoc
// @Summary GetMovieReplies
// @Description Get movie reply list
// @ID movie-reply-list-id
// @Param limit query int true "movie_id"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {array} models.ReplyModel
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/reply/ [get]
func (t *ReplyDelivery) GetMovieReplies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.BadMethodHTTPResponse(&w)
		return
	}

	movieID, movieOk := mux.Vars(r)[replies.MovieIDQuery]
	limit, limitOk := mux.Vars(r)[movieservice.LimitQuery]
	page, pageOk := mux.Vars(r)[movieservice.PageQuery]
	if !movieOk || !limitOk || !pageOk {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	castedMovieID, err := strconv.Atoi(movieID)
	if err != nil {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	castedLimit, err := strconv.Atoi(limit)
	if err != nil {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	castedPage, err := strconv.Atoi(page)
	if err != nil {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	resp, err := t.useCase.GetMovieReplies(castedMovieID, castedLimit, castedPage)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	outputBuf, _ := json.Marshal(resp)
	w.Write(outputBuf)
}
