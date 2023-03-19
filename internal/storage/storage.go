package storage

import (
	city "main/internal/city"
	"sync"
)

type MainSotrage struct {
	Store         map[int]*city.City
	autoincrement int
	sync.Mutex
}

func NewMainStorage() *MainSotrage {
	return &MainSotrage{
		Store:         make(map[int]*city.City),
		autoincrement: 0,
	}
}

func (ms *MainSotrage) AddNew(city *city.City) {
	ms.Lock()
	defer ms.Unlock()
	ms.Store[ms.autoincrement] = city
	ms.autoincrement += 1
}

func (ms *MainSotrage) Create() {

}

func (ms *MainSotrage) GetById(id int) *city.City {
	for _, value := range ms.Store {
		if value.Id == id {
			return value
		}
	}
	return nil
}

func (ms *MainSotrage) Delete(id int) bool {
	for k, value := range ms.Store {
		if value.Id == id {
			delete(ms.Store, k)
			return true
		}
	}
	return false
}

func (ms *MainSotrage) UpdatePopulation(id int, popultation int) bool {
	for _, value := range ms.Store {
		if value.Id == id {
			value.Population = popultation
			return true
		}
	}
	return false
}

func (ms *MainSotrage) GetByRegion(region string) (answer []*city.City) {
	for _, value := range ms.Store {
		if value.Region == region {
			answer = append(answer, value)

		}
	}
	return answer
}

func (ms *MainSotrage) GetByDistrict(district string) (answer []*city.City) {
	for _, value := range ms.Store {
		if value.District == district {
			answer = append(answer, value)

		}
	}
	return answer
}
func (ms *MainSotrage) GetByPopulationDiapazon(from int, to int) (answer []*city.City) {
	for _, value := range ms.Store {
		if value.Population >= from && value.Population <= to {
			answer = append(answer, value)
		}
	}
	return answer
}

func (ms *MainSotrage) GetByFoundationDiapazon(from int, to int) (answer []*city.City) {
	for _, value := range ms.Store {
		if value.Foundation >= from && value.Foundation <= to {
			answer = append(answer, value)
		}
	}
	return answer
}
