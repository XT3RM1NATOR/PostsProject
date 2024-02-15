package main

import (
	"github.com/XT4RM1NATOR/PostsProject/initializers"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func init() {
	initializers.LoadEnvVariables()
	DB = initializers.ConnectToDb()
}
