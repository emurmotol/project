package main

import (
	"log"

	"github.com/emurmotol/project/user/pkg/service"
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
	user := service.User{
		Username: username,
		Email:    "testuser@example.com",
		Password: "secret",
		Role:     "user",
	}
	if err := db.FirstOrCreate(&user, service.User{Username: username}).Error; err != nil {
		panic(err)
	}
	log.Println(user)

	casbinRules := []service.CasbinRule{
		service.CasbinRule{
			PType: "p",
			V0:    "/pb.User/CreateUser",
			V1:    "read",
		},
		service.CasbinRule{
			PType: "p",
			V0:    "/pb.User/GetByUsername",
			V1:    "read",
		},
	}
	for _, cr := range casbinRules {
		if err := db.FirstOrCreate(&cr, service.CasbinRule{V0: cr.V0, V1: cr.V1}).Error; err != nil {
			panic(err)
		}
		log.Println(cr)
	}
}
