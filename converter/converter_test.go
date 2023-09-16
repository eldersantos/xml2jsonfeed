package converter_test

import (
	"fmt"
	"os"
	"testing"
	"xml2jsonfeed/converter"
)

func TestConverter(t *testing.T) {

	file, err := os.Open("./xml/input.xml")
	if err != nil {
		t.Error(err)
	}

	json, err := converter.Convert(file)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(json.String())
}
