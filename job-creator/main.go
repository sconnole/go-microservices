package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bxcodec/faker/v4"
)

type Job struct {
	Data string `json:"data"`
}

func main() {
	fakeSentence := faker.Word()

	job := Job{Data: fakeSentence}
	jsonData, err := json.Marshal(job)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://api.default.svc.cluster.local/jobs", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Printf("Sent Job: %s\nResponseStatus: %s\n", fakeSentence, resp.Status)
}
