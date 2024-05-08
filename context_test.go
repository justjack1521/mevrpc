package mevrpc_test

import (
	"context"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserIDFromOutgoingContext(t *testing.T) {
	var user = uuid.NewV4()
	var out = mevrpc.NewOutgoingContext(context.Background(), user, uuid.Nil)
	assert.Equal(t, user, mevrpc.UserIDFromOutgoingContext(out))
}

func TestUserIDFromIncomingContext(t *testing.T) {
	var user = uuid.NewV4()
	var out = mevrpc.NewIncomingContext(context.Background(), user, uuid.Nil)
	assert.Equal(t, user, mevrpc.UserIDFromIncomingContext(out))
}
