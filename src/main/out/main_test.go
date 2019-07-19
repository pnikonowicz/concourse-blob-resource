package main_test

import (
	"concourse-blob-resource/src/concourse"
	"fmt"
	"testing"
)
import "concourse-blob-resource/src/assetssets/out"

func TestSerializeResponse(t *testing.T) {
	actual, err := main.Serialize(func() concourse.OutResponse {
		return concourse.OutResponse{
			Version: concourse.Version{ProductVersion:"product_version_value"},
			Metadata: []concourse.Metadata{},
		}
	})

	if err != nil {
		t.Errorf("does not return error %s", err)
	}

	if len(actual) == 0 {
		t.Errorf("len is greater than zero")
	}

	expected := "{\"version\":{\"product_version\":\"product_version_value\"}}\n"

	if actual != expected {
		t.Errorf("expected equals actual\n%q\n%q", expected, actual)
	}
}


func TestCreateConcourseOutResponse(t *testing.T) {
	actual := main.CreateConcourseOutResponse()

	expected := concourse.OutResponse{}

	if fmt.Sprintf("%q", expected) != fmt.Sprintf("%q", actual) {
		t.Errorf("expected equals actual\n%q\n%q", expected, actual)
	}
}
