SERVICE_FILE_NAME?=pRpc1
SERVICE_NAME?=pRpc1
SERVICE_PREFIX?=rpc1

#每次发版本需要维护版本
SERVICE_VERSION?=v1.0.0

service-name:
	@echo ${SERVICE_NAME} ${SERVICE_VERSION}

# go快捷命令
tidy:
	go mod tidy

# goctl 快捷命令
goctl-api:
	# goctl api plugin -p goctl-go-compact -api buybuybuyService.api -dir . -style goZero
	# goctl api go -api api/${SERVICE_FILE_NAME}.api -dir api/. -style goZero

goctl-rpc:
	goctl rpc proto -src ${SERVICE_FILE_NAME}.proto -dir . -style goZero

test-rpc:
	go test ./rpc/test -v

test:
	make test-rpc


swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename docs/${SERVICE_PREFIX}.json" -api api/${SERVICE_FILE_NAME}.api -dir .

swagger-file:
	@echo docs/${SERVICE_PREFIX}.json


version:
	@echo ${SERVICE_VERSION}


image-api-prefix:
	@echo ${SERVICE_NAME}-api:${SERVICE_VERSION}

image-rpc-prefix:
	@echo ${SERVICE_NAME}-rpc:${SERVICE_VERSION}

