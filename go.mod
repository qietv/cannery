module cannery

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/boltdb/bolt v1.3.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/hanskorg/logkit v1.0.0
	github.com/hashicorp/raft v1.2.0
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea
	github.com/hpcloud/tail v1.0.0
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/qietv/qgrpc v0.0.2
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	google.golang.org/grpc v1.35.0
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

go 1.15

//replace (
//     github.com/qietv/qgrpc => ../qgrpc
// )
