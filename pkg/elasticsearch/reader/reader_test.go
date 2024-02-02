package reader_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch/reader"
)

type Person struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

const (
	singleObjectJson = `{ "name": "unknown", "age": 20 }`
	listOfObjects    = `[ { "name": "unknown", "age": 20 }, { "name": "new person", "age": 75 } ]`
)

func TestReadJson(t *testing.T) {
	t.Run("Returns and", func(t *testing.T) {
		_, err := reader.ReadJson[Person](fstest.MapFS{}, "this path does not exists")

		want := reader.InvalidPathError
		if err != want {
			t.Errorf("expected %s but got %s", want, err)
		}
	})

	t.Run("read a single object from json", func(t *testing.T) {
		fs := fstest.MapFS{"data.json": {Data: []byte(singleObjectJson)}}
		got, err := reader.ReadJson[Person](fs, "data.json")

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		want := Person{Name: "unknown", Age: 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("read a list of objects", func(t *testing.T) {
		fs := fstest.MapFS{"data.json": {Data: []byte(listOfObjects)}}
		got, err := reader.ReadJson[[]Person](fs, "data.json")

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		want := []Person{{Name: "unknown", Age: 20}, {Name: "new person", Age: 75}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
