package population

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"main/internal/city"
	"net/http"
	"strconv"
)

type Updater interface {
	UpdatePopulation(int, int) bool
}

func New(bd Updater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("patch population method")
		cityId, err := strconv.Atoi(chi.URLParam(r, "cityId"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Город не найден"))
			return
		}
		defer r.Body.Close()
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var c city.City
		err = json.Unmarshal(content, &c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
			return
		}
		if bd.UpdatePopulation(cityId, c.Population) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
