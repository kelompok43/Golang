package membership

import (
	handlerAPI "github.com/kelompok43/Golang/membership/handler/api"
	repoMySQL "github.com/kelompok43/Golang/membership/repository/mysql"
	service "github.com/kelompok43/Golang/membership/service"
	"gorm.io/gorm"
)

func NewMembershipFactory(db *gorm.DB) (membershipHandler handlerAPI.MembershipHandler) {
	membershipRepo := repoMySQL.NewMembershipRepository(db)
	membershipService := service.NewMembershipService(membershipRepo)
	membershipHandler = handlerAPI.NewMembershipHandler(membershipService)
	return
}
