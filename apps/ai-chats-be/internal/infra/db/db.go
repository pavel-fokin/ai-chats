package db

type Config struct {
	DATABASE_TYPE string `env:"AICHATS_DATABASE_TYPE" envDefault:"sqlite"`
	DATABASE_URL  string `env:"AICHATS_DATABASE_URL" envDefault:"file:memdb1?mode=memory&cache=shared"`
}
