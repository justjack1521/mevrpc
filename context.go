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
	errUnableExtractUserIDFromContext = func(err error) error {
		return fmt.Errorf("unable to extract user id from context: %w", err)
	}
	errUnableExtractPlayerIDFromContext = func(err error) error {
		return fmt.Errorf("unable to extract player id from context: %w", err)
	}
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

func UserIDFromContext(ctx context.Context) uuid.UUID {
	out, ok := metadata.FromOutgoingContext(ctx)
	if ok == true {
		user, err := userIDFromMetadata(out)
		if err == nil {
			return user
		}
	}
	in, ok := metadata.FromIncomingContext(ctx)
	if ok == true {
		user, err := userIDFromMetadata(in)
		if err == nil {
			return user
		}
	}
	return uuid.Nil
}

func PlayerIDFromContext(ctx context.Context) uuid.UUID {
	out, ok := metadata.FromOutgoingContext(ctx)
	if ok == true {
		player, err := playerIDFromMetadata(out)
		if err == nil {
			return player
		}
	}
	in, ok := metadata.FromIncomingContext(ctx)
	if ok == true {
		player, err := playerIDFromMetadata(in)
		if err == nil {
			return player
		}
	}
	return uuid.Nil
}

func MustUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	out, ok := metadata.FromOutgoingContext(ctx)
	if ok == true {
		user, err := userIDFromMetadata(out)
		if err == nil {
			return user, nil
		}
	}
	in, ok := metadata.FromIncomingContext(ctx)
	if ok == true {
		user, err := userIDFromMetadata(in)
		if err == nil {
			return user, nil
		}
	}
	return uuid.Nil, errUnableExtractUserIDFromContext(errUnableToParseMetaData)
}

func MustPlayerIDFromContext(ctx context.Context) (uuid.UUID, error) {
	out, ok := metadata.FromOutgoingContext(ctx)
	if ok == true {
		player, err := playerIDFromMetadata(out)
		if err == nil {
			return player, nil
		}
	}
	in, ok := metadata.FromIncomingContext(ctx)
	if ok == true {
		player, err := playerIDFromMetadata(in)
		if err == nil {
			return player, nil
		}
	}
	return uuid.Nil, errUnableExtractPlayerIDFromContext(errUnableToParseMetaData)
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
