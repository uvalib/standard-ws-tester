# project specific definitions
SRCDIR = cmd

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

#build: darwin

#all: darwin linux

test:
	@ \
	(cd "$(SRCDIR)/tests" && $(GOTEST) -v) ;

#clean:
#	$(GOCLEAN)
#	rm -rf bin
