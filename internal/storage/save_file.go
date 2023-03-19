package storage

import (
	"fmt"
	"log"
	"os"
)

func getData(st *MainSotrage) (result string) {
	for _, city := range st.Store {
		result += fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", city.Id, city.Name, city.Region, city.District, city.Population, city.Foundation)
	}
	return
}

func SaveFile(path string, sotrage *MainSotrage) error {
	err := os.WriteFile(path, []byte(getData(sotrage)), 0644)
	if err != nil {
		log.Fatal("Ошибка сохранения файла: ", err)
		return err
	}
	return nil
}
