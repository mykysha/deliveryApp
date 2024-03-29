golangci-lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

genservice:
	mockgen -source=pkg/service/courierservice/courierservice.go \
    -destination=pkg/mocks/mock_courierservice.go

genstorage:
	mockgen -source=pkg/service/courierservice/storage.go \
	-destination=pkg/mocks/mock_courierstorage.go

go-test-unit:
	go test api/v1/handler/courierhandler/courierhandler_test.go -v
	go test pkg/service/courierservice/courierservice_test.go -v

go-test-integrational:
	go test pkg/storage/courierstorage/courierstorage_integrational_test.go -v

go-test-all:
	make go-test-unit
	make go-test-integrational

status-check:
	-curl http://localhost:8082/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm courier_app

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start courier-service

docker-stop:
	docker container stop courier-service

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/nndergunov/deliveryApp/pkg/api
	go get -u github.com/nndergunov/deliveryApp/pkg/configreader
	go get -u github.com/nndergunov/deliveryApp/pkg/logger
	go get -u github.com/nndergunov/deliveryApp/pkg/server
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18

swagger:
	swagger generate spec -o ./api/v1/swagger.yaml --scan-models
