cd tfapi
protoc -I . --go_out=. --go-triple_out=. telegraf.proto
protoc-go-inject-tag -input="telegraf.pb.go"
sed -i "" -e "s/,omitempty//g" telegraf.pb.go
cd ..
