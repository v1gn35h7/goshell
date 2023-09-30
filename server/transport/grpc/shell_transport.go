package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/pb"
)

// Shell serviec contracts
type GetScriptRequest struct {
	AgentId         string `protobuf:"bytes,1,opt,name=AgentId,proto3" json:"AgentId,omitempty"`
	HostName        string `protobuf:"bytes,1,opt,name=HostName,proto3" json:"HostName,omitempty"`
	Platform        string `protobuf:"bytes,2,opt,name=Platform,proto3" json:"Platform,omitempty"`
	OperatingSystem string `protobuf:"bytes,3,opt,name=OperatingSystem,proto3" json:"OperatingSystem,omitempty"`
	Architecture    string `protobuf:"bytes,4,opt,name=Architecture,proto3" json:"Architecture,omitempty"`
	Frequency       string `protobuf:"bytes,4,opt,name=Frequency,proto3" json:"Frequency,omitempty"`
}

type GetScriptResponse struct {
	Scripts []*goshell.Script `json:"scripts"`
}

type ShellFragmentRquest struct {
	Outputs string `protobuf:"bytes,4,opt,name=Output,proto3" json:"Output,omitempty"`
}

type FragmentResponse struct {
	Awknowledgement int32 `protobuf:"varint,1,opt,name=Awknowledgement,proto3" json:"Awknowledgement,omitempty"`
}

// Grpc server
type grpcServer struct {
	getScripts   grpctransport.Handler
	sendFragment grpctransport.Handler
	pb.UnimplementedShellServiceServer
}

func NewServer(endpoints grpcEndpoints) *grpcServer {

	return &grpcServer{
		getScripts: grpctransport.NewServer(
			endpoints.GetScriptEndpoint,
			decodeGetScriptsRequest,
			encodeGetScriptsResponse,
		),
		sendFragment: grpctransport.NewServer(
			endpoints.SendFragmentEndpoint,
			decodeSendFragmentRequest,
			encodeFragmentResponse,
		),
	}

}

func (s *grpcServer) GetScripts(ctx context.Context, req *pb.ShellRequest) (*pb.ShellResponse, error) {
	_, rep, err := s.getScripts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ShellResponse), nil
}

func (s *grpcServer) SendFragment(ctx context.Context, req *pb.ShellFragmentRquest) (*pb.FragmentResponse, error) {
	_, rep, err := s.sendFragment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FragmentResponse), nil
}

// decodeGetScriptsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sum request to a user-domain sum request. Primarily useful in a server.
func decodeGetScriptsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ShellRequest)
	return req, nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func encodeGetScriptsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetScriptResponse)
	scripts := make([]*pb.ShellScript, 0)
	for _, v := range resp.Scripts {
		scripts = append(scripts, &pb.ShellScript{
			Script:    v.Script,
			Args:      "tt",
			Type:      "S",
			Frequency: v.Frequency,
			Id:        v.Id,
		})
	}
	return &pb.ShellResponse{Scripts: scripts}, nil
}

// decodeGetScriptsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sum request to a user-domain sum request. Primarily useful in a server.
func decodeSendFragmentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ShellFragmentRquest)
	return req, nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func encodeFragmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(FragmentResponse)
	return &pb.FragmentResponse{Awknowledgement: resp.Awknowledgement}, nil
}
