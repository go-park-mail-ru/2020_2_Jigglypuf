package usecase

import(
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation/mock"
	"github.com/golang/mock/gomock"
	"sync"
	"testing"
	"time"
)



func TestFirstExample(t *testing.T){
	controller := gomock.NewController(t)
	repMock := mock.NewMockRepository(controller)
	RepositoryResponse := []models.RecommendationDataFrame{
		{
			uint64(1),
			"first movie",
			uint64(1),
			6,
			7.5,
			10,
		},
		{
			uint64(1),
			"fourth movie",
			uint64(4),
			6,
			7.5,
			10,
		},
		{
			uint64(2),
			"first movie",
			uint64(1),
			11,
			7.5,
			10,
		},
		{
			uint64(9),
			"first movie",
			uint64(1),
			6,
			7.5,
			10,
		},

		{
			uint64(5),
			"second movie",
			uint64(2),
			11,
			7.5,
			10,
		},
		{
			uint64(9),
			"second movie",
			uint64(2),
			10,
			7.5,
			10,
		},
		{
			uint64(1),
			"second movie",
			uint64(2),
			10,
			7.5,
			10,
		},
		{
			uint64(9),
			"third movie",
			uint64(3),
			8,
			7.5,
			10,
		},
		{
			uint64(5),
			"third movie",
			uint64(3),
			11,
			7.5,
			10,
		},
	}


	repMock.EXPECT().GetMovieRatingsDataset().AnyTimes().Return(&RepositoryResponse, nil)
	sys := NewRecommendationSystemUseCase(repMock, time.Minute*10, &sync.RWMutex{})
	_, _ = sys.MakeMovieRecommendations(uint64(1))

}