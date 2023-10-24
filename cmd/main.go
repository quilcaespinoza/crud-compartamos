package main

import (
	"crud-compartamos/config"
	"crud-compartamos/repository"
	"crud-compartamos/routes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func main() {
	dbConfig := config.LoadDatabaseConfig()
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DatabaseName)
	driver := "mysql"

	db, err := sql.Open(driver, databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Executing migrations")
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	fmt.Println("Migration successful")

	n, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Error aplicando migraciones: %v", err)
	}

	fmt.Printf("Se aplicaron %d migraciones\n", n)

	userService := repository.NewUserService(db)

	r := gin.Default()
	routes.InitializeRoutes(r, userService)

	err = r.Run(":8080")
	if err != nil {
		return
	}

}
