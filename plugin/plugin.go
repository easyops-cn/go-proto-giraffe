package plugin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/vanity"

	"github.com/easyops-cn/go-proto-giraffe"

	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

// generatedCodeVersion indicates a version of the generated code.
// It is incremented whenever an incompatibility between the generated code and
// the grpc package is introduced; the generated code references
// a constant, grpc.SupportPackageIsVersionN (where N is generatedCodeVersion).
const generatedCodeVersion = 4

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	ioPkgPath           = "io"
	contextPkgPath      = "context"
	giraffePkgPath      = "github.com/easyops-cn/giraffe-micro"
	giraffeProtoPkgPath = "github.com/easyops-cn/go-proto-giraffe"
)

func init() {
	generator.RegisterPlugin(new(giraffeMicro))
}

func NewPlugin(opts ...Option) generator.Plugin {
	g := &giraffeMicro{
		useGogoImport: false,
	}

	for _, o := range opts {
		o(g)
	}

	return g
}

// grpc is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for gRPC support.
type giraffeMicro struct {
	useGogoImport bool
	gen           *generator.Generator
	ioPkg         generator.Single
	contextPkg    generator.Single
	giraffePkg    generator.Single
}

type Option func(g *giraffeMicro)

var UseGogoImport = func(g *giraffeMicro) {
	g.useGogoImport = true
}

// Name returns the name of this plugin, "grpc".
func (g *giraffeMicro) Name() string {
	return "giraffe"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	ioPkg           string
	contextPkg      string
	giraffePkg      string
	giraffeProtoPkg string
)

// Init initializes the plugin.
func (g *giraffeMicro) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *giraffeMicro) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *giraffeMicro) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *giraffeMicro) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *giraffeMicro) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	if !g.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}

	ioPkg = string(g.gen.AddImport(ioPkgPath))
	contextPkg = string(g.gen.AddImport(contextPkgPath))
	giraffePkg = string(g.gen.AddImport(giraffePkgPath))
	giraffeProtoPkg = string(g.gen.AddImport(giraffeProtoPkgPath))

	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ = ", ioPkg, ".EOF")
	g.P("var _ ", contextPkg, ".Context")
	g.P("var _ ", giraffePkg, ".Client")
	g.P("var _ ", giraffeProtoPkg, ".Contract")
	g.P()

	// Assert version compatibility.
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the grpc package it is being compiled against.")
	g.P("const _ = ", giraffePkg, ".SupportPackageIsVersion", generatedCodeVersion, " // please upgrade the ", giraffePkg, " package")
	g.P()

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *giraffeMicro) GenerateImports(file *generator.FileDescriptor) {}

// reservedClientName records whether a client name is reserved on the client side.
var reservedClientName = map[string]bool{
	// TODO: do we need any in gRPC?
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

func methodDescVarName(method *pb.MethodDescriptorProto) string {
	return fmt.Sprintf("_%sMethodDesc", method.GetName())
}

func endpointName(method *pb.MethodDescriptorProto) string {
	return fmt.Sprintf("_%sEndpoint", method.GetName())
}

// deprecationComment is the standard comment added to deprecated
// messages, fields, enums, and enum values.
var deprecationComment = "// Deprecated: Do not use."

// generateService generates all the code for the named service.
func (g *giraffeMicro) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.

	serviceName := file.GetPackage()
	deprecated := service.GetOptions().GetDeprecated()

	g.P()
	g.P(fmt.Sprintf(`// Client is the client API for %s service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.`, serviceName))

	// Client interface.
	if deprecated {
		g.P("//")
		g.P(deprecationComment)
	}
	g.P("type Client interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateClientSignature(method))
	}
	g.P("}")
	g.P()

	// Client structure.
	g.P("type client struct {")
	g.P("c ", giraffePkg, ".Client")
	g.P("}")
	g.P()

	// NewClient factory.
	if deprecated {
		g.P(deprecationComment)
	}
	g.P("func NewClient (c ", giraffePkg, ".Client) Client {")
	g.P("return &client{")
	g.P("c: c,")
	g.P("}")
	g.P("}")
	g.P()

	// Client method implementations.
	for _, method := range service.Method {
		g.generateClientMethod(method)
	}

	// Server interface.
	serverType := "Service"
	g.P("// ", serverType, " is the server API for ", serviceName, " service.")
	if deprecated {
		g.P("//")
		g.P(deprecationComment)
	}
	g.P("type ", serverType, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateServerSignature(method))
	}
	g.P("}")
	g.P()

	for _, method := range service.Method {
		if !method.GetServerStreaming() && !method.GetClientStreaming() {
			g.generateServiceEndpoint(method)
		} else {
			g.generateServiceEndpoint(method)
		}
	}

	// Service registration.
	if deprecated {
		g.P(deprecationComment)
	}
	g.P("func RegisterService(s ", giraffePkg, ".Server, srv ", serverType, ") {")
	for _, method := range service.Method {
		if !method.GetServerStreaming() && !method.GetClientStreaming() {
			g.P("s.RegisterUnaryEndpoint(", methodDescVarName(method), ", ", endpointName(method), "(srv))")
		} else {
			g.P("s.RegisterStreamEndpoint(", methodDescVarName(method), ", ", endpointName(method), "(srv))")
		}
	}
	g.P("}")
	g.P()

	// Service descriptor.
	g.P("// Method Description")
	methodServiceName := fmt.Sprintf("%s.%s", serviceName, service.GetName())
	for _, method := range service.Method {
		if !method.GetClientStreaming() && !method.GetServerStreaming() {
			g.generateMethodDesc(methodServiceName, method)
		} else {
			g.generateStreamDesc(methodServiceName, method)
		}
	}
	g.P()
}

