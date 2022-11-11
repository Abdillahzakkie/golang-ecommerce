package helpers

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	reg := regexp.MustCompile(`^(.*` + "golang-ecommerce" + `)`)
	cwd, _ := os.Getwd()
	rootPath := reg.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}