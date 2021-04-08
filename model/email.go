package model

type EmailTo struct {
	UserId  int64
	Name    string
	Address string
}

type EmailFrom struct {
	Name    string
	Address string
}

type EmailBody interface{}

type EmailData struct {
	Category []string
	Subject  string
	From     EmailFrom
	To       EmailTo
	Body     EmailBody
}
