package schema

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
