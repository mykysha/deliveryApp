update-gomod-all:
	make download-gomod-all
	make gomod-tidy-all

download-gomod-all:
	make -C ./api download-gomod
	make -C ./configreader download-gomod
	make -C ./grpcserver download-gomod
	make -C ./messagebroker download-gomod
	make -C ./server download-gomod

gomod-tidy-all:
	make -C ./api gomod-tidy
	make -C ./configreader gomod-tidy
	make -C ./grpcserver gomod-tidy
	make -C ./messagebroker gomod-tidy
	make -C ./server gomod-tidy
