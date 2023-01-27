APP_NAME=scale_maker
APP_VERSION=0.0.1
NUC_DOCKER_REGISTRY=nuc.lliu.ca

build:
	go build  ./main.go
doc:
	swag init
run:
	go run ./main.go
test:
	go test `go list ./... | grep -v /docs`
testv:
	go test `go list ./... | grep -v /docs` -v -cover
cov:
	go test `go list ./... | grep -v /docs` -coverprofile cp.out
	go tool cover -html=cp.out
rpm:
	rm -rf ./releases/
	mkdir -p ./releases/
	nfpm pkg -f ./ops/nfpm.yaml --packager rpm --target ./releases/${APP_NAME}-${APP_VERSION}.rpm
docker_build:
	docker build -t ${APP_NAME} .
docker_tag:
	docker tag $(APP_NAME) $(NUC_DOCKER_REGISTRY)/$(APP_NAME):${APP_VERSION}
docker_nuc_push:
	docker push $(NUC_DOCKER_REGISTRY)/$(APP_NAME):${APP_VERSION}
localtest:
	docker-compose up
localtest_down:
	docker-compose down