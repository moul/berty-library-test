GOPKG ?=	moul.io/berty-library-test
DOCKER_IMAGE ?=	moul/berty-library-test
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ berty-library-test' > .tmp/usage.txt
	berty-library-test 2>&1 >> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
