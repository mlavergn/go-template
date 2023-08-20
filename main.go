package main

import (
	"demo"
	"log"
)

func main() {
	log.Println(demo.Message())

	log.Println("DB client demo")
	db := NewDBClient("demo.db")
	db.exec("create table dual (name varchar(64) null)")
	db.exec("insert into dual (name) values ('X')")
	result := db.query("select * from dual")
	log.Println(result)

	log.Println("HTTP client demo")
	client := NewHTTPClient()
	client.get("http://example.com")

	log.Println("HTTP server demo")
	server := NewHTTPServer()
	server.start()
}
