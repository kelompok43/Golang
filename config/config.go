package config

import (
	"fmt"
	"os"

	repoAdmin "github.com/kelompok43/Golang/admin/repository/mysql"
	repoPM "github.com/kelompok43/Golang/payment_method/repository/mysql"
	repoTrainer "github.com/kelompok43/Golang/trainer/repository/mysql"
	repoTransaction "github.com/kelompok43/Golang/transaction/repository/mysql"
	repoUser "github.com/kelompok43/Golang/user/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
}

var Conf Config

func Init() {
	Conf = Config{
		DBNAME: os.Getenv("DB_NAME"),
		DBUSER: os.Getenv("DB_USER"),
		DBPASS: os.Getenv("DB_PASS"),
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
	}
}

func DBInit() (DB *gorm.DB) {
	DB, _ = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				Conf.DBUSER,
				Conf.DBPASS,
				Conf.DBHOST,
				Conf.DBPORT,
				Conf.DBNAME,
			),
		),
	)
	return
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&repoUser.User{},
		&repoUser.UserDetail{},
		&repoAdmin.Admin{},
		&repoTrainer.Trainer{},
		&repoPM.PaymentMethod{},
		&repoTransaction.Transaction{},
	)
}
