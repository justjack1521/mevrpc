package mevrpc

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

const (
	UserIDMetadataKey   = "X-API-USER"
	PlayerIDMetadataKey = "X-API-PLAYER"
)

var (
	errUnableExtractUserIDFromOutgoingContext = func(err error) error {
		return fmt.Errorf("unable to extract user id from outgoing context: %w", err)
	}
	errUnableExtractUserIDFromIncomingContext = func(err error) error {
		return fmt.Errorf("unable to extract user id from incoming context: %w", err)
	}
	errUnableExtractPlayerIDFromOutgoingContext = func(err error) error {
		return fmt.Errorf("unable to extract player id from outgoing context: %w", err)
	}
	errUnableExtractPlayerIDFromIncomingContext = func(err error) error {
		return fmt.Errorf("unable to extract player id from incoming context: %w", err)
	}
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

func NewIncomingContext(ctx context.Context, user uuid.UUID, player uuid.UUID) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
		UserIDMetadataKey:   user.String(),
		PlayerIDMetadataKey: player.String(),
	}))
}

func NewOutgoingContext(ctx context.Context, user uuid.UUID, player uuid.UUID) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		UserIDMetadataKey:   user.String(),
		PlayerIDMetadataKey: player.String(),
	}))
}

func UserIDFromOutgoingContext(ctx context.Context) uuid.UUID {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok == false {
		return uuid.Nil
	}
	if len(md.Get(UserIDMetadataKey)) == 0 {
		return uuid.Nil
	}
	val := md.Get(UserIDMetadataKey)[0]
	return uuid.FromStringOrNil(val)
}

func UserIDFromIncomingContext(ctx context.Context) uuid.UUID {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil
	}
	if len(md.Get(UserIDMetadataKey)) == 0 {
		return uuid.Nil
	}
	val := md.Get(UserIDMetadataKey)[0]
	return uuid.FromStringOrNil(val)
}

func MustUserIDFromOutgoingContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok == false {
		return uuid.Nil, errUnableExtractUserIDFromOutgoingContext(errUnableToParseMetaData)
	}
	value, err := userIDFromMetadata(md)
	if err != nil {
		return uuid.Nil, errUnableExtractUserIDFromOutgoingContext(err)
	}
	return value, nil
}

func MustUserIDFromIncomingContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil, errUnableExtractUserIDFromIncomingContext(errUnableToParseMetaData)
	}
	value, err := userIDFromMetadata(md)
	if err != nil {
		return uuid.Nil, errUnableExtractUserIDFromIncomingContext(err)
	}
	return value, nil
}
func PlayerIDFromOutgoingContext(ctx context.Context) uuid.UUID {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok == false {
		return uuid.Nil
	}
	if len(md.Get(PlayerIDMetadataKey)) == 0 {
		return uuid.Nil
	}
	val := md.Get(PlayerIDMetadataKey)[0]
	return uuid.FromStringOrNil(val)
}

func PlayerIDFromIncomingContext(ctx context.Context) uuid.UUID {
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

func MustPlayerIDFromOutgoingContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok == false {
		return uuid.Nil, errUnableExtractPlayerIDFromOutgoingContext(errUnableToParseMetaData)
	}
	value, err := playerIDFromMetadata(md)
	if err != nil {
		return uuid.Nil, errUnableExtractPlayerIDFromOutgoingContext(err)
	}
	return value, nil
}

func MustPlayerIDFromIncomingContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return uuid.Nil, errUnableExtractPlayerIDFromIncomingContext(errUnableToParseMetaData)
	}
	value, err := playerIDFromMetadata(md)
	if err != nil {
		return uuid.Nil, errUnableExtractPlayerIDFromIncomingContext(err)
	}
	return value, nil
}

func userIDFromMetadata(meta metadata.MD) (uuid.UUID, error) {
	if len(meta.Get(UserIDMetadataKey)) == 0 {
		return uuid.Nil, errUnableExtractUserIDFromMetadata(errIDMissingFromMetaData)
	}
	client := meta.Get(UserIDMetadataKey)[0]
	user, err := uuid.FromString(client)
	if uuid.Equal(user, uuid.Nil) || err != nil {
		return uuid.Nil, errUnableExtractUserIDFromMetadata(errMetaDataContainsMalformedID)
	}
	return user, nil
}

func playerIDFromMetadata(meta metadata.MD) (uuid.UUID, error) {
	if len(meta.Get(PlayerIDMetadataKey)) == 0 {
		return uuid.Nil, errUnableExtractPlayerIDFromMetadata(errIDMissingFromMetaData)
	}
	client := meta.Get(PlayerIDMetadataKey)[0]
	player, err := uuid.FromString(client)
	if uuid.Equal(player, uuid.Nil) || err != nil {
		return uuid.Nil, errUnableExtractPlayerIDFromMetadata(errMetaDataContainsMalformedID)
	}
	return player, nil
}
