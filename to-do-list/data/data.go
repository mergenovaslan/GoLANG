package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"to-do-list/config"

	_ "github.com/lib/pq"
)

const (
	port = 5432
	host = "127.0.0.1"
)

var DB *sql.DB
var Config config.Config

func StorageConfig(conf config.Config) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, conf.DBusername, conf.DBpassword, conf.DBname,
	)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("We have problems with connection database", err)
		os.Exit(1)
	}
}

func init() {
	Config = config.NewConfig()
	StorageConfig(Config)
	config.InitDB(DB)
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
