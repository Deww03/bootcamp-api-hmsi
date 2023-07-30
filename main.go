package main

import (
	"fmt"
	"os"

	"github.com/Deww03/bootcamp-api-hmsi/connectDB"
	"github.com/Deww03/bootcamp-api-hmsi/modules/customers/customerHandler"
	"github.com/Deww03/bootcamp-api-hmsi/modules/customers/customerRepository"
	"github.com/Deww03/bootcamp-api-hmsi/modules/customers/customerUsecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	log.Info().Str("version", os.Getenv("APP_VERSION")).Msg("Application started")

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("PORT")
	DB_DRIVER := os.Getenv("DB_DRIVER")

	log.Info().Msg("DB_HOST: " + DB_HOST)
	log.Info().Msg("DB_PORT: " + DB_PORT)
	log.Info().Msg("DB_USER: " + DB_USER)
	log.Info().Msg("DB_PASSWORD: " + DB_PASSWORD)
	log.Info().Msg("DB_NAME: " + DB_NAME)
	log.Info().Msg("PORT: " + PORT)
	log.Info().Msg("DB_DRIVER: " + DB_DRIVER)

	db, errConn := connectDB.GetConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)
	if errConn != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println("Successfully connected")

	//---Inisialisasi router---
	var router = gin.Default()

	//---Inisialisasi modules---
	customerRepo := customerRepository.NewCustomerRepository(db)
	customerUC := customerUsecase.NewCustomerUsecase(customerRepo)
	customerHandler.NewCustomerHandler(router, customerUC)

	log.Info().Msg("Server running on port" + PORT)
	router.Run(":" + PORT)

	// DB struct initialize
	//DB := query.DB{Conn: db}

	// --create---
	// err = DB.Create(&query.Customers{
	// 	Name:  "Adam",
	// 	Phone: "085711228241",
	// 	Email: "adam@gmail.com",
	// 	Age:   19,
	// })

	// fmt.Println(err)

	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }

	// fmt.Println("Insert Data Berhasil")

	// ---read---
	// result, err := DB.Read()

	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println(result)

	// ---update---
	// err = DB.Update(&query.Customers{
	// 	Id:    1,
	// 	Name:  "Adew",
	// 	Phone: "085711228241",
	// 	Email: "adew@gmail.com",
	// 	Age:   19,
	// })

	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }

	// fmt.Println("Update Data Berhasil")

	// ---delete---
	// err = DB.Delete(2)
	// if err != nil {
	// 	log.Error().Msg(errConn.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("Delete Data Berhasil")
}
