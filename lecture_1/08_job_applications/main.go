package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Applicant struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

const url = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"
const outputFileName = "output.txt"

var requiredSkills = map[string]struct{}{
	"Java": {},
	"Go":   {},
}

func containsAnyRequiredSkill(skills []string) bool {
	for _, skill := range skills {
		_, contains := requiredSkills[skill]
		if contains {
			return true
		}
	}

	return false
}

func main() {
	httpClient := &http.Client{}

	rawResponse, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var applicants []Applicant
	err = json.Unmarshal(body, &applicants)
	if err != nil {
		log.Fatal(err)
	}

	var filteredApplicants []Applicant
	for _, applicant := range applicants {
		if applicant.Passed && containsAnyRequiredSkill(applicant.Skills) {
			filteredApplicants = append(filteredApplicants, applicant)
		}
	}

	file, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, applicant := range filteredApplicants {
		file.WriteString(applicant.Name)
		file.WriteString(" - ")
		file.WriteString(strings.Join(applicant.Skills, ", "))
		file.WriteString("\n")
	}

	file.Sync()
}
