package repository

import (
	"backend/internal/pkg/models"
	"sync"
)

type CinemaRepository struct {
	Cinema []models.Cinema
	mu     *sync.RWMutex
}

func NewCinemaRepository(mutex *sync.RWMutex) *CinemaRepository {
	return &CinemaRepository{
		Cinema: []models.Cinema{

			{
				ID: 1,
				Name:"CinemaScope1",
				Address: "Москва, Первая улица, д.1",
			},
			{
				ID:2,
				Name:"CinemaScope2",
				Address: "Москва, Первая улица, д.2",
			},
			{
				ID:3,
				Name:"CinemaScope3",
				Address:"Москва, Первая улица, д.3",
			},
			{
				ID:4,
				Name:"CinemaScope4",
				Address:"Москва, Первая улица, д.4",
			},
		},
		mu: mutex,
	}
}

type CinemaAlreadyExists struct{}
type CinemaNotFound struct{}
type PageNotFound struct{}

func (t CinemaAlreadyExists) Error() string {
	return "CinemaAlreayExists!"
}

func (t CinemaNotFound) Error() string {
	return "CinemaNotFound!"
}

func (t PageNotFound) Error() string {
	return "PageNotFound!"
}

func (t *CinemaRepository) CreateCinema(cinema *models.Cinema) error {
	success := true
	var id uint64 = 0
	t.mu.RLock()
	{
		id = t.Cinema[len(t.Cinema)-1].ID + 1
		for _, val := range t.Cinema {
			if val.Name == cinema.Name {
				success = false
				break
			}
		}
	}
	t.mu.Unlock()

	if !success {
		return CinemaAlreadyExists{}
	}
	cinema.ID = id
	t.mu.Lock()
	{
		t.Cinema = append(t.Cinema, *cinema)
	}
	t.mu.Unlock()

	return nil
}

func (t *CinemaRepository) GetCinema(id uint64) (*models.Cinema, error) {
	resultCinema := new(models.Cinema)
	success := false
	t.mu.RLock()
	{
		for _, val := range t.Cinema {
			if val.ID == id {
				*resultCinema = val
				success = true
				break
			}
		}
	}
	t.mu.RUnlock()

	if !success {
		return resultCinema, CinemaNotFound{}
	}

	return resultCinema, nil
}

func (t *CinemaRepository) GetCinemaList(limit int, page int) (*[]models.Cinema, error) {
	resultArray := make([]models.Cinema, 0)
	success := true

	t.mu.RLock()
	{
		startIndex := (page - 1) * limit
		endIndex := len(t.Cinema) - startIndex
		if startIndex > len(t.Cinema) {
			success = false
		} else {
			if endIndex > limit {
				endIndex = limit
			}
			resultArray = t.Cinema[startIndex:endIndex]
		}
	}
	t.mu.RUnlock()

	if !success {
		return &resultArray, PageNotFound{}
	}
	return &resultArray, nil
}

func (t *CinemaRepository) UpdateCinema(cinema *models.Cinema) error {
	return nil
}

func (t *CinemaRepository) DeleteCinema(cinema *models.Cinema) error {
	return nil
}
