# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.15.4-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/berty-library-test
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.12
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="berty-library-test" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/berty-library-test/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/berty-library-test" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/berty-library-test" \
                org.label-schema.help="docker exec -it $CONTAINER berty-library-test --help"
COPY            --from=builder /go/bin/berty-library-test /bin/
ENTRYPOINT      ["/bin/berty-library-test"]
#CMD             []
