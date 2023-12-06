package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type fileRequest struct {
	fileName   string
	data       []byte
	permission uint32
}

func TestCreateFile(t *testing.T) {
	testCases := []struct {
		description string
		name        string
		request     fileRequest
		want        error
	}{
		{
			description: "it should return an error",
			name:        "return_error_when_creating_file",
			request: fileRequest{
				fileName:   "test/test.txt",
				data:       []byte("Hello World"),
				permission: 0777,
			},
			want: &os.PathError{},
		},
		{
			description: "it should return create a file",
			name:        "create_a_file",
			request: fileRequest{
				fileName:   "test.txt",
				data:       []byte("Hello World"),
				permission: 0777,
			},
			want: nil,
		},
	}

	f := New()
	for _, testCase := range testCases {
		err := f.Create(testCase.request.fileName, testCase.request.data, testCase.request.permission)
		assert.IsType(t, err, testCase.want)
	}

}

func TestReadFile(t *testing.T) {
	testCases := []struct {
		description string
		name        string
		request     string
		want        error
	}{
		{
			description: "it should return an error",
			name:        "return_error_when_reading_file",
			request:     "",
			want:        &os.PathError{},
		},
		{
			description: "it should read file",
			name:        "get_file",
			request:     "test.txt",
			want:        nil,
		},
	}

	f := New()
	for _, testCase := range testCases {
		_, err := f.Read(testCase.request)
		assert.IsType(t, err, testCase.want)
	}
}

func TestRemoveFile(t *testing.T) {
	testCases := []struct {
		description string
		name        string
		request     string
		want        error
	}{
		{
			description: "it should return an error",
			name:        "return_error_when_removing_file",
			request:     "",
			want:        &os.PathError{},
		},
		{
			description: "it should read remove",
			name:        "get_file",
			request:     "test.txt",
			want:        nil,
		},
	}

	f := New()
	for _, testCase := range testCases {
		err := f.Remove(testCase.request)
		assert.IsType(t, err, testCase.want)
	}
}
