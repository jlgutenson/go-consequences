package structureprovider

import (
	"fmt"

	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/structures"
	"github.com/dewberry/gdal"
)

func StructureSchema() []string {
	s := make([]string, 10)
	s[0] = "fd_id"
	s[1] = "cbfips"
	s[2] = "x"
	s[3] = "y"
	s[4] = "st_damcat"
	s[5] = "occtype"
	s[6] = "val_struct"
	s[7] = "val_cont"
	s[8] = "found_ht"
	s[9] = "found_type"
	return s
}
func featuretoStructure(f *gdal.Feature, m map[string]structures.OccupancyTypeStochastic, defaultOcctype structures.OccupancyTypeStochastic, idxs []int) (structures.StructureStochastic, error) {
	s := structures.StructureStochastic{}
	s.Name = fmt.Sprintf("%v", f.FieldAsInteger(idxs[0]))
	OccTypeName := f.FieldAsString(idxs[5])
	var occtype = defaultOcctype
	//dont have access to foundation type in the structure schema yet.
	if idxs[9] > 0 {
		if otf, okf := m[OccTypeName+"-"+f.FieldAsString(idxs[9])]; okf {
			occtype = otf
		}
	} else {
		if ot, ok := m[OccTypeName]; ok {
			occtype = ot
		} else {
			occtype = defaultOcctype
			msg := "Using default " + OccTypeName + " not found"
			fmt.Println(msg)
			//return s, errors.New(msg)
		}
	}

	s.OccType = occtype
	s.X = f.FieldAsFloat64(idxs[2])
	s.Y = f.FieldAsFloat64(idxs[3])
	s.DamCat = f.FieldAsString(idxs[4])
	s.StructVal = consequences.ParameterValue{Value: f.FieldAsFloat64(idxs[6])}
	s.ContVal = consequences.ParameterValue{Value: f.FieldAsFloat64(idxs[7])}
	s.FoundHt = consequences.ParameterValue{Value: f.FieldAsFloat64(idxs[8])}

	return s, nil
}
