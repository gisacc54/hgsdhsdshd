package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   uint64 `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}

func main() {
	app := fiber.New()

	db := DB()
	if db == nil {
		println("Fail")
		return
	}

	//Command
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		println("Command mode")
		if args[0] == "migrate:seed" {
			println("Migrate Data Base and seed!")
			db.AutoMigrate(&User{})
		}
		return
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Docker! <3")
	})

	app.Post("/users", func(c *fiber.Ctx) error {

		var users User
		err := c.BodyParser(&users)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  false,
				"message": "Fail To create User",
			})
		}

		db.Create(&users)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  true,
			"message": "User created successful!",
		})
	})
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []User
		db.Find(&users)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": true,
			"users":  users,
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(struct{ Status string }{Status: "OK"})
	})

	app.Listen(":8080")
}

func DB() *gorm.DB {
	dsn := "root:root_password@tcp(sis-mysql-container:3306)/sis_system_setting_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("err")
		return nil
	}

	return db
}
