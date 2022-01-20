package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Credentials struct {
	Username string
	Password string
	Host     string
	Port     int
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

	return final
}

func ParseCredentials(credentials string) Credentials {

	creds := strings.Split(credentials, "\n")
	for i, e := range creds {
		creds[i] = strings.TrimSpace(e)
	}

	p, _ := strconv.ParseInt(creds[3], 10, 32)
	i := int(p)

	return Credentials{
		Username: creds[0],
		Password: creds[1],
		Host:     creds[2],
		Port:     i,
		Database: creds[4],
		Sslmode:  creds[5],
	}
}
