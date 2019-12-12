export PATH=~/go/bin:$PATH
protoc --go_out=./ ./protobuf/judge.proto