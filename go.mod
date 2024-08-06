module moul.io/berty-library-test

go 1.15

require (
	berty.tech/berty/v2 v2.148.0
	github.com/gogo/protobuf v1.3.1
	github.com/mdp/qrterminal/v3 v3.0.0
	github.com/tailscale/depaware v0.0.0-20200914232109-e09ee10c1824
	go.uber.org/goleak v1.1.10
	google.golang.org/grpc v1.56.3
	moul.io/u v1.10.0
)

replace (
	bazil.org/fuse => bazil.org/fuse v0.0.0-20200117225306-7b5117fecadc // specific version for iOS building
	github.com/agl/ed25519 => github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // latest commit before the author shutdown the repo; see https://github.com/golang/go/issues/20504
	github.com/ipld/go-ipld-prime => github.com/ipld/go-ipld-prime v0.0.2-0.20191108012745-28a82f04c785 // specific version needed indirectly
	github.com/ipld/go-ipld-prime-proto => github.com/ipld/go-ipld-prime-proto v0.0.0-20191113031812-e32bd156a1e5 // specific version needed indirectly
	github.com/libp2p/go-libp2p-kbucket => github.com/libp2p/go-libp2p-kbucket v0.4.2 // specific version needed indirectly
	github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.18.0 // required by go1.15
)
