package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Credentials struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Sslmode  string
}

var CredentialsInstance Credentials = ParseCredentials(ReadCredentials())

func ReadCredentials() string {
	file, ferr := os.Open("cred/info.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	final := ""

	for scanner.Scan() {
		line := scanner.Text()
		final += "\n" + line
	}

	return final[1:]
}

func ParseCredentials(credentials string) Credentials {

	creds := strings.Split(credentials, "\n")
	for i, e := range creds {
		creds[i] = strings.SplitAfter(e, "= ")[1]
	}

	return Credentials{
		Username: creds[0],
		Password: creds[1],
		Host:     creds[2],
		Port:     creds[3],
		Database: creds[4],
		Sslmode:  creds[5],
	}
}

func MapCredentials(credentials Credentials) string {
	return fmt.Sprintf(
		"host=%s user=%s password='%s' dbname=%s port=%s sslmode=%s TimeZone=America/New_York",
		credentials.Host, credentials.Username, credentials.Password, credentials.Database, credentials.Port, credentials.Sslmode,
	)
}
