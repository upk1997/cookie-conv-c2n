package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type cookie struct {
	Domain            string
	IncludeSubdomains string
	Path              string
	HTTPOnly          string
	Expiration        int64
	Name              string
	Value             string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("[ERROR] No file input.")
	}
	fileContent, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatalln("[ERROR]", err.Error())
	}
	file, err := os.Create(fmt.Sprintf("output_%s.txt", time.Now().Format("20060102150405")))
	if err != nil {
		log.Fatalln("[ERROR]", err.Error())
	}
	defer file.Close()
	rawCookies := strings.Split(string(fileContent), "\n")
	_, errHeader := file.WriteString("# Netscape HTTP Cookie File\n")
	if errHeader != nil {
		log.Fatalln("[ERROR]", errHeader.Error())
	}
	for _, rawCookie := range rawCookies {
		splittedCookie := strings.Split(rawCookie, "\t")
		if len(splittedCookie) < 7 {
			continue
		}
		cookie := &cookie{
			Domain:            splittedCookie[2],
			IncludeSubdomains: "TRUE",
			Path:              splittedCookie[3],
			HTTPOnly:          "FALSE",
			Expiration:        0,
			Name:              splittedCookie[0],
			Value:             splittedCookie[1],
		}
		if cookie.Name == "" {
			continue
		}
		if strings.ToLower(splittedCookie[4]) == "session" {
			cookie.Expiration = time.Now().Unix() + 86400
		} else {
			expiration, err := time.Parse("2006-01-02T15:04:05.000Z", splittedCookie[4])
			if err != nil {
				fmt.Println("[ERROR]", err.Error())
				continue
			}
			cookie.Expiration = expiration.Unix()
		}
		if splittedCookie[6] == "✓" {
			cookie.HTTPOnly = "TRUE"
		}
		if !strings.HasPrefix(cookie.Domain, ".") {
			cookie.Domain = "." + cookie.Domain
		}
		_, err := file.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%d\t%s\t%s\n",
			cookie.Domain, cookie.IncludeSubdomains, cookie.Path, cookie.HTTPOnly, cookie.Expiration, cookie.Name, cookie.Value))
		if err != nil {
			log.Fatalln("[ERROR]", err.Error())
		}
		fmt.Printf("[COOKIE] %s\t%s\t%s\t%s\t%d\t%s\n",
			cookie.Domain, cookie.IncludeSubdomains, cookie.Path, cookie.HTTPOnly, cookie.Expiration, cookie.Name)
	}
}
