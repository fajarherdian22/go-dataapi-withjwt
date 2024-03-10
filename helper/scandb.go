package helper

import (
	"reflect"
)

func ScanDB(isStruct interface{}) []interface{} {
	datamodel := isStruct
	s := reflect.ValueOf(&datamodel).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	return columns
}
