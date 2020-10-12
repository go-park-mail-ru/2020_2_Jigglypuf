package repository

import (
	"backend/internal/pkg/models"
	"sync"
)

type MovieRepository struct {
	Movies  []models.Movie
	Ratings map[string][]models.RatingSet
	mutex   *sync.RWMutex
}

type MovieAlreadyExists struct{}
type MovieNotFound struct{}
type PageNotFound struct{}
type MovieAlreadyRated struct{}

func (t PageNotFound) Error() string {
	return "PageNotFound!"
}

func (t MovieAlreadyExists) Error() string {
	return "Movie Already Exists!"
}

func (t MovieNotFound) Error() string {
	return "Movie Not Found!"
}

func (t MovieAlreadyRated) Error() string {
	return "Movie Already rated!"
}

func NewMovieRepository(mutex *sync.RWMutex) *MovieRepository {
	return &MovieRepository{
		Movies: []models.Movie{
			{
				ID:           1,
				Name:         "Гренландия",
				Description:  "Greenland description",
				PathToAvatar: "/media/greenland.jpg",
			},
			{
				ID:           2,
				Name:         "Антибеллум",
				Description:  "Антибеллум description",
				PathToAvatar: "/media/antibellum.jpg",
			},
			{
				ID:           3,
				Name:         "Довод",
				Description:  "Довод description",
				PathToAvatar: "/media/dovod.jpg",
			},
			{
				ID:           4,
				Name:         "Гнездо",
				Description:  "Гнездо description",
				PathToAvatar: "/media/gnezdo.jpg",
			},
			{
				ID:           5,
				Name:         "Сделано в Италии",
				Description:  "Италиан description",
				PathToAvatar: "/media/italian.jpg",
			},
			{
				ID:           6,
				Name:         "Мулан",
				Description:  "Мулан description",
				PathToAvatar: "/media/mulan.jpg",
			},
			{
				ID:           7,
				Name:         "Никогда всегда всегда никогда",
				Description:  "Никогда description",
				PathToAvatar: "/media/nikogda.jpg",
			},
			{
				ID:           8,
				Name:         "После",
				Description:  "После description",
				PathToAvatar: "/media/posle.jpg",
			},
			{
				ID:           9,
				Name:         "Стрельцов",
				Description:  "Стрельцов description",
				PathToAvatar: "/media/strelcov.jpg",
			},
		},
		mutex:   mutex,
		Ratings: make(map[string][]models.RatingSet),
	}
}

func (t *MovieRepository) CreateMovie(movie *models.Movie) error {
	success := true
	var id uint64 = 0
	t.mutex.RLock()
	{
		id = t.Movies[len(t.Movies)-1].ID + 1
		for _, val := range t.Movies {
			if val.Name == movie.Name {
				success = false
				break
			}
		}
	}
	t.mutex.Unlock()

	if !success {
		return MovieAlreadyExists{}
	}
	movie.ID = id
	t.mutex.Lock()
	{
		t.Movies = append(t.Movies, *movie)
	}
	t.mutex.Unlock()

	return nil
}

func (t *MovieRepository) UpdateMovie(movie *models.Movie) error {
	return nil
}

func (t *MovieRepository) GetMovie(name string) (*models.Movie, error) {
	resultMovie := new(models.Movie)
	success := false
	t.mutex.RLock()
	{
		for _, val := range t.Movies {
			if val.Name == name {
				*resultMovie = val
				success = true
				break
			}
		}
	}
	t.mutex.RUnlock()

	if !success {
		return resultMovie, MovieNotFound{}
	}

	return resultMovie, nil
}

func (t *MovieRepository) GetMovieList(limit, page int) (*[]models.Movie, error) {
	resultArray := make([]models.Movie, 0)
	success := true

	t.mutex.RLock()
	{
		startIndex := (page - 1) * limit
		endIndex := len(t.Movies) - startIndex
		if startIndex > len(t.Movies) {
			success = false
		} else {
			if endIndex > limit {
				endIndex = limit
			}
			resultArray = t.Movies[startIndex:endIndex]
		}
	}
	t.mutex.RUnlock()

	if !success {
		return &resultArray, PageNotFound{}
	}
	return &resultArray, nil
}

func (t *MovieRepository) RateMovie(user *models.User, name string, rating int64) error {
	movie, err := t.GetMovie(name)
	if err != nil {
		return err
	}
	success := true
	t.mutex.RLock()
	{
		for _, val := range t.Ratings[user.Username] {
			if val.MovieRating.Name == movie.Name {
				success = false
			}
		}
	}
	t.mutex.RUnlock()

	if !success {
		return MovieAlreadyRated{}
	}

	t.mutex.Lock()
	{
		t.Ratings[user.Username] = append(t.Ratings[user.Username], models.RatingSet{
			MovieRating: movie,
			Rating:      rating,
		})
	}
	t.mutex.Unlock()

	return nil
}

func (t *MovieRepository) GetRating(user *models.User, name string) (int64, error) {
	movie, err := t.GetMovie(name)
	if err != nil {
		return 0, err
	}
	success := false
	var result int64
	t.mutex.RLock()
	{
		for _, val := range t.Ratings[user.Username] {
			if val.MovieRating.Name == movie.Name {
				success = true
				result = val.Rating
			}
		}
	}
	t.mutex.RUnlock()
	if !success {
		return result, MovieNotFound{}
	}
	return result, nil
}
