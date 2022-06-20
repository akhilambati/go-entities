package goentities

import (
	"reflect"
	"strings"
)

func castStructs(field, value *reflect.Value) {
	if (field.Type().Kind() != reflect.Struct) || (value.Type().Kind() != reflect.Struct) {
		panic("Input and Output Type must be struct")
	}

	outputType := field.Type()

	for i := 0; i < field.NumField(); i++ {
		innerField := field.Field(i)
		key := strings.Split(outputType.Field(i).Tag.Get("entity"), ",")[0]
		if key == "" {
			continue
		}
		innerValue := value.FieldByName(key)

		// If the Field is not present in input then continue
		if innerValue.Kind() == reflect.Invalid {
			continue
		}

		castField(&innerField, &innerValue)
	}
}

func castSliceofStructs(field, value *reflect.Value) interface{} {
	if field.Type().Kind() != reflect.Struct || value.Type().Kind() != reflect.Slice {
		panic("Unsupported Input and output types")
	}

	outputType := reflect.SliceOf(field.Type())
	output := reflect.MakeSlice(outputType, value.Len(), value.Len())

	castSlices(&output, value)

	return output.Interface()
}