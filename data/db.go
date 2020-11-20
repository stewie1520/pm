package data

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"crypto/tls"

	"github.com/joho/godotenv"
	"github.com/stewie1520/pm/constants"
	"github.com/stewie1520/pm/data/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client mongo.

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	mongoUri := os.Getenv("DB_CONNECTION")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
}

// GetLastLogin return the most recent time user login
func GetLastLogin() (*time.Time, error) {
	var lastLoginAtv *model.Activity

	Activities := session.DB(os.Getenv("DB_NAME")).C("activities")

	err := Activities.Find(struct {
		Action string
	}{
		Action: constants.ActionLogin,
	}).One(lastLoginAtv)

	if err != nil {
		return nil, err
	}

	return &lastLoginAtv.Time, nil
}

// CheckMasterKey check if password is correct
func CheckMasterKey(password string) bool {
	return false
}
