swagger-generate:
	swagger generate spec --scan-models --output=./api/v1/swagger.yaml

golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

status-check:
	-curl http://localhost:8085/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm order_app

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start order-service

docker-stop:
	docker container stop order-service

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf

	go get -u github.com/adrianbrad/psqldocker
	go get -u github.com/adrianbrad/psqltest
	go get -u github.com/friendsofgo/errors
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq
	go get -u github.com/stretchr/testify
	go get -u github.com/volatiletech/sqlboiler/v4
	go get -u github.com/volatiletech/strmangle

	go get -u github.com/nndergunov/deliveryApp/pkg/api
	go get -u github.com/nndergunov/deliveryApp/pkg/configreader
	go get -u github.com/nndergunov/deliveryApp/pkg/grpcserver
	go get -u github.com/nndergunov/deliveryApp/pkg/logger
	go get -u github.com/nndergunov/deliveryApp/pkg/messagebroker
	go get -u github.com/nndergunov/deliveryApp/pkg/server
	go get -u github.com/nndergunov/deliveryApp/services/accounting/api/v1/rest/accountingapi
	go get -u github.com/nndergunov/deliveryApp/services/restaurant/api/v1/restaurantapi

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18

proto:
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative order.proto
