package schema

import "fmt"

var OrderBy = map[string]bool{
	"asc":  true,
	"desc": true,
}

func In[T any](source []T) ([]any, string) {
	whereIn := ""
	args := []any{}
	for i, id := range source {
		args = append(args, id)
		if i > 0 {
			whereIn += ","
		}
		whereIn += "?"
	}
	return args, whereIn
}

func Join[T any](source []T, sep string) string {
	output := ""
	for i, v := range source {
		str := fmt.Sprint(v)
		if i > 0 {
			output += sep
		}
		output += str
	}
	return output
}
