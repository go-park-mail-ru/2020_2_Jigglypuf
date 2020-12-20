package replyserver

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/replies/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type ReplyService struct {
	ReplyRepository replies.Repository
	ReplyUseCase    replies.UseCase
	ReplyDelivery   *delivery.ReplyDelivery
	ReplyRouter     *mux.Router
}

func configureReplyRouter(handler *delivery.ReplyDelivery) *mux.Router {
	replyRouter := mux.NewRouter()
	replyRouter.HandleFunc(utils.ReplyURLPattern, handler.CreateReply).Methods(http.MethodPost)
	replyRouter.HandleFunc(utils.ReplyURLPattern, handler.GetMovieReplies).Methods(http.MethodGet)

	return replyRouter
}

func Start(profileConn profileService.ProfileServiceClient, connection *sql.DB) (*ReplyService, error) {
	if connection == nil {
		return nil, models.ErrFooNoDBConnection
	}
	replyRep := repository.NewReplyRepository(connection)
	replyUC := usecase.NewReplyUseCase(replyRep, profileConn)
	replyDelivery := delivery.NewReplyDelivery(replyUC)
	router := configureReplyRouter(replyDelivery)
	return &ReplyService{
		ReplyDelivery:   replyDelivery,
		ReplyUseCase:    replyUC,
		ReplyRepository: replyRep,
		ReplyRouter:     router,
	}, nil
}
