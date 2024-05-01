package internal

type contextKeyType struct{}

var (
	ContextKey = contextKeyType(struct{}{})
)
