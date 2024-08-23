package ctxkey

// ctxKey Key type for context.context
type ctxKey int

const (
	// UserID User UUID key
	UserID ctxKey = iota
)
