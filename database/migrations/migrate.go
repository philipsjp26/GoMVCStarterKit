package migrations

import (
	"GoMVCStarterKit/config"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"time"
)

var (
	flags = flag.NewFlagSet("db:migrate", flag.ExitOnError)
)

func Migrate() {
	flags.Parse(os.Args[2:])
	args := flags.Args()
	cfg := config.NewEnv()

	// fmt.Println(args[0]) // create ${filename}
	switch args[0] {
	case "create":
		migrationName := args[1]
		timestamp := time.Now().Format("20060102150405")
		upFileName := fmt.Sprintf("database/migrations/%s_%s.up.sql", timestamp, migrationName)
		downFileName := fmt.Sprintf("database/migrations/%s_%s.down.sql", timestamp, migrationName)
		createFile(upFileName)
		createFile(downFileName)
		fmt.Println("Success created migrations")
	case "up":
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("failed to get current working directory: %v", err)
		}
		migrationsDir := filepath.Join(wd, "database/migrations")

		m, err := migrate.New("file://"+migrationsDir, fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
			cfg.Database.Username, cfg.Database.Password, cfg.Database.Hostname, cfg.Database.Port, cfg.Database.Name))
		if err != nil {
			log.Fatal("error : ", err)
		}
		if err = m.Up(); err != nil {
			log.Fatal("apa error : ", err)
		}
		fmt.Println("success migrate")
	case "down":
		// TODO: Migrate down
	}
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file :%s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println("Created file")
}
