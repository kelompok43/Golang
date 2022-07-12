package news

import (
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	repoHourmailer "github.com/kelompok43/Golang/hourmailer/repository/hourmailer"
	repoMySQLM "github.com/kelompok43/Golang/membership/repository/mysql"
	serviceM "github.com/kelompok43/Golang/membership/service"
	handlerAPI "github.com/kelompok43/Golang/news/handler/api"
	repoMySQL "github.com/kelompok43/Golang/news/repository/mysql"
	service "github.com/kelompok43/Golang/news/service"
	repoMySQLU "github.com/kelompok43/Golang/user/repository/mysql"
	serviceU "github.com/kelompok43/Golang/user/service"
	"gorm.io/gorm"
)

func NewNewsFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (newsHandler handlerAPI.NewsHandler) {
	hourmailerRepo := repoHourmailer.NewHourmailerRepository()
	membershipRepo := repoMySQLM.NewMembershipRepository(db)
	membershipService := serviceM.NewMembershipService(membershipRepo)
	userRepo := repoMySQLU.NewUserRepository(db)
	userService := serviceU.NewUserService(userRepo, configJWT, membershipService)
	newsRepo := repoMySQL.NewNewsRepository(db)
	newsService := service.NewNewsService(newsRepo, hourmailerRepo, userService)
	newsHandler = handlerAPI.NewNewsHandler(newsService)
	return
}
