package utils

func GetWhere(data map[string]interface{}) (string, []interface{}) {
	var i = 0
	var command = ""
	var request []interface{}
	for index, value := range data {
		if i != 0 {
			command += " and "
		}
		command += index + " = ? "
		request = append(request, value)
		i++
	}
	return command, request
}
