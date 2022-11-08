update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/spf13/viper

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
