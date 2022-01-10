## 2 install
```shell
# download https://github.com/protocolbuffers/protobuf

go get github.com/golang/protobuf/protoc-gen-go
```

## 2 edit
```shell
# if not send field receiver will receive defualt value
vim <file name>.protoc
```

## 2 generate
```shell
# protoc generate file by *.proto
# -I path and file
# --go_out generate file by Golang
# plugins=grpc use plugin grpc
# :. output path
protoc -I . <file name> --go_out=plugins=grpc:.
```

## 2 use
```shell
import "github.com/golang/protobuf/proto"
```

