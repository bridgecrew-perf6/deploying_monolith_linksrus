module deploying_monolith_linksrus

go 1.17

replace work_scheduling => ../work_scheduling

require (
	github.com/blevesearch/bleve v1.0.14
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/juju/clock v0.0.0-20190205081909-9c5c9712527c
	github.com/lib/pq v1.10.4
	github.com/microcosm-cc/bluemonday v1.0.17
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.43.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	work_scheduling v0.0.0-00010101000000-000000000000
)

require (
	github.com/RoaringBitmap/roaring v0.4.23 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/blevesearch/go-porterstemmer v1.0.3 // indirect
	github.com/blevesearch/mmap-go v1.0.2 // indirect
	github.com/blevesearch/segment v0.9.0 // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/blevesearch/zap/v11 v11.0.14 // indirect
	github.com/blevesearch/zap/v12 v12.0.14 // indirect
	github.com/blevesearch/zap/v13 v13.0.6 // indirect
	github.com/blevesearch/zap/v14 v14.0.5 // indirect
	github.com/blevesearch/zap/v15 v15.0.3 // indirect
	github.com/couchbase/vellum v1.0.2 // indirect
	github.com/glycerine/go-unsnap-stream v0.0.0-20190901134440-81cf024a9e0a // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/juju/errors v0.0.0-20210818161939-5560c4c073ff // indirect
	github.com/juju/loggo v0.0.0-20210728185423-eebad3a902c4 // indirect
	github.com/juju/testing v0.0.0-20211215003918-77eb13d6cad2 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/steveyen/gtreap v0.1.0 // indirect
	github.com/tinylib/msgp v1.1.1 // indirect
	github.com/willf/bitset v1.1.10 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
