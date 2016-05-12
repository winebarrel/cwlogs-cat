PREFIX=/usr/local
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

ifdef GOPATH
  RUNTIME_GOPATH=$(GOPATH):`pwd`
else
  RUNTIME_GOPATH=`pwd`
endif

all: cwlogs-cat

go-get:
	go get github.com/aws/aws-sdk-go

cwlogs-cat: go-get main.go src/cwlogs_cat/optparse.go src/cwlogs_cat/cwlogs_cat.go
	GOPATH=$(RUNTIME_GOPATH) go build -o cwlogs-cat main.go

install: cwlogs-cat
	install -m 755 cwlogs-cat $(DESTDIR)$(PREFIX)/bin/

clean:
	rm -f cwlogs-cat *.gz

package: clean cwlogs-cat
	gzip -c cwlogs-cat > cwlogs-cat-$(VERSION)-$(GOOS)-$(GOARCH).gz

deb:
	dpkg-buildpackage -us -uc
