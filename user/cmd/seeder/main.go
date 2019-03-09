package main

import (
	"log"

	"github.com/emurmotol/project/user/pkg/service"
	"github.com/emurmotol/project/user/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://root:postgres@localhost:5433/project?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	username := "testuser"
	role := "user"
	user := service.User{
		Username: username,
		Email:    "testuser@example.com",
		Password: "secret",
		Role:     role,
	}
	if err := db.FirstOrCreate(&user, service.User{Username: username}).Error; err != nil {
		panic(err)
	}
	log.Println(user)

	casbinRules := []utils.CasbinRule{
		utils.CasbinRule{
			PType: "p",
			V0:    role,
			V1:    "/pb.User/CreateUser",
			V2:    "read",
		},
		utils.CasbinRule{
			PType: "p",
			V0:    role,
			V1:    "/pb.User/GetByUsername",
			V2:    "read",
		},
	}
	for _, cr := range casbinRules {
		if err := db.FirstOrCreate(&cr, utils.CasbinRule{V0: cr.V0, V1: cr.V1, V2: cr.V2}).Error; err != nil {
			panic(err)
		}
		log.Println(cr)
	}
}
