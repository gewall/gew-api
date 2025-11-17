package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	sqlDb, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Fail to connect database: %s", err)
		return nil
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fail to connect database: %s", err)
		return nil
	}

	if err := sqlDb.Ping(); err != nil {
		log.Fatalf("Fail to connect database: %s", err)
		return nil
	}

	fmt.Print("Database Connected!")

	return gormDb
}
