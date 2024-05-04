package mevrpc

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

var ErrUnableToParseMetaData = errors.New("unable to parse metadata")
var ErrMetaDataContainsMalformedUserID = errors.New("metadata contains malformed user uuid")

func UserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil, ErrUnableToParseMetaData
	}
	client := md.Get("X-API-USER")[0]
	user, err := uuid.FromString(client)
	if uuid.Equal(user, uuid.Nil) || err != nil {
		return uuid.Nil, ErrMetaDataContainsMalformedUserID
	}
	return user, nil
}

func PlayerIDFromContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil, ErrUnableToParseMetaData
	}
	client := md.Get("X-API-PLAYER")[0]
	player, err := uuid.FromString(client)
	if uuid.Equal(player, uuid.Nil) || err != nil {
		return uuid.Nil, ErrMetaDataContainsMalformedUserID
	}
	return player, nil
}
