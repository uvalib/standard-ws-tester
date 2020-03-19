# project specific definitions
SRCDIR = cmd

GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

test:
	@ \
	(cd "$(SRCDIR)/tests" && $(GOTEST) -v) ;

#deps:
#	cd $(SRCDIR); $(GOGET) -u
#	$(GOMOD) tidy
