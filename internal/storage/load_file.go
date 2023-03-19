package storage

import (
	"bufio"
	"fmt"
	"log"
	"main/internal/city"
	"os"
	"strconv"
	"strings"
)

type Storage interface {
	AddNew(city *city.City)
	GetById(int) *city.City
}

func LoadFileRAM(path string, bd *MainSotrage) *MainSotrage {

	rows, err := readFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for i, row := range rows {
		params := strings.Split(row, ",")
		if len(params) != 6 {
			log.Fatalf("Входной файл имеет неподдерживаемый формат около строки %v", i)
			return nil
		}
		id, err := strconv.Atoi(params[0])
		if err != nil {
			log.Fatal(err)
			return nil
		}
		population, err := strconv.Atoi(params[4])
		if err != nil {
			log.Fatal(err)
			return nil
		}
		foundation, err := strconv.Atoi(params[5])
		if err != nil {
			log.Fatal(err)
			return nil
		}
		bd.AddNew(&city.City{
			Id:         id,
			Name:       params[1],
			Region:     params[2],
			District:   params[3],
			Population: population,
			Foundation: foundation})
	}
	return bd
}

func readFile(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
			return nil, err
		}
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return rows, nil

}
