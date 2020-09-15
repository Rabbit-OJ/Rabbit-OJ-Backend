export PATH=~/go/bin:$PATH
protoc --go_out=./ ./models/protobuf/case.proto
protoc --go_out=./ ./services/judger/protobuf/judger.proto