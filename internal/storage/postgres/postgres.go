package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type databaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadDBConfig() *databaseConfig {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return &databaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

func ConnectDatabase() *gorm.DB {
	config := loadDBConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar com PostgreSQL: %v", err)
	}

	if err = runMigrations(db); err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Erro ao obter database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Conectado ao PostgreSQL com sucesso!")
	return db
}

func TestConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func runMigrations(db *gorm.DB) error {
	log.Println("Executando migrações...")

	// Drop existing tables in reverse order to avoid foreign key constraints
	if err := db.Migrator().DropTable(); err != nil {
		return fmt.Errorf("erro ao dropar tabelas: %v", err)
	}

	// Recreate tables with new schema
	if err := db.AutoMigrate(); err != nil {
		return fmt.Errorf("erro ao executar migrações: %v", err)
	}
	return nil
}
