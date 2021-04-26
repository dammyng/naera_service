package billsgrpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


var(
	InternalError = status.Errorf(codes.Internal, "Internal service error")
)