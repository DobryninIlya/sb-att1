package getByDmain

import (
	"encoding/json"
	"fmt"
	"io"
	"main/internal/city"
	"net/http"
)

type diapazon struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type Getter interface {
	GetByFoundationDiapazon(int, int) []*city.City
}

func New(bd Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get by population diapazon method")
		defer r.Body.Close()
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		var d diapazon
		err = json.Unmarshal(content, &d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		city := bd.GetByFoundationDiapazon(d.From, d.To)
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
