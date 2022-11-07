package reader

import (
	"embed"
)

//go:embed resources/old-swedish-dictionary.json
var jsonFile embed.FS

func ReadJsonDictionary() ([]byte, error) {
	file, err := jsonFile.ReadFile("resources/old-swedish-dictionary.json")

	return file, err
}
