package usecase

import (
	set "github.com/deckarep/golang-set"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cosinesimilarity"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	config "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/rake"
	"log"
)

type ContentBasedRecommender struct {
	RecommendationRepository config.Repository
	MovieRepository 		 movieservice.MovieRepository
	Movies					 map[uint64]models.MovieList
	CounterVector			 [][]float64
	MovieIDVector			 map[int]uint64
	CosineSimilarity		 [][]float64
}

func NewContentBasedRecommender(rep config.Repository, movieRep movieservice.MovieRepository) (*ContentBasedRecommender,error){
	t :=  &ContentBasedRecommender{
		RecommendationRepository: rep,
		MovieRepository: movieRep,
		Movies: make(map[uint64]models.MovieList),
		CounterVector: make([][]float64, 0),
		MovieIDVector: make(map[int]uint64),
		CosineSimilarity: make([][]float64, 0),
	}
	err := t.createDataSet()
	if err != nil{
		return nil, err
	}
	return t, nil
}


func (t *ContentBasedRecommender) createDataSet() error{
	MovieInformation, err := t.MovieRepository.GetAllMovies()
	if err != nil{
		return err
	}
	wordsSet := set.NewSet()
	wordsCount := make(map[uint64]map[string]float64)
	for _, val := range MovieInformation{
		t.Movies[val.ID] = val
		producer := val.Producer
		country := val.Country
		genres := val.GenreList.Join(" ")
		actors := val.ActorList.Join(" ")
		description := val.Description
		cond := rake.RunRake(producer + country + genres + actors + description)
		for _, value := range cond{
			wordsSet.Add(value.Key)
			if wordsCount[val.ID] == nil{
				wordsCount[val.ID] = make(map[string]float64)
			}
			wordsCount[val.ID][value.Key] = value.Value
		}
	}
	for _, val := range MovieInformation{
		counter := make([]float64, 0)
		for _, value := range wordsSet.ToSlice(){
			amount, _ := wordsCount[val.ID][value.(string)]
			counter = append(counter, amount)
		}
		t.CounterVector = append(t.CounterVector, counter)
		t.MovieIDVector[len(t.CounterVector) - 1] =  val.ID
	}
	t.CosineSimilarity = cosinesimilarity.Compute(t.CounterVector, t.CounterVector)
	return nil
}

func (t *ContentBasedRecommender) MakeRecommendation(userID uint64, limit int) (*models.Recommendation, error){
	UserRatingHistory, err := t.RecommendationRepository.GetLastUserRating(userID)
	if err != nil || len(UserRatingHistory) < 1{
		return t.GetPopularMovies()
	}

	found := false
	lastRating := UserRatingHistory[0]
	movieIDSet := set.NewSet()
	for _,val := range UserRatingHistory{
		if val.MovieRating > config.RatingBorder && found == false{
			found = true
			lastRating = val
		}
		movieIDSet.Add(val.MovieID)
	}

	if !found{
		return t.GetPopularMovies()
	}
	index, err := t.findMovieIndex(lastRating.MovieID)
	if err != nil{
		return t.GetPopularMovies()
	}
	movieCorrelation := t.CosineSimilarity[index]
	resultArr := make([]models.MovieList, 0)
	log.Println("BEFORE SORT", movieCorrelation, t.MovieIDVector)
	Sort(movieCorrelation, t.MovieIDVector)
	log.Println("AFTER SORT", movieCorrelation, t.MovieIDVector)
	for i := 0; len(resultArr) < limit && i < len(movieCorrelation); i += 1{
		if !movieIDSet.Contains(t.MovieIDVector[i]) && movieCorrelation[i] > 0{
			resultArr = append(resultArr, t.Movies[t.MovieIDVector[i]])
		}
	}
	if len(resultArr) < 1{
		return t.GetPopularMovies()
	}
	return &models.Recommendation{RatingMovieName: t.Movies[lastRating.MovieID].Name, Movie: &resultArr}, nil
}

func Sort(input []float64, index map[int]uint64){
	for i := 0; i < len(input); i++{
		for j := i; j < len(input); j++{
			if input[j] > input[i]{
				input[j], input[i] = input[i], input[j]
				index[j], index[i] = index[i], index[j]
			}
		}
	}
}

func (t *ContentBasedRecommender) GetPopularMovies()(*models.Recommendation, error){
	movieList, err :=  t.RecommendationRepository.GetPopularMovies()
	if err != nil{
		return nil, err
	}
	return &models.Recommendation{Movie: movieList}, nil
}

func (t *ContentBasedRecommender) findMovieIndex(movieID uint64) (int, error){
	for i, val := range t.MovieIDVector{
		if val == movieID{
			return i, nil
		}
	}
	return -1, models.ErrFooIncorrectInputInfo
}




