package main

import (
	"context"
	"log"
	"os"

	"codeid.northwind/config"
	repositories "codeid.northwind/repositories/dbContext"
	"codeid.northwind/server"

	_ "github.com/lib/pq"
)

func main() {
	// ketika menjalankan server, jangan pakai fmt.print, tetapi pakai log.println
	log.Println("starting northwind")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database...")
	dbHandler := server.InitDatabase(config)
	// log.Println(dbHandler)

	log.Println("Initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()

	// test insert to category
	ctx := context.Background()
	queries := repositories.New(dbHandler)

	newCategory, err := queries.CreateCategory(ctx,
		repositories.CreateCategoryParams{
			CategoryID:   101,
			CategoryName: "Mainan",
			Description:  "Mainan Anak",
			Picture:      nil,
		},
	)

	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Println(newCategory)
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "nortwind" + env
	}

	return "northwind"
}
