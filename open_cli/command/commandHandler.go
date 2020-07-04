package command

type ApiCommand interface {
	Request() string
	Response(strArr []string) string
}
