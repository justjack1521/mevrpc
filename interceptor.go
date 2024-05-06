package mevrpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	errFailedExtractIdentity = func(err error) error {
		return fmt.Errorf("failed to extract identity from service interceptor: %w", err)
	}
	errFailedCopyIdentity = func(err error) error {
		return fmt.Errorf("failed to copy identity from client interceptor: %w", err)
	}
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

func IdentityCopyInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	user, err := MustUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	player, err := MustPlayerIDFromContext(ctx)
	if err != nil {
		return err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, UserIDCMetadataKey, user.String(), PlayerIDMetadataKey, player.String())
	return invoker(ctx, method, req, reply, cc, opts...)
}
