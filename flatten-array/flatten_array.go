package flatten

func Flatten(array interface{}) []interface{} {
	output := []interface{}{}

	for _, value := range array.([]interface{}) {
		switch value.(type) {
		case []interface{}:
			output = append(output, Flatten(value)...)
		case nil:
		default:
			output = append(output, value)
		}
	}
	return output
}
