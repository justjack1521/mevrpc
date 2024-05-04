package mevrpc

import (
	"context"
	"google.golang.org/grpc"
)

var (
	UserIDContextKey   = "user_id"
	PlayerIDContextKey = "player_id"
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
