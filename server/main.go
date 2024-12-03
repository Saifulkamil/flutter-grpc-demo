

type GreeterServer struct {
    helloworld.UnimplementedGreeterServer
}




func (s *GreeterServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
    return &helloworld.HelloReply{
        Message: "Hello " + req.GetName(),
    }, nil
}

func (s *GreeterServer) SayHelloAgain(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
    return &helloworld.HelloReply{
        Message: "Hello again " + req.GetName(),
    }, nil
}
