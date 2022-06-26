package validate

// package validate

// go get -d github.com/envoyproxy/protoc-gen-validate
// make build

// 在proto文件中
// import "validate/validate.proto";

// server 使用拦截器
func server() {
	intercptor := func() {}
	if r, ok := req.(Validator); ok {
		if err := r.Validator(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	return handler(ctx, req)
}
