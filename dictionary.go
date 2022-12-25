package dictionary

import (
	"encoding/json"

	"github.com/stscoundrel/old-swedish-dictionary-go/internal/reader"
)

type DictionaryEntry struct {
	Headword          string   `json:"a"`
	PartOfSpeech      []string `json:"b"`
	GrammaticalAspect string   `json:"c"`
	Information       string   `json:"d"`
	Definitions       []string `json:"e"`
	AlternativeForms  []string `json:"f"`
}

func parseDictionary(bytes []byte) ([]DictionaryEntry, error) {
	var entries []DictionaryEntry

	err := json.Unmarshal(bytes, &entries)

	return entries, err
}

func GetDictionary() ([]DictionaryEntry, error) {
	rawDictionary, readError := reader.ReadJsonDictionary()

	if readError != nil {
		return []DictionaryEntry{}, readError
	}

	dictionary, parseError := parseDictionary(rawDictionary)

	if parseError != nil {
		return []DictionaryEntry{}, parseError
	}

	return dictionary, nil
}
