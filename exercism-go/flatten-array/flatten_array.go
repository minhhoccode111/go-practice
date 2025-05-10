package flatten

import "reflect"

func Flatten(nested any) []any {
	reflectNested := reflect.ValueOf(nested)

	// this function is guaranteed to be called with a slice or nil
	if reflectNested.Kind() != reflect.Slice {
		return nil
	}

	result := []any{}
	for i := 0; i < reflectNested.Len(); i++ {
		elem := reflectNested.Index(i).Interface()

		reflectElem := reflect.ValueOf(elem)

		if !reflectElem.IsValid() {
			continue
		}

		if reflectElem.Kind() != reflect.Slice {
			result = append(result, elem)
			continue
		}

		flattenedElem := Flatten(elem)
		result = append(result, flattenedElem...)
	}

	return result
}
