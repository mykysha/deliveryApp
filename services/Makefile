update-gomod-all:
	make download-gomod-all
	make gomod-tidy-all

download-gomod-all:
	make -C ./accounting/ download-gomod
	make -C ./consumer/ download-gomod
	make -C ./courier/ download-gomod
	make -C ./delivery/ download-gomod
	make -C ./kitchen/ download-gomod
	make -C ./order/ download-gomod
	make -C ./restaurant/ download-gomod

gomod-tidy-all:
	make -C ./accounting/ gomod-tidy
	make -C ./consumer/ gomod-tidy
	make -C ./courier/ gomod-tidy
	make -C ./delivery/ gomod-tidy
	make -C ./kitchen/ gomod-tidy
	make -C ./order/ gomod-tidy
	make -C ./restaurant/ gomod-tidy
