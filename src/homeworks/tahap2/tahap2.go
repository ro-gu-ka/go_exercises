package tahap2

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func GenerateConsecutiveNextNik(nik string, numberGenerated int) []string {
	var coll []string

	splitted := strings.Split(nik, "-")
	// validation
	if !strings.Contains(splitted[0], "ARN") || strings.Contains(splitted[0], "ART") {
		log.Fatal("Invalid NIK format: ", splitted[0])
	}
	if len([]rune(splitted[1])) != 5 {
		log.Fatal("Invalid NIK number format: ", splitted[1])
	}

	// string to int
	n, err := strconv.Atoi(splitted[1])
	if err != nil {
		// handle error
		panic(err)
	}
	for i := 1; i <= numberGenerated; i++ {
		n = n + 1
		coll = append(coll, splitted[0]+"-"+convertToStringNumber(n))
	}

	return coll
}

func GenerateNik(gender string, year int) string {
	// gender code
	g, err := getGenderCode(gender)
	if err != nil {
		log.Fatal(err)
	}

	// last 2 digits of year
	y := strconv.Itoa(year)[2:]

	// nik
	nik := "AR" + g + y + randomSemester() + "-" + randomNumber()
	return nik
}

func GenerateNikCollection(gender string, year int, numberGenerated int) []string {
	var coll []string
	for i := 1; i <= numberGenerated; i++ {
		nik := GenerateNik(gender, year)
		coll = append(coll, nik)
	}
	return coll
}

func getGenderCode(gender string) (string, error) {
	if gender == "male" || gender == "m" {
		return "N", nil
	} else if gender == "female" || gender == "f" {
		return "T", nil
	}
	return gender, errors.New("Wrong gender input!")
}

func randomSemester() string {
	// Only two values, "1" is b/w 1-6 and "2" is b/w 7-12
	formats := []string{"1", "2"}

	// Return a randomly selected semester
	return formats[rand.Intn(len(formats))]
}

func randomNumber() string {
	return fmt.Sprintf("%05d", rand.Intn(1000))
}

func convertToStringNumber(num int) string {
	return fmt.Sprintf("%05d", num)
}
