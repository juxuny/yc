package services

import "strings"

var reactTsDataTypeMapper = map[string]string {
	"dt.ID": "string | number",
	"uint64": "number",
	"uint32": "number",
	"float": "number",
	"double": "number",
	"dt.NullInt64": "string | number",
	"dt.NullInt32": "string | number",
	"dt.NullBool": "string",
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
