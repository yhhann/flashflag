export GOPATH := ${GOPATH}:$(shell pwd)
export LDFLAG=" -s -X main.buildTime=`date -u '+%Y%m%d-%I%M%S%Z'`"

all: conf
	@echo "make conf         : build conf"
	@echo "make clean        : clean conf"

.PHONY: conf
conf:
	go install -ldflags ${LDFLAG} flashflag/cmd

clean:
	rm -fr bin
	rm -fr pkg