func getHttpRule(method *pb.MethodDescriptorProto) giraffeproto.HttpRule {
	if i, err := proto.GetExtension(method.GetOptions(), giraffeproto.E_Http); err == nil {
		return *(i.(*giraffeproto.HttpRule))
	}
	return giraffeproto.HttpRule{}
}

func getContract(method *pb.MethodDescriptorProto) giraffeproto.Contract {
	if i, err := proto.GetExtension(method.GetOptions(), giraffeproto.E_Contract); err == nil {
		return *(i.(*giraffeproto.Contract))
	}
	return giraffeproto.Contract{}
}

func (g *giraffeMicro) generateMethodDesc(serviceName string, method *pb.MethodDescriptorProto) {
	contract := getContract(method)
	httpRule := getHttpRule(method)
	g.P("var ", methodDescVarName(method), " = &", giraffePkg, ".MethodDesc{")
	g.P("Contract: &", giraffeProtoPkg, ".Contract{")
	g.P("Name: ", strconv.Quote(contract.GetName()), ",")
	g.P("Version: ", strconv.Quote(contract.GetVersion()), ",")
	g.P("},")
	g.P("ServiceName: ", strconv.Quote(serviceName), ",")
	g.P("MethodName: ", strconv.Quote(method.GetName()), ",")
	g.P("RequestType: (*", g.typeName(method.GetInputType()), ")(nil),")
	g.P("ResponseType: (*", g.typeName(method.GetOutputType()), ")(nil),")
	g.P("HttpRule: &", giraffeProtoPkg, ".HttpRule{")
	switch {
	case httpRule.GetGet() != "":
		g.P("Pattern: &", giraffeProtoPkg, ".HttpRule_Get{")
		g.P("Get: ", strconv.Quote(httpRule.GetGet()), ",")
		g.P("},")
	case httpRule.GetPost()	!= "":
		g.P("Pattern: &", giraffeProtoPkg, ".HttpRule_Post{")
		g.P("Post: ", strconv.Quote(httpRule.GetPost()), ",")
		g.P("},")
	case httpRule.GetPut() != "":
		g.P("Pattern: &", giraffeProtoPkg, ".HttpRule_Put{")
		g.P("Put: ", strconv.Quote(httpRule.GetPut()), ",")
		g.P("},")
	case httpRule.GetDelete() != "":
		g.P("Pattern: &", giraffeProtoPkg, ".HttpRule_Delete{")
		g.P("Delete: ", strconv.Quote(httpRule.GetDelete()), ",")
		g.P("},")
	case httpRule.GetPatch() != "":
		g.P("Pattern: &", giraffeProtoPkg, ".HttpRule_Patch{")
		g.P("Patch: ", strconv.Quote(httpRule.GetPatch()), ",")
		g.P("},")
	}
	g.P("Body: ", strconv.Quote(httpRule.GetBody()), ",")
	g.P("ResponseBody: ", strconv.Quote(httpRule.GetResponseBody()), ",")
	g.P("},")
	g.P("}")
	g.P()
}

