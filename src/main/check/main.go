package main

import (
	"bytes"
	"concourse-blob-resource/src/concourse"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	output, err := Serialize(CreateConcourseCheckRequest)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	} else {
		fmt.Println(output)
	}
}

func Serialize(action func() concourse.CheckResponse) (string, error) {
	response := action()
	outStream := bytes.NewBufferString("")

	err := json.NewEncoder(outStream).Encode(response)
	if err != nil {
		return "", fmt.Errorf("encode json: %s", err)
	}

	return outStream.String(), nil
}

func CreateConcourseCheckRequest() concourse.CheckResponse {
	return concourse.CheckResponse([]concourse.Version{{ProductVersion: "NotImplemented"}})
}