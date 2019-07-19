package main

import (
	"bytes"
	"concourse-blob-resource/src/concourse"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	output, err := Serialize(CreateConcourseOutResponse)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	} else {
		fmt.Println(output)
	}
}

func Serialize(action func() concourse.OutResponse) (string, error) {
	response := action()
	outStream := bytes.NewBufferString("")

	err := json.NewEncoder(outStream).Encode(response)
	if err != nil {
		return "", fmt.Errorf("encode json: %s", err)
	}

	return outStream.String(), nil
}

func CreateConcourseOutResponse() concourse.OutResponse {
	return concourse.OutResponse{}
}
