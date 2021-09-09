package cmd

import (
	"naerarauth/internals/implementation"
	"naerarauth/internals/persistence"
	"naerarshared/models"

	"naerarauth/memstore"
	"naerarauth/pkg"
	"os"
)

var port = os.Getenv("AuthPort")

// RunServer -
func RunServer() {
	var eliest = pkg.NaerarAuthHandler{}

	//The Db
	AuthDbConfig := persistence.InitAuthDbConfig()
	psql := implementation.NewPsqlLayer(persistence.Config(&AuthDbConfig))
	psql.NaerarUserDb.AutoMigrate(models.UserAccount{})
	eliest.InitDb(psql)

	//The Cache
	redis := memstore.NewRedisClient(os.Getenv("RedisHost"), os.Getenv("RedisPass"))
	eliest.InitMemStore(redis)

	//The Router
	eliest.InitRouter()
	eliest.StartHttp(port)
}
