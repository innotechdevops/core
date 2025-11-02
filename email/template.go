package email

type Template interface {
	Text(body map[string]interface{}) []byte
}
