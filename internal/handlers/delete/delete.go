package delete

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type Deleter interface {
	Delete(id int) bool
}

func New(bd Deleter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("delete method")
		cityId, err := strconv.Atoi(chi.URLParam(r, "cityId"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Город не найден"))
			return
		}
		if bd.Delete(cityId) {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