func (g *giraffeMicro) generateStreamDesc(serviceName string, method *pb.MethodDescriptorProto) {
	contract := getContract(method)
	g.P("var ", methodDescVarName(method), " = &", giraffePkg, ".MethodDesc{")
	g.P("Contract: &", giraffeProtoPkg, ".Contract{")
	g.P("Name: ", strconv.Quote(contract.GetName()), ",")
	g.P("Version: ", strconv.Quote(contract.GetVersion()), ",")
	g.P("},")
	g.P("ServiceName: ", strconv.Quote(serviceName), ",")
	g.P("StreamName: ", strconv.Quote(method.GetName()), ",")
	g.P("ClientStreams: ", method.GetClientStreaming(), ",")
	g.P("ServerStreams: ", method.GetServerStreaming(), ",")
	g.P("}")
	g.P()
}

// generateClientSignature returns the client-side signature for a method.
func (g *giraffeMicro) generateClientSignature(method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}
	reqArg := ", in *" + g.typeName(method.GetInputType())
	if method.GetClientStreaming() {
		reqArg = ""
	}
	respName := "*" + g.typeName(method.GetOutputType())
	if method.GetServerStreaming() || method.GetClientStreaming() {
		respName = generator.CamelCase(origMethName) + "Client"
	}
	return fmt.Sprintf("%s(ctx %s.Context%s) (%s, error)", methName, contextPkg, reqArg, respName)
}

func (g *giraffeMicro) generateClientMethod(method *pb.MethodDescriptorProto) {
	methName := generator.CamelCase(method.GetName())
	inType := g.typeName(method.GetInputType())
	outType := g.typeName(method.GetOutputType())

	if method.GetOptions().GetDeprecated() {
		g.P(deprecationComment)
	}
	g.P("func (c *client) ", g.generateClientSignature(method), "{")
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		g.P("out := new(", outType, ")")
		// TODO: add call options
		g.P("err := c.c.Invoke(ctx, ", methodDescVarName(method), ", in, out)")
		g.P("if err != nil { return nil, err }")
		g.P("return out, nil")
		g.P("}")
		g.P()
		return
	}
	streamType := unexport(methName) + "Client"
	g.P("stream, err := c.c.NewStream(ctx, ", methodDescVarName(method), ")")
	g.P("if err != nil { return nil, err }")
	g.P("x := &", streamType, "{stream}")
	if !method.GetClientStreaming() {
		g.P("if err := x.ClientStream.SendMsg(in); err != nil { return nil, err }")
		g.P("if err := x.ClientStream.CloseSend(); err != nil { return nil, err }")
	}
	g.P("return x, nil")
	g.P("}")
	g.P()

	genSend := method.GetClientStreaming()
	genRecv := method.GetServerStreaming()
	genRecvAll := method.GetServerStreaming() && !method.GetClientStreaming()
	genCloseAndRecv := !method.GetServerStreaming()

	// Stream auxiliary types and methods.
	g.P("type ", methName, "Client interface {")
	if genSend {
		g.P("Send(*", inType, ") error")
	}
	if genRecv {
		g.P("Recv() (*", outType, ", error)")
	}
	if genRecvAll {
		g.P("RecvAll() ([]", outType, ", error)")
	}
	if genCloseAndRecv {
		g.P("CloseAndRecv() (*", outType, ", error)")
	}
	g.P(giraffePkg, ".ClientStream")
	g.P("}")
	g.P()

	g.P("type ", streamType, " struct {")
	g.P(giraffePkg, ".ClientStream")
	g.P("}")
	g.P()

	if genSend {
		g.P("func (x *", streamType, ") Send(m *", inType, ") error {")
		g.P("return x.ClientStream.SendMsg(m)")
		g.P("}")
		g.P()
	}
	if genRecv {
		g.P("func (x *", streamType, ") Recv() (*", outType, ", error) {")
		g.P("m := new(", outType, ")")
		g.P("if err := x.ClientStream.RecvMsg(m); err != nil { return nil, err }")
		g.P("return m, nil")
		g.P("}")
		g.P()
	}
	if genRecvAll {
		g.P("func (x *", streamType, ") RecvAll() ([]", outType, ", error) {")
		g.P("var resp []", outType)
		g.P("for {")
		g.P("m := new(", outType, ")")
		g.P("err := x.ClientStream.RecvMsg(m);")
		g.P("if err == io.EOF { break }")
		g.P("if err != nil { return nil, err }")
		g.P("resp = append(resp, *m)")
		g.P("}")
		g.P("return resp, nil")
		g.P("}")
		g.P()
	}
	if genCloseAndRecv {
		g.P("func (x *", streamType, ") CloseAndRecv() (*", outType, ", error) {")
		g.P("if err := x.ClientStream.CloseSend(); err != nil { return nil, err }")
		g.P("m := new(", outType, ")")
		g.P("if err := x.ClientStream.RecvMsg(m); err != nil { return nil, err }")
		g.P("return m, nil")
		g.P("}")
		g.P()
	}
}

