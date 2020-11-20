package data

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/stewie1520/pm/constants"
	"github.com/stewie1520/pm/data/model"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// session, err = mgo.Dial(os.Getenv("DB_CONNECTION"))
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{os.Getenv("DB_ADDRS")},
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PW"),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
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
