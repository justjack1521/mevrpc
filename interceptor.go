package mevrpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func IdentityExtractionInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	_, err = MustUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, err = MustPlayerIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, err = handler(ctx, req)
	return
}

func IdentityCopyInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	user, err := MustUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	player, err := MustPlayerIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, UserIDCMetadataKey, user.String(), PlayerIDMetadataKey, player.String())
	resp, err = handler(ctx, req)
	return
}
