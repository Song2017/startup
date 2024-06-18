package pkg

import "reflect"

// Map
func JsonMerge(dst, src map[string]interface{}) map[string]interface{} {
	return MapMerge(dst, src, 0)
}

func MapMerge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > 10 {
		return dst
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := MapMapify(srcVal)
			dstMap, dstMapOk := MapMapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = MapMerge(dstMap, srcMap, depth+1)
			}
		}

		dst[key] = srcVal
	}

	return dst
}

func MapMapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}

	return map[string]interface{}{}, false
}

func GetMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m)) // create a slice of keys with initial capacity

	for k := range m {
		keys = append(keys, k) // append the key to the slice
	}

	return keys
}

func GetMapValues[T any](myMap map[string]T) []T {
	values := make([]T, 0, len(myMap))

	// 遍历map，并将值追加到切片中
	for _, value := range myMap {
		values = append(values, value)
	}

	return values
}
