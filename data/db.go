package data

import (
	"fmt"
	"os"
	"time"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stewie1520/pm/data/postgres"
)

type Context struct {
	 activities *postgres.ActivityModel
}

func getDb() (*sql.DB, error) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getContext() (*Context, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	return &Context{ activities: &postgres.ActivityModel{DB: db} }, nil
}


// GetLastLogin return the most recent time user login
func GetLastActivityTime() *time.Time {
	// default value of now is 20 minutes ago
	now := time.Now().Add(-20*time.Minute)

	ctx, err := getContext()
	if err != nil {
		return &now
	}

	actv := ctx.activities.Latest()

	if actv == nil {
		return &now
	}

	return &actv.Time
}

// CheckMasterKey check if password is correct
func CheckMasterKey(password string) bool {
	return false
}
