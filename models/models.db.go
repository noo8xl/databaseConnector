package models

type MongoOpts struct {
	User string
	Pwd  string
	Name string
}

type MySqlOpts struct {
	User string
	Pwd  string
	Name string
	Host string
	Port string
}
