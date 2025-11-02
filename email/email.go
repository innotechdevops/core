package email

type Mail struct {
	From         string
	FromPassword string
	To           []string
	Subject      string
	HTML         []byte
}
