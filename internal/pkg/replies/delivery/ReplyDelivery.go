package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
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
// @Param movie_id query int true "movie_id"
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
	w.Header().Set("Content-Type", "application/json")

	movie := r.URL.Query()[replies.MovieIDQuery]
	Limit := r.URL.Query()[movieservice.LimitQuery]
	Page := r.URL.Query()[movieservice.PageQuery]
	if len(Limit) == 0 || len(Page) == 0 || len(movie) == 0 {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}
	movieID, movieErr := strconv.Atoi(movie[0])
	limit, limitErr := strconv.Atoi(Limit[0])
	page, pageErr := strconv.Atoi(Page[0])
	if limitErr != nil || pageErr != nil || movieErr != nil {
		models.IncorrectGetParamsHTTPResponse(&w)
		return
	}

	resp, err := t.useCase.GetMovieReplies(movieID, limit, page)
	if err != nil {
		models.BadBodyHTTPResponse(&w, err)
		return
	}

	outputBuf, _ := json.Marshal(resp)
	w.Write(outputBuf)
}


// Reply godoc
// @Summary UpdateReply
// @Description Update Reply
// @Param Reply_info body models.ReplyUpdateInput true "Reply Update information"
// @Success 200
// @Failure 400 {object} models.ServerResponse
// @Failure 405 {object} models.ServerResponse
// @Router /api/reply/ [put]
func (t *ReplyDelivery) UpdateReply(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut{
		models.BadMethodHTTPResponse(&w)
		return
	}

	isAuth := r.Context().Value(cookieService.ContextIsAuthName)
	userID := r.Context().Value(cookieService.ContextUserIDName)
	if isAuth == nil || !isAuth.(bool) {
		models.UnauthorizedHTTPResponse(&w)
		return
	}
	input := new(models.ReplyUpdateInput)
	defer func(){
		_ = r.Body.Close()
	}()
	decodeErr := json.NewDecoder(r.Body).Decode(&input)
	if decodeErr != nil{
		models.BadBodyHTTPResponse(&w, models.ErrFooIncorrectInputInfo)
		return
	}

	err := t.useCase.UpdateReply(input, userID.(uint64))
	if err != nil{
		models.BadBodyHTTPResponse(&w, err)
		return
	}
}