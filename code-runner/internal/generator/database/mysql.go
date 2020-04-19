package database

func GenMysqlConf(port string) []byte {
	data := `
[client]
port=`+port+`
socket=/tmp/mysql.sock

[mysqld]
port=`+port+`
socket=/tmp/mysql.sock
key_buffer_size=16M
max_allowed_packet=128M

[mysqldump]
quick `
	return []byte(data)
}
