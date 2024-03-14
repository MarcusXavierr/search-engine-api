package reader

import (
	"encoding/json"
	"errors"
	"io/fs"
)

var (
	InvalidPathError       = errors.New("The given path is invalid. Or the file wasn't found")
	InvalidJsonFormatError = errors.New("There was an error running Unmarshal on your file's content, description below ")
)

// Simply reads a Json file and return it's content as Struct
func ReadJson[T any](fileSystem fs.FS, path string) (T, error) {
	var data T
	bytes, err := fs.ReadFile(fileSystem, path)

	if err != nil {
		return data, InvalidPathError
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return data, errors.Join(InvalidJsonFormatError, err)
	}

	return data, nil
}
