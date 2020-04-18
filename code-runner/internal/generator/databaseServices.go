package generator

import "code-runner/internal/generator/database"

func (app *GenApp) generateMysqlService() {
	app.Data = make(map[string][]byte)
	app.Data["config/mysql.cnf"] = database.GenMysqlConf(app.App.Spec["port"])
}

func (app *GenApp) generateMongoService() {
	app.Data = make(map[string][]byte)
	app.Data["config/mongod.conf"] = database.GenMongoConf(app.App.Spec["port"])
}

func (app *GenApp) generateRedisService() {
	app.Data = make(map[string][]byte)
	app.Data["config/redis.conf"] = database.GenRedisConf(app.App.Spec["port"])
}
