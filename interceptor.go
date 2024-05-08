package mevrpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

var (
	errFailedExtractIdentity = func(err error) error {
		return fmt.Errorf("failed to extract identity from service interceptor: %w", err)
	}
	errFailedCopyIdentity = func(err error) error {
		return fmt.Errorf("failed to copy identity from client interceptor: %w", err)
	}
)

func IdentityExtractionUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	_, err = MustUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, errFailedExtractIdentity(err)
	}
	_, err = MustPlayerIDFromIncomingContext(ctx)
	if err != nil {
		return nil, errFailedExtractIdentity(err)
	}
	resp, err = handler(ctx, req)
	return
}

func IdentityCopyUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	user, err := MustUserIDFromIncomingContext(ctx)
	if err != nil {
		return errFailedCopyIdentity(err)
	}
	player, err := MustPlayerIDFromIncomingContext(ctx)
	if err != nil {
		return errFailedCopyIdentity(err)
	}
	return invoker(NewOutgoingContext(ctx, user, player), method, req, reply, cc, opts...)
}