// generateServerSignature returns the server-side signature for a method.
func (g *giraffeMicro) generateServerSignature(method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}

	var reqArgs []string
	ret := "error"
	reqArgs = append(reqArgs, contextPkg+".Context")
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		ret = "(*" + g.typeName(method.GetOutputType()) + ", error)"
	}
	if !method.GetClientStreaming() {
		reqArgs = append(reqArgs, "*"+g.typeName(method.GetInputType()))
	}
	if method.GetServerStreaming() || method.GetClientStreaming() {
		reqArgs = append(reqArgs, generator.CamelCase(origMethName)+"Stream")
	}

	return methName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func (g *giraffeMicro) generateServiceEndpoint(method *pb.MethodDescriptorProto) string {
	methodName := generator.CamelCase(method.GetName())
	endpointName := endpointName(method)
	inType := g.typeName(method.GetInputType())
	outType := g.typeName(method.GetOutputType())

	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		g.P("func ", endpointName, "(s Service) ", giraffePkg, ".UnaryEndpoint {")
		g.P("return func(ctx context.Context, req interface{}) (interface{}, error) {")
		g.P("return s.", methodName, "(ctx, req.(*", inType, "))")
		g.P("}")
		g.P("}")
		g.P()
		return endpointName
	}

	streamType := "_" + methodName + "Stream"
	g.P("func ", endpointName, "(s Service) ", giraffePkg, ".StreamEndpoint {")
	g.P("return func(ctx context.Context, stream ", giraffePkg, ".ServiceStream) error {")
	if !method.GetClientStreaming() {
		g.P("m := new(", inType, ")")
		g.P("if err := stream.RecvMsg(m); err != nil { return err }")
		g.P("return s.", methodName, "(ctx, m, &", streamType, "{stream})")
	} else {
		g.P("return s.", methodName, "(ctx, &", streamType, "{stream})")
	}
	g.P("}")
	g.P("}")
	g.P()

	genSend := method.GetServerStreaming()
	genSendAll := method.GetServerStreaming() && !method.GetClientStreaming()
	genSendAndClose := !method.GetServerStreaming()
	genRecv := method.GetClientStreaming()

	// Stream auxiliary types and methods.
	g.P("type ", methodName, "Stream interface {")
	if genSend {
		g.P("Send(*", outType, ") error")
	}
	if genSendAll {
		g.P("SendAll([]", outType, ") error")
	}
	if genSendAndClose {
		g.P("SendAndClose(*", outType, ") error")
	}
	if genRecv {
		g.P("Recv() (*", inType, ", error)")
	}
	g.P("}")
	g.P()

	g.P("type ", streamType, " struct {")
	g.P(giraffePkg, ".ServiceStream")
	g.P("}")
	g.P()

	if genSend {
		g.P("func (x *", streamType, ") Send(m *", outType, ") error {")
		g.P("return x.ServiceStream.SendMsg(m)")
		g.P("}")
		g.P()
	}
	if genSendAll {
		g.P("func (x *", streamType, ") SendAll(m []", outType, ") error {")
		g.P("for _, v := range m {")
		g.P("err := x.ServiceStream.SendMsg(v)")
		g.P("if err != nil { return err }")
		g.P("}")
		g.P("return nil")
		g.P("}")
		g.P()
	}
	if genSendAndClose {
		g.P("func (x *", streamType, ") SendAndClose(m *", outType, ") error {")
		g.P("return x.ServiceStream.SendMsg(m)")
		g.P("}")
		g.P()
	}
	if genRecv {
		g.P("func (x *", streamType, ") Recv() (*", inType, ", error) {")
		g.P("m := new(", inType, ")")
		g.P("if err := x.ServiceStream.RecvMsg(m); err != nil { return nil, err }")
		g.P("return m, nil")
		g.P("}")
		g.P()
	}

	return endpointName
}
