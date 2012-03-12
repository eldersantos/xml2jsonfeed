package xml

import (
	"testing"
	"gokogiri/help"
)

func TestParseDocumentFragment(t *testing.T) {
	defer help.CheckXmlMemoryLeaks(t)

	doc, err := Parse(nil, DefaultEncodingBytes, nil, DefaultParseOption, DefaultEncodingBytes)
	if err != nil {
		t.Error("parsing error:", err.String())
		return
	}
	docFragment, err := doc.ParseFragment([]byte("<foo></foo><!-- comment here --><bar>fun</bar>"), nil, DefaultParseOption)
	if err != nil {
		t.Error(err.String())
		doc.Free()
		return
	}
	if (docFragment.Children.Length() != 3) {
		t.Error("the number of children from the fragment does not match")
	}
	doc.Free()
}

func TestSearchDocumentFragment(t *testing.T) {
	defer help.CheckXmlMemoryLeaks(t)

	doc, err := Parse([]byte("<moovweb><z/><s/></moovweb>"), DefaultEncodingBytes, nil, DefaultParseOption, DefaultEncodingBytes)
	if err != nil {
		t.Error("parsing error:", err.String())
		return
	}
	docFragment, err := doc.ParseFragment([]byte("<foo></foo><!-- comment here --><bar>fun</bar>"), nil, DefaultParseOption)
	if err != nil {
		t.Error(err.String())
		doc.Free()
		return
	}
	nodes, err := docFragment.Search(".//*")
	if err != nil {
		t.Error("fragment search has error")
		doc.Free()
		return
	}
	if len(nodes) != 2 {
		t.Error("the number of children from the fragment does not match")
	}
	nodes, err = docFragment.Search("//*")

	if err != nil {
		t.Error("fragment search has error")
		doc.Free()
		return
	}

	if len(nodes) != 3 {
		t.Error("the number of children from the fragment's document does not match")
	}

	doc.Free()
}

func TestSearchDocumentFragmentWithEmptyDoc(t *testing.T) {
	defer help.CheckXmlMemoryLeaks(t)

	doc, err := Parse(nil, DefaultEncodingBytes, nil, DefaultParseOption, DefaultEncodingBytes)
	if err != nil {
		t.Error("parsing error:", err.String())
		return
	}
	docFragment, err := doc.ParseFragment([]byte("<foo></foo><!-- comment here --><bar>fun</bar>"), nil, DefaultParseOption)
	if err != nil {
		t.Error(err.String())
		doc.Free()
		return
	}
	nodes, err := docFragment.Search(".//*")
	if err != nil {
		t.Error("fragment search has error")
		doc.Free()
		return
	}
	if len(nodes) != 2 {
		t.Error("the number of children from the fragment does not match")
	}
	nodes, err = docFragment.Search("//*")

	if err != nil {
		t.Error("fragment search has error")
		doc.Free()
		return
	}

	if len(nodes) != 0 {
		t.Error("the number of children from the fragment's document does not match")
	}

	doc.Free()
}