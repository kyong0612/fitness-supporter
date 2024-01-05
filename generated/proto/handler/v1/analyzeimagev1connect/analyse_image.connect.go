// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/handler/v1/analyse_image.proto

package analyzeimagev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/kyong0612/fitness-supporter/generated/proto/handler/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "proto.handler.v1.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServiceAnalyzeImageProcedure is the fully-qualified name of the GreetService's AnalyzeImage
	// RPC.
	GreetServiceAnalyzeImageProcedure = "/proto.handler.v1.GreetService/AnalyzeImage"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	greetServiceServiceDescriptor            = v1.File_proto_handler_v1_analyse_image_proto.Services().ByName("GreetService")
	greetServiceAnalyzeImageMethodDescriptor = greetServiceServiceDescriptor.Methods().ByName("AnalyzeImage")
)

// GreetServiceClient is a client for the proto.handler.v1.GreetService service.
type GreetServiceClient interface {
	AnalyzeImage(context.Context, *connect.Request[v1.AnalyzeImageRequest]) (*connect.Response[v1.AnalyzeImageResponse], error)
}

// NewGreetServiceClient constructs a client for the proto.handler.v1.GreetService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greetServiceClient{
		analyzeImage: connect.NewClient[v1.AnalyzeImageRequest, v1.AnalyzeImageResponse](
			httpClient,
			baseURL+GreetServiceAnalyzeImageProcedure,
			connect.WithSchema(greetServiceAnalyzeImageMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// greetServiceClient implements GreetServiceClient.
type greetServiceClient struct {
	analyzeImage *connect.Client[v1.AnalyzeImageRequest, v1.AnalyzeImageResponse]
}

// AnalyzeImage calls proto.handler.v1.GreetService.AnalyzeImage.
func (c *greetServiceClient) AnalyzeImage(ctx context.Context, req *connect.Request[v1.AnalyzeImageRequest]) (*connect.Response[v1.AnalyzeImageResponse], error) {
	return c.analyzeImage.CallUnary(ctx, req)
}

// GreetServiceHandler is an implementation of the proto.handler.v1.GreetService service.
type GreetServiceHandler interface {
	AnalyzeImage(context.Context, *connect.Request[v1.AnalyzeImageRequest]) (*connect.Response[v1.AnalyzeImageResponse], error)
}

// NewGreetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreetServiceHandler(svc GreetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	greetServiceAnalyzeImageHandler := connect.NewUnaryHandler(
		GreetServiceAnalyzeImageProcedure,
		svc.AnalyzeImage,
		connect.WithSchema(greetServiceAnalyzeImageMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.handler.v1.GreetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreetServiceAnalyzeImageProcedure:
			greetServiceAnalyzeImageHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) AnalyzeImage(context.Context, *connect.Request[v1.AnalyzeImageRequest]) (*connect.Response[v1.AnalyzeImageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.handler.v1.GreetService.AnalyzeImage is not implemented"))
}