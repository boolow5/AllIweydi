package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/boolow5/BolWeydiWararka/g"
	/* _ "github.com/go-sql-driver/mysql"
	   _ "github.com/lib/pq" */
	_ "github.com/mattn/go-sqlite3"
)

var (
	o     orm.Ormer
	DEBUG bool
)

type MyModel interface {
	Valid() bool
	SetId(id int)
	String() string
}

func init() {
	fmt.Println("Initialing model's defaults")
	dbHome := os.Getenv("DB_HOME")
	if dbHome == "" {
		dbHome = os.Getenv("HOME")
	}

	fmt.Printf("Connecting to \"%s\"\n", dbHome+"/news.db")

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", dbHome+"/news.db")

	orm.RegisterModel(new(User), new(Profile), new(Website))

	if g.AUTO_MIGRATE {
		name := "default"                          // Database alias.
		force := true                              // Drop table and re-create.
		verbose := true                            // Print log
		err := orm.RunSyncdb(name, force, verbose) // Sync with database
		if err != nil {
			Verbose(err.Error())
		}
	}
	o = orm.NewOrm()
	o.Raw("PRAGMA foreign_keys = ON")
}

func Verbose(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format+"\n", args...)
	}
}
