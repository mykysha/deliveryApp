docker-build-all:
	docker network inspect delivery_app_net >/dev/null 2>&1 || \
        docker network create --driver bridge delivery_app_net
	make -C ./pkg/messagebroker docker-build
	make -C ./services/accounting docker-build
	make -C ./services/consumer docker-build
	make -C ./services/courier docker-build
	make -C ./services/delivery docker-build
	make -C ./services/kitchen docker-build
	make -C ./services/order docker-build
	make -C ./services/restaurant docker-build

cl-start-all:
	make -C ./services/accounting cl-start
	make -C ./services/consumer cl-start
	make -C ./services/courier cl-start
	make -C ./services/delivery cl-start
	make -C ./services/kitchen cl-start
	make -C ./services/order cl-start
	make -C ./services/restaurant cl-start

update-gomod-all:
	make -C ./pkg update-gomod-all
	make -C ./services update-gomod-all

download-gomod-all:
	make -C ./pkg download-gomod-all
	make -C ./services download-gomod-all

gomod-tidy-all:
	make -C ./pkg gomod-tidy-all
	make -C ./services gomod-tidy-all

check-all:
	make -C ./services/accounting status-check
	make -C ./services/consumer status-check
	make -C ./services/courier status-check
	make -C ./services/delivery status-check
	make -C ./services/kitchen status-check
	make -C ./services/order status-check
	make -C ./services/restaurant status-check
