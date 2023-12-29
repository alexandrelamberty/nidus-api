package config

import (
	"errors"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

var rex = regexp.MustCompile("(\\w+)=(\\w+)")

func CheckConfig() error {
	// 	conn := `ENV=dev
	// DATABASE_URI=mongodb://nidus:nidus@localhost:27017/nidus
	// PAIRING_KEY=9fca54477c8ad4e70dc5e1084f884aad
	// BCRYPT_HASH=7f91317e30a02bc7b87205e95b842df2
	// JWT_SECRET=d7a481461577ba4c3c4c6946cca7204b
	// JWT_EXPIRE=90`
	// 	data := rex.FindAllStringSubmatch(conn, -1)

	// 	res := make(map[string]string)
	// 	for _, kv := range data {
	// 		k := kv[1]
	// 		v := kv[2]
	// 		res[k] = v
	// 	}
	// 	fmt.Println(res)

	// Environments variables
	_, exist := os.LookupEnv("DATABASE_URI")
	if !exist {
		log.Println("[Config] No system environment variables found !")
		log.Println("[Config] Looking for environment variables files ...")
		err := godotenv.Load(".env")
		if err != nil {
			return errors.New("No environment variables files found !")
		} else {
			log.Println("[Config] Environment variables file found !")
			env := os.Getenv("DATABASE_URI")
			if env == "" {
				return errors.New("Environment not found !!")
			}
		}
	}
	return nil

}
