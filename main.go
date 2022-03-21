package main

import (
	"fmt"
	"log"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.Dbname)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@examle.com"
	// u.PassWord = "testtest"

	// fmt.Println(u)

	// u.CreateUser()

	// u, err := models.GetUser(1)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(u)

	// u.Name = "Test3sss3"
	// u.Email = "test2@test.com"

	// u.UpdateUser()

	// u, _ = models.GetUser(1)

	// u.DeleteUser()

	// u, _ = models.GetUser(1)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(u)

	u, err := models.GetUser(1)

	if err != nil {
		log.Fatalln(err)
	}

	u.DeleteUser()

	u, _ = models.GetUser(1)

	fmt.Println(u)

}
