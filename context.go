package mevrpc

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

const (
	UserIDCMetadataKey  = "X-API-USER"
	PlayerIDMetadataKey = "X-API-PLAYER"
)

var (
	errUnableExtractUserIDFromMetadata = func(err error) error {
		return fmt.Errorf("unable to extract user id from metadata: %w", err)
	}
	errUnableExtractPlayerIDFromMetadata = func(err error) error {
		return fmt.Errorf("unable to extract player id from metadata: %w", err)
	}
)

var (
	errUnableToParseMetaData       = errors.New("unable to parse metadata")
	errMetaDataContainsMalformedID = errors.New("metadata contains malformed uuid")
	errIDMissingFromMetaData       = errors.New("id missing from metadata")
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
		return uuid.Nil, errUnableExtractUserIDFromMetadata(errUnableToParseMetaData)
	}
	if len(md.Get(UserIDCMetadataKey)) == 0 {
		return uuid.Nil, errUnableExtractUserIDFromMetadata(errIDMissingFromMetaData)
	}
	client := md.Get(UserIDCMetadataKey)[0]
	user, err := uuid.FromString(client)
	if uuid.Equal(user, uuid.Nil) || err != nil {
		return uuid.Nil, errUnableExtractUserIDFromMetadata(errMetaDataContainsMalformedID)
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
		return uuid.Nil, errUnableExtractPlayerIDFromMetadata(errUnableToParseMetaData)
	}
	if len(md.Get(PlayerIDMetadataKey)) == 0 {
		return uuid.Nil, errUnableExtractPlayerIDFromMetadata(errIDMissingFromMetaData)
	}
	client := md.Get(PlayerIDMetadataKey)[0]
	player, err := uuid.FromString(client)
	if uuid.Equal(player, uuid.Nil) || err != nil {
		return uuid.Nil, errUnableExtractPlayerIDFromMetadata(errMetaDataContainsMalformedID)
	}
	return player, nil
}
