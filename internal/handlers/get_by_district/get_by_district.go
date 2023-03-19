package getByDistrict

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"main/internal/city"
	"net/http"
)

type Getter interface {
	GetByDistrict(string) []*city.City
}

func New(bd Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get by district method")
		region := chi.URLParam(r, "district")
		city := bd.GetByDistrict(region)
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
