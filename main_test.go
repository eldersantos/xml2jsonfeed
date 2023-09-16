package xml2jsonfeed

import (
	"io/ioutil"
	"testing"
	"xml2jsonfeed/help"
)

func TestParseXml(t *testing.T) {
	input := "<foo></foo>"
	expected := `<?xml version="1.0" encoding="utf-8"?>
<foo/>
`
	doc, err := ParseXml([]byte(input))
	if err != nil {
		t.Error("Parsing has error:", err)
		return
	}

	if doc.String() != expected {
		t.Error("the output of the xml doc does not match the expected")
	}

	expected = `<?xml version="1.0" encoding="utf-8"?>
<foo>
  <bar/>
</foo>
`
	doc.Root().AddChild("<bar/>")
	if doc.String() != expected {
		t.Error("the output of the xml doc does not match the expected")
	}
	doc.Free()
	CheckXmlMemoryLeaks(t)
}

func CheckXmlMemoryLeaks(t *testing.T) {
	// LibxmlCleanUpParser() should only be called once during the lifetime of the
	// program, but because there's no way to know when the last test of the suite
	// runs in go, we can't accurately call it strictly once, so just avoid calling
	// it for now because it's known to cause crashes if called multiple times.
	//help.LibxmlCleanUpParser()

	if !help.LibxmlCheckMemoryLeak() {
		t.Errorf("Memory leaks: %d!!!", help.LibxmlGetMemoryAllocation())
		help.LibxmlReportMemoryLeak()
	}
}

func TestLoadAllNodes(t *testing.T) {
	input, err := ioutil.ReadFile("./xml/tests/document/rss/input.txt")

	if err != nil {
		t.Error(err)
	}

	res, err := ParseXml(input)
	if err != nil {
		t.Error(err)
	}
	root := res.Root()

	s := root.String()
	t.Log(s)
	f, err := root.Search("channel")
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(f); i++ {
		t.Log(f[i].Name())
	}
}
