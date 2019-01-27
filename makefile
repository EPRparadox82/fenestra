version=$(shell git tag --points-at HEAD)
date=$(shell date "+%b %Y")
name=fenestra
.PHONY: all

all:
	@echo " make <cmd>"
	@echo ""
	@echo "Commands:"
	@echo "  build             - runs go build with ldflags version=${version} & date=${date}"
	@echo "  fetch_stuff       - installs all needed libaries"
	@echo ""


build: clean
	@go build -v -ldflags '-X "main.version=${version}" -X "main.compdate=${date}"' -o ${name}

clean:
	@rm -f ${name}

fetch_stuff:
	@go get -u github.com/aarzilli/nucular
	@go get -u github.com/pborman/getopt/v2
	@echo "All nessecary libaries installed"
