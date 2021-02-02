package echo

import context "context"

// Server struct
type Server struct{}

// Echo responds with the request message
func (s *Server) Echo(ctx context.Context, request *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{
		Message: request.Message,
	}, nil
}

func (s *Server) mustEmbedUnimplementedEchoServiceServer() {}
