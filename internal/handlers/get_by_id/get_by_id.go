package getById

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"main/internal/city"
	"net/http"
	"strconv"
)

type Getter interface {
	GetById(int) *city.City
}

func New(bd Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get_by_id method")
		cityId, err := strconv.Atoi(chi.URLParam(r, "cityId"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Город не найден"))
			return
		}
		city := bd.GetById(cityId)
		if city == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if cityJson, err := json.Marshal(&city); err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write(cityJson)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
