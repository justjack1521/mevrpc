package mevrpc_test

import (
	"context"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserIDFromContext(t *testing.T) {
	var user = uuid.NewV4()
	var out = mevrpc.NewIncomingContext(context.Background(), user, uuid.Nil)
	var in = mevrpc.NewIncomingContext(context.Background(), user, uuid.Nil)
	assert.Equal(t, user, mevrpc.UserIDFromContext(out))
	assert.Equal(t, user, mevrpc.UserIDFromContext(in))
}

func TestMustUserIDFromContext(t *testing.T) {
	var user = uuid.NewV4()
	var out = mevrpc.NewIncomingContext(context.Background(), user, uuid.Nil)
	var in = mevrpc.NewIncomingContext(context.Background(), user, uuid.Nil)

	o, err := mevrpc.MustUserIDFromContext(out)
	assert.Equal(t, o, user)
	assert.NoError(t, err)

	i, err := mevrpc.MustUserIDFromContext(in)
	assert.Equal(t, i, user)
	assert.NoError(t, err)
}

func TestPlayerIDFromContext(t *testing.T) {
	var player = uuid.NewV4()
	var out = mevrpc.NewIncomingContext(context.Background(), uuid.Nil, player)
	var in = mevrpc.NewIncomingContext(context.Background(), uuid.Nil, player)
	assert.Equal(t, player, mevrpc.PlayerIDFromContext(out))
	assert.Equal(t, player, mevrpc.PlayerIDFromContext(in))
}

func TestMustPlayerIDFromContext(t *testing.T) {
	var player = uuid.NewV4()
	var out = mevrpc.NewIncomingContext(context.Background(), uuid.Nil, player)
	var in = mevrpc.NewIncomingContext(context.Background(), uuid.Nil, player)

	o, err := mevrpc.MustPlayerIDFromContext(out)
	assert.Equal(t, o, player)
	assert.NoError(t, err)

	i, err := mevrpc.MustPlayerIDFromContext(in)
	assert.Equal(t, i, player)
	assert.NoError(t, err)
}
