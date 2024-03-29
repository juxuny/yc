package services

import "strings"

var reactTsDataTypeMapper = map[string]string{
	"dt.ID":          "string | number",
	"uint64":         "number",
	"uint32":         "number",
	"float":          "number",
	"double":         "number",
	"dt.NullInt64":   "string | number",
	"dt.NullInt32":   "string | number",
	"dt.NullBool":    "string",
	"dt.NullFloat64": "number",
}

func ConvertProtoTypeToReactTsDataType(dataType string) (finalType string, nullable bool) {
	finalType, b := reactTsDataTypeMapper[dataType]
	if b {
		if strings.Contains(dataType, "dt") {
			nullable = false
		}
	} else {
		finalType = dataType
		nullable = true
	}
	return
}

var cSharpDataTypeMapper = map[string]string{
	"int64":  "long",
	"int32":  "int",
	"uint32": "uint",
	"uint64": "ulong",
}

func ConvertProtoTypeToCSharpDataType(dataType string) (finalType string) {
	finalType, b := cSharpDataTypeMapper[dataType]
	if !b {
		return dataType
	}
	return
}

var goDataTypeMapper = map[string]string{
	"double": "float64",
	"float":  "float32",
}

func ConvertProtoTypeToGoDataType(dataType string) string {
	finalType, b := goDataTypeMapper[dataType]
	if !b {
		return dataType
	}
	return finalType
}
