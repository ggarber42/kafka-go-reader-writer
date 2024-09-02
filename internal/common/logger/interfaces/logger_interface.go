package ilogger

type ILogger interface {
	Info(msg string, info ...any)
	Error(msg string, info ...any)
}