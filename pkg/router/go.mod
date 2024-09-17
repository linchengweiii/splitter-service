module github.com/linchengweiii/splitter/pkg/router

require (
	github.com/linchengweiii/splitter/pkg/expense v0.0.0
	github.com/linchengweiii/splitter/pkg/group v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

go 1.23.1

replace github.com/linchengweiii/splitter/pkg/expense => ../expense

replace github.com/linchengweiii/splitter/pkg/group => ../group
