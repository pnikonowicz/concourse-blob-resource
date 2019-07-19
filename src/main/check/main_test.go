package main_test

import (
	"concourse-blob-resource/src/concourse"
	"fmt"
	"testing"
	"concourse-blob-resource/src/main/check"
)

func TestSerializeResponse(t *testing.T) {
	actual, err := main.Serialize(func() concourse.CheckResponse {
		return concourse.CheckResponse{{ProductVersion: "expectedProductVersion"}}
	})

	if err != nil {
		t.Errorf("does not return error %s", err)
	}

	if len(actual) == 0 {
		t.Errorf("len is greater than zero")
	}

	expected := `[{"product_version":"expectedProductVersion"}]
`
	if actual != expected {
		t.Errorf("expected equals actual\n%q\n%q", expected, actual)
	}
}


func TestCreateConcourseCheckRequest(t *testing.T) {
	actual := main.CreateConcourseCheckRequest()

	expected := concourse.CheckResponse([]concourse.Version{{ProductVersion: "NotImplemented"}})

	if fmt.Sprintf("%q", expected) != fmt.Sprintf("%q", actual) {
		t.Errorf("expected equals actual\n%q\n%q", expected, actual)
	}
}
