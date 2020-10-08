package repository

import (
	"models"
	"sync"
)

type CinemaRepository struct{
	Cinema []models.Cinema
	mu *sync.RWMutex
}

func NewCinemaRepository (mutex *sync.RWMutex) *CinemaRepository{
	return &CinemaRepository{
		Cinema: []models.Cinema{
			models.Cinema{
				1,
				"CinemaScope1",
				"Москва, Первая улица, д.1",
			},
			models.Cinema{
				2,
				"CinemaScope2",
				"Москва, Первая улица, д.2",
			},
			models.Cinema{
				3,
				"CinemaScope3",
				"Москва, Первая улица, д.3",
			},
			models.Cinema{
				4,
				"CinemaScope4",
				"Москва, Первая улица, д.4",
			},
		},
		mu: mutex,
	}
}

type CinemaAlreadyExists struct{}
type CinemaNotFound struct{}
type PageNotFound struct{}

func (t CinemaAlreadyExists) Error()string{
	return "CinemaAlreayExists!"
}

func (t CinemaNotFound) Error()string{
	return "CinemaNotFound!"
}

func (t PageNotFound) Error() string{
	return "PageNotFound!"
}

func (t *CinemaRepository) CreateCinema(cinema *models.Cinema ) error{
	success := true
	var id uint64 = 0
	t.mu.RLock()
	{
		id = t.Cinema[len(t.Cinema) -1].Id + 1
		for _, val := range t.Cinema {
			if val.Name == cinema.Name {
				success = false
				break
			}
		}
	}
	t.mu.Unlock()

	if !success{
		return CinemaAlreadyExists{}
	}
	cinema.Id = id
	t.mu.Lock()
	{
		t.Cinema = append(t.Cinema, *cinema)
	}
	t.mu.Unlock()

	return nil
}


func (t *CinemaRepository) GetCinema( name *string )( *models.Cinema, error ){
	resultCinema := new(models.Cinema)
	success := false
	t.mu.RLock()
	{
		for _,val := range t.Cinema{
			if val.Name == *name{
				*resultCinema = val
				success = true
				break
			}
		}
	}
	t.mu.RUnlock()

	if !success{
		return resultCinema, CinemaNotFound{}
	}

	return resultCinema, nil
}

func (t *CinemaRepository) GetCinemaList(limit int, page int)( *[]models.Cinema, error ){
	resultArray := make([]models.Cinema,0)
	success := true

	t.mu.RLock()
	{
		startIndex := (page-1)*limit
		endIndex := len(t.Cinema) - startIndex
		if startIndex > len(t.Cinema){
			success = false
		}else{
			if endIndex > limit{
				endIndex = limit
			}
			resultArray = t.Cinema[startIndex:endIndex]
		}
	}
	t.mu.RUnlock()

	if !success{
		return &resultArray,PageNotFound{}
	}
	return &resultArray, nil
}

func (t *CinemaRepository) UpdateCinema( cinema *models.Cinema ) error{
	return nil
}

func (t *CinemaRepository) DeleteCinema( cinema *models.Cinema ) error{
	return nil
}