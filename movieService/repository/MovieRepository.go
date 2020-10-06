package repository

import (
	"models"
	"sync"
)

type MovieRepository struct{
	Movies []models.Movie
	mutex *sync.RWMutex
}

type MovieAlreadyExists struct{}
type MovieNotFound struct{}
type PageNotFound struct{}

func (t PageNotFound) Error() string{
	return "PageNotFound!"
}

func (t MovieAlreadyExists) Error() string{
	return "Movie Already Exists!"
}

func (t MovieNotFound) Error() string{
	return "Movie Already Exists!"
}

func NewMovieRepository(mutex *sync.RWMutex) *MovieRepository{
	return &MovieRepository{
		Movies: []models.Movie{
			models.Movie{
				Id: 1,
				Name: "Гренландия",
				Description: "Greenland description",
				PathToAvatar: "/media/greenland.jpg",
			},
			models.Movie{
				Id: 2,
				Name: "Антибеллум",
				Description: "Антибеллум description",
				PathToAvatar: "/media/antibellum.jpg",
			},
			models.Movie{
				Id: 3,
				Name: "Довод",
				Description: "Довод description",
				PathToAvatar: "/media/dovod.jpg",
			},
			models.Movie{
				Id: 4,
				Name: "Гнездо",
				Description: "Гнездо description",
				PathToAvatar: "/media/gnezdo.jpg",
			},
			models.Movie{
				Id: 5,
				Name: "Сделано в Италии",
				Description: "Италиан description",
				PathToAvatar: "/media/italian.jpg",
			},
			models.Movie{
				Id: 6,
				Name: "Мулан",
				Description: "Мулан description",
				PathToAvatar: "/media/mulan.jpg",
			},
			models.Movie{
				Id: 7,
				Name: "Никогда всегда всегда никогда",
				Description: "Никогда description",
				PathToAvatar: "/media/nikogda.jpg",
			},
			models.Movie{
				Id: 8,
				Name: "После",
				Description: "После description",
				PathToAvatar: "/media/posle.jpg",
			},
			models.Movie{
				Id: 9,
				Name: "Стрельцов",
				Description: "Стрельцов description",
				PathToAvatar: "/media/strelcov.jpg",
			},
		},
		mutex: mutex,
	}
}


func (t *MovieRepository) CreateMovie( movie *models.Movie )error{
	success := true
	var id uint64 = 0
	t.mutex.RLock()
	{
		id = t.Movies[len(t.Movies) -1].Id + 1
		for _, val := range t.Movies {
			if val.Name == movie.Name {
				success = false
				break
			}
		}
	}
	t.mutex.Unlock()

	if !success{
		return MovieAlreadyExists{}
	}
	movie.Id = id
	t.mutex.Lock()
	{
		t.Movies = append(t.Movies, *movie)
	}
	t.mutex.Unlock()

	return nil
}

func (t *MovieRepository) UpdateMovie(movie *models.Movie)error{
	return nil
}

func (t *MovieRepository) GetMovie(name string)(*models.Movie, error){
	resultMovie := new(models.Movie)
	success := false
	t.mutex.RLock()
	{
		for _,val := range t.Movies{
			if val.Name == name{
				*resultMovie = val
				success = true
				break
			}
		}
	}
	t.mutex.RUnlock()

	if !success{
		return resultMovie, MovieNotFound{}
	}

	return resultMovie, nil
}

func (t *MovieRepository) GetMovieList(limit, page int)(*[]models.Movie, error){
	resultArray := make([]models.Movie,0)
	success := true

	t.mutex.RLock()
	{
		startIndex := (page-1)*limit
		endIndex := len(t.Movies) - startIndex
		if startIndex > len(t.Movies){
			success = false
		}else{
			if endIndex > limit{
				endIndex = limit
			}
			resultArray = t.Movies[startIndex:endIndex]
		}
	}
	t.mutex.RUnlock()

	if !success{
		return &resultArray,PageNotFound{}
	}
	return &resultArray, nil
}