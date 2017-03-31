package g

import (
	"log"

	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

func InitEnv() {
	globalSessions, err := session.NewManager("redis", &session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600, ProviderConfig: "127.0.0.1:6379"})
	if err != nil {
		log.Fatalln(err)
	}
	go globalSessions.GC()
}
