package flatten

func Flatten(array interface{}) []interface{} {
	output := []interface{}{}
	return flatten(array, output)
}

func flatten(input interface{}, output []interface{}) []interface{} {
	if array, ok := input.([]interface{}); ok {
		for _, element := range array {
			output = flatten(element, output)
		}
	} else if input != nil {
		output = append(output, input)
	}
	return output
}
