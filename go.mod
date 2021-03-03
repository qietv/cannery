module cannery

require (
	github.com/boltdb/bolt v1.3.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/qietv/qgrpc v0.0.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	google.golang.org/grpc v1.35.0
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

go 1.15

replace (
    github.com/qietv/qgrpc => ../qgrpc
)