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
			uint64(1),
			"some movie",
			uint64(1),
			5,
			7.5,
			10,
		},
		{
			uint64(2),
			"some another movie",
			uint64(2),
			6,
			7.5,
			10,
		},
		{
			uint64(9),
			"some another movie",
			uint64(2),
			10,
			7.5,
			10,
		},
		{
			uint64(1),
			"some another movie",
			uint64(2),
			10,
			7.5,
			10,
		},
		{
			uint64(9),
			"some another movie",
			uint64(1),
			11,
			7.5,
			10,
		},
		{
			uint64(2),
			"some another movie",
			uint64(1),
			11,
			7.5,
			10,
		},
		{
			uint64(5),
			"some 3rd movie",
			uint64(3),
			7,
			7.5,
			10,
		},
	}
	repMock.EXPECT().GetMovieRatingsDataset().Return(&RepositoryResponse, nil)
	_, _ = sys.makeRecommendations()
}