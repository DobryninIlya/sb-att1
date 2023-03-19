package create

import (
	"encoding/json"
	"fmt"
	"io"
	"main/internal/city"
	"net/http"
)

type Adder interface {
	AddNew(*city.City)
	GetById(int) *city.City
}

func New(bd Adder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("add method")
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
		if bd.GetById(c.Id) == nil {
			bd.AddNew(&c)
			w.WriteHeader(http.StatusCreated)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Already exist"))
		}
		return

	}
}
