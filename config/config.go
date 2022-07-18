package config

import (
	"fmt"
	"os"

	repoAdmin "github.com/kelompok43/Golang/admin/repository/mysql"
	repoBook "github.com/kelompok43/Golang/book/repository/mysql"
	repoClass "github.com/kelompok43/Golang/class/repository/mysql"
	repoMembership "github.com/kelompok43/Golang/membership/repository/mysql"
	repoNews "github.com/kelompok43/Golang/news/repository/mysql"
	repoPM "github.com/kelompok43/Golang/payment_method/repository/mysql"
	repoTrainer "github.com/kelompok43/Golang/trainer/repository/mysql"
	repoTransaction "github.com/kelompok43/Golang/transaction/repository/mysql"
	repoUser "github.com/kelompok43/Golang/user/repository/mysql"
	"gorm.io/driver/postgres"
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

// func DBInit() (DB *gorm.DB) {
// 	DB, _ = gorm.Open(
// 		mysql.Open(
// 			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 				Conf.DBUSER,
// 				Conf.DBPASS,
// 				Conf.DBHOST,
// 				Conf.DBPORT,
// 				Conf.DBNAME,
// 			),
// 		),
// 	)
// 	return
// }

func DBInit() (DB *gorm.DB) {
	dbURL := fmt.Sprintf(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			Conf.DBHOST,
			Conf.DBUSER,
			Conf.DBPASS,
			Conf.DBNAME,
			Conf.DBPORT,
		),
	)

	DB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURL, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,  // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	return
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&repoUser.User{},
		&repoUser.UserDetail{},
		&repoAdmin.Admin{},
		&repoTrainer.Trainer{},
		&repoPM.PaymentMethod{},
		&repoMembership.MembershipCategory{},
		&repoMembership.Membership{},
		&repoTransaction.Transaction{},
		&repoTransaction.TransactionDetail{},
		&repoClass.ClassCategory{},
		&repoClass.OnlineClass{},
		&repoClass.OfflineClass{},
		&repoBook.BookOnlineClass{},
		&repoBook.BookOfflineClass{},
		&repoNews.NewsCategory{},
		&repoNews.News{},
	)
}
