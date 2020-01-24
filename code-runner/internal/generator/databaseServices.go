package generator

func (app *GenApp) generateMysqlService() {
	app.Data = make(map[string][]byte)
	app.Data["config/mysql.cnf"] = app.genMysqlConf()
}

func (app *GenApp) genMysqlConf() []byte {
	index := `
[client]
port=`+app.App.Spec["port"]+`
socket=/tmp/mysql.sock

[mysqld]
port=`+app.App.Spec["port"]+`
socket=/tmp/mysql.sock
key_buffer_size=16M
max_allowed_packet=128M

[mysqldump]
quick `
	return []byte(index)
}

func (app *GenApp) generateMongoService() {
	app.Data = make(map[string][]byte)
	app.Data["config/mongod.conf"] = app.genMongoConf()
}

func (app *GenApp) genMongoConf() []byte {
	index :=
		` 
# mongod.conf

# for documentation of all options, see:
#   http://docs.mongodb.org/manual/reference/configuration-options/
# Where and how to store data.
#storage:
#  dbPath: /var/lib/mongodb
#  journal:
#    enabled: true
#  engine:
#  mmapv1:
#  wiredTiger:

# where to write logging data.
#systemLog:
#  destination: file
#  logAppend: true
#  path: /var/log/mongodb/mongod.log

# network interfaces
net:
  port: ` + app.App.Spec["port"] + `
  bindIp: 0.0.0.0

# how the process runs
#processManagement:
#  timeZoneInfo: /usr/share/zoneinfo

#security:

#operationProfiling:

#replication:

#sharding:

## Enterprise-Only Options:

#auditLog:

#snmp:
`
	return []byte(index)
}
