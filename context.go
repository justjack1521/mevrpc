package mevrpc

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

const (
	UserIDCMetadataKey  = "X-API-USER"
	PlayerIDMetadataKey = "X-API-PLAYER"
)

var (
	errUnableToParseMetaData             = errors.New("unable to parse metadata")
	errUserIDMissingFromMetaData         = errors.New("user id missing from metadata")
	errMetaDataContainsMalformedUserID   = errors.New("metadata contains malformed user uuid")
	errMetaDataContainsMalformedPlayerID = errors.New("metadata contains malformed player uuid")
	errPlayerIDMissingFromMetaData       = errors.New("player id missing from metadata")
)

func NewOutgoingContext(ctx context.Context, user uuid.UUID, player uuid.UUID) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		UserIDCMetadataKey:  user.String(),
		PlayerIDMetadataKey: player.String(),
	}))
}

func UserIDFromContext(ctx context.Context) uuid.UUID {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil
	}
	if len(md.Get(UserIDCMetadataKey)) == 0 {
		return uuid.Nil
	}
	val := md.Get(UserIDCMetadataKey)[0]
	return uuid.FromStringOrNil(val)
}

func MustUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil, errUnableToParseMetaData
	}
	if len(md.Get(UserIDCMetadataKey)) == 0 {
		return uuid.Nil, errUserIDMissingFromMetaData
	}
	client := md.Get(UserIDCMetadataKey)[0]
	user, err := uuid.FromString(client)
	if uuid.Equal(user, uuid.Nil) || err != nil {
		return uuid.Nil, errMetaDataContainsMalformedUserID
	}
	return user, nil
}

func PlayerIDFromContext(ctx context.Context) uuid.UUID {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil
	}
	if len(md.Get(PlayerIDMetadataKey)) == 0 {
		return uuid.Nil
	}
	val := md.Get(PlayerIDMetadataKey)[0]
	return uuid.FromStringOrNil(val)
}

func MustPlayerIDFromContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil, errUnableToParseMetaData
	}
	if len(md.Get(PlayerIDMetadataKey)) == 0 {
		return uuid.Nil, errPlayerIDMissingFromMetaData
	}
	client := md.Get(PlayerIDMetadataKey)[0]
	player, err := uuid.FromString(client)
	if uuid.Equal(player, uuid.Nil) || err != nil {
		return uuid.Nil, errMetaDataContainsMalformedPlayerID
	}
	return player, nil
}
