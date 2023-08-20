module app

go 1.21.0

require (
	demo v0.0.0
	// modernc.org/sqlite v1.25.0
	github.com/mattn/go-sqlite3 v1.14.17
)

replace demo => ./demo
