package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

// service the interacts with the db
type Service interface {
	Health() map[string]string

	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dbName     = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	dbInstance *service
)

func New() Service {

	//reuse active conection
	if dbInstance != nil {
		return dbInstance
	}

	// opening a driver
	// will not connect to the db
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName))
	if err != nil {
		// dns parse error or initialization error
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// pinging the db
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["errpr"] = fmt.Sprintf("database down: %v", err)
		log.Fatalf("database down: %v", err)
		return stats
	}

	// get db stats like open connections, idle, etc.
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// evaluate stats and provide health message
	if dbStats.OpenConnections >= 40 {
		stats["message"] = "The database is experiencing heavy load"
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events"
	}
	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the conection pool settings"
	}
	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// this func closes the database connection
// logs a message indicatinh wich db was disconnected
// if its successfully closed, return nil
// if not, returns the error
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbName)
	return s.db.Close()
}
