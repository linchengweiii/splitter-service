module github.com/linchengweiii/splitter

go 1.23.1

require (
	github.com/gorilla/mux v1.8.1
	github.com/linchengweiii/splitter/pkg/expense v0.0.0
	github.com/linchengweiii/splitter/pkg/group v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

replace github.com/linchengweiii/splitter/pkg/expense => ./pkg/expense

replace github.com/linchengweiii/splitter/pkg/group => ./pkg/group
