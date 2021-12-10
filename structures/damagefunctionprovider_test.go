package structures

import (
	"fmt"
	"testing"

	"github.com/USACE/go-consequences/utils"
)

func Test_readJson(t *testing.T) {
	path := "/workspaces/go-consequences/data/DF.json"
	var c RawDFStruct
	err := utils.ReadJson(path, &c)

	if err != nil {
		t.Errorf("Unable to parse Json from file")
	}
}

func Test_IngestDFStore(t *testing.T) {
	path := "/workspaces/go-consequences/data/DF.json"
	d, err := IngestDFStore(path)
	if err != nil {
		t.Errorf("Unable to parse Json from file")
	}
	fmt.Println(d)
}
