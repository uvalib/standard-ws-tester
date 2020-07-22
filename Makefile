# project specific definitions
SRCDIR = cmd

GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod
GOFMT = $(GOCMD) fmt
GOVET = $(GOCMD) vet

test:
	@ \
	(cd "$(SRCDIR)/tests" && $(GOTEST) -v) ;

dep:
	cd $(SRCDIR)/tests; $(GOGET) -u
	$(GOMOD) tidy
	$(GOMOD) verify

fmt:
	cd $(SRCDIR)/tests; $(GOFMT)

vet:
	cd $(SRCDIR)/tests; $(GOVET)
