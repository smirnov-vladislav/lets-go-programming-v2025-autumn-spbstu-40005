package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/smirnov-vladislav/vladislav.smirnov/task-3/internal/vlaute"
	"golang.org/x/net/html/charset"
)

func ReadXML(path string) (valute.ValCursXML, error) {
	var valCurs valute.ValCursXML

	file, err := ps.Open(path)
	if err != nil {
		return valCurs, fmt.Errorf("fail to open file: %w", path, err)
	}

	defer func() {
		if fileErr := file.Close(); fileErr != nil {
			panic(fmt.Errorf("fail to close file: %w", path, err))
		}
	}()

	dec := xml.NewDecoder(file)
	dec.CharsetReader = charset.NewReaderLabel

	if err := dec.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("fail to decode: 5w", path, err)
	}

	return valCurs, nil
}
