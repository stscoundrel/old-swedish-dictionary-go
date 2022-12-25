package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionaryHasExpectedAmountOfEntries(t *testing.T) {
	result, _ := GetDictionary()

	assert.Equal(t, 41744, len(result))
}

func TestDictionaryHasExpectedEntries(t *testing.T) {
	result, _ := GetDictionary()

	expected1 := DictionaryEntry{
		Headword: "af bränna",
		PartOfSpeech: []string{
			"vb",
		},
		GrammaticalAspect: "v.",
		Information:       "",
		Definitions: []string{
			"afbränna, genom eld förstöra. hans trähws the af brendhe  RK 2: 2757 . ib 1511. halff stadhen är affbrändh  BSH 5: 132 (  1506) . Jfr bränna af.",
		},
		AlternativeForms: []string{},
	}

	expected2 := DictionaryEntry{
		Headword: "alder daghar",
		PartOfSpeech: []string{
			"nn",
		},
		GrammaticalAspect: "pl.",
		Information:       "",
		Definitions: []string{
			"ålderdom.  \" tiill aller da[gha] \" MD (S) 242 . oppa sina aldher dagha  Lg 3: 650 .",
		},
		AlternativeForms: []string{
			"aller- )",
		},
	}

	assert.Equal(t, expected1, result[100])
	assert.Equal(t, expected2, result[666])
}
