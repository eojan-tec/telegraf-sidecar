IMAGE=registry.cn-shanghai.aliyuncs.com/eojan/telegraf-sidecar:v1.0.0
run:
	go run main.go
build-image:
	# CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o strong main.go
	docker build -f Dockerfile -t ${IMAGE} .
test-image:
	docker run -v `pwd`:/workspace/code --rm  ${IMAGE}
push-image:
	docker push ${IMAGE}
gen-doc:
	swag init --parseDependency --parseInternal --parseDepth 1