# Old Swedish Dictionary

Old Swedish Dictionary for Golang. The dictionary consists of 40 000+ Old Swedish words with Swedish translations.

Based on K.F. Söderwall's Medieval Swedish Dictionary

### Install

`go get github.com/stscoundrel/old-swedish-dictionary-go`

### Usage

The library exposes one function for getting the whole dictionary dataset.

```go
package main

import (
    "fmt"

    dictionary "github.com/stscoundrel/old-swedish-dictionary-go"
)

func main() {
  result, err := dictionary.GetDictionary()

  if err != nil {
    fmt.Println(err)
  }
  
  // Contains 41 000+ DictionaryEntries.
  for _, entry := range result {
    fmt.Println(entry.Headword)
  }
}
```

The entries are structs of:

```go
type DictionaryEntry struct {
  Headword          string
  PartOfSpeech      string
  GrammaticalAspect string
  Definitions       []string
  AlternativeForms  []string
}

```

### About "Dictionary of Old Swedish"

_"Ordbok Öfver svenska medeltids-språket"_ dictionary was published in late 1884—1918 by K.F. Söderwall. Additional supplement to it was published in 1953—1973.

Old Swedish developed from Old East Norse, the eastern dialect of Old Norse, at the end of the Viking Age. Early Old Swedish was spoken from about 1225 until about 1375, and Late Old Swedish was spoken from about 1375 until about 1526.

The original material is licenced under [Creative Commons International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/), made available by University of Gothenburg. The source code for this library is under MIT licence.

- https://spraakbanken.gu.se/en/resources/soederwall
- https://spraakbanken.gu.se/en/resources/soederwall-supp

