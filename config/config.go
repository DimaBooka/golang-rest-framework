package config

import (
	"flag"
)

var (
	IsSeed   = flag.Bool("make.seed", false, "Clear database")
	DB_ADDR  = flag.String("db.address", "localhost", "Database address (localhost in original)")
	DB_NAME  = flag.String("db.name", "test", "Database name(test in original)")
	HttpAddr = flag.String("http.addr", ":8001", "Address for HTTP (JSON) server")
)
