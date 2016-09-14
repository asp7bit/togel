package conf

import (
	"log"
	"database/sql"
	"github.com/ruizu/gcfg"
	_ "github.com/lib/pq" // Dependency to connect postgres
)

var DB DBBundle
// ReadConfig : read file .ini
func ReadConfig(c *Config, filePath string) bool {
	if err := gcfg.ReadFileInto(c, filePath); err != nil {
		log.Printf("%s\n", err)
		return false
	}
	return true
}

func InitDB(dbCore *string) {

	core, err := sql.Open("postgres", *dbCore)
	if err != nil {
		log.Fatal(err, "db.core not available with config ", *dbCore)
	}

	DB = DBBundle {
		core,
	}

	_, err = DB.Core.Query("SELECT 1")
	if err != nil {
		log.Fatal("TKP Core data not accessible, please check config")
	}
	
}