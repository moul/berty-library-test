GOPKG ?=	moul.io/berty-library-test
DOCKER_IMAGE ?=	moul/berty-library-test
GOBINS ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ berty-library-test -h' > .tmp/usage.txt
	berty-library-test -h 2>> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
