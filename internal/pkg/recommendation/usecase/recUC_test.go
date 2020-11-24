package usecase

import(
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/mock"
	"github.com/golang/mock/gomock"
	"testing"
)



func TestFirstExample(t *testing.T){
	controller := gomock.NewController(t)
	repMock := mock.NewMockRepository(controller)
	sys := NewRecommendationSystemUseCase(repMock)
	RepositoryResponse := []models.RecommendationDataFrame{
		{
			1,
			"some movie",
			5,
			7.5,
			10,
		},
		{
			2,
			"another",
			5,
			7.5,
			10,
		},
	}
	repMock.EXPECT().GetMovieRatingsDataset().Return(&RepositoryResponse, nil)
	_, _ = sys.makeRecommendations()
}