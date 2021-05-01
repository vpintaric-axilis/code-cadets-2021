package main

import (
	"io/ioutil"
	"log"
	"time"

	// NOTE - last element of the path is the package name
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

// NOTE - inferred `string` type
const yesNoURL = "http://yesno.wtf/api"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func main() {
	// NOTE - public/private is resolved via capitalisation
	// NOTE - `:=` vs `=` assignment
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	// NOTE - there are 2 values returned
	// NOTE - error handling
	httpResponse, err := httpClient.Get(yesNoURL)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	// NOTE - the `%v` verb
	// NOTE - type conversion `[]byte` to `string`
	log.Printf("Response from yesno: %v", string(bodyContent))
}
