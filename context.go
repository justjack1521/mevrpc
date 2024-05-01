package mevrpc

import (
	"context"
	"github.com/justjack1521/mevrpc/internal"
	uuid "github.com/satori/go.uuid"
)

type Context struct {
	UserID   uuid.UUID
	PlayerID uuid.UUID
}

func NewContext(ctx context.Context, cxn *Context) context.Context {
	return context.WithValue(ctx, internal.ContextKey, cxn)
}

func FromContext(ctx context.Context) *Context {
	if ctx == nil {
		return nil
	}
	h, _ := ctx.Value(internal.ContextKey).(*Context)
	return h
}

func UserIDFromContext(ctx context.Context) uuid.UUID {
	var cxn = FromContext(ctx)
	if cxn == nil {
		return uuid.Nil
	}
	return cxn.UserID
}

func PlayerIDFromContext(ctx context.Context) uuid.UUID {
	var cxn = FromContext(ctx)
	if cxn == nil {
		return uuid.Nil
	}
	return cxn.PlayerID
}
