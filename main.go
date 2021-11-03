package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kyeeego/flowershop/config"
	"github.com/kyeeego/flowershop/delivery/gql"
	"github.com/kyeeego/flowershop/repository"
	"github.com/kyeeego/flowershop/server"
	"github.com/kyeeego/flowershop/service"
	"log"
	"net/url"
	"strconv"
)

func main() {

	conf := config.Init()

	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(conf.DBUsername, conf.DBPassword),
		Host:     fmt.Sprintf("%s:%s", conf.DBHost, conf.DBPort),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err.Error())
	}

	repos := repository.New(db)
	services := service.New(repos)

	api, err := gql.New(services, &gql.Customer{}, &gql.Bouquet{}, &gql.Purchase{}, &gql.Seller{})
	if err != nil {
		log.Fatalf("Unable to initialize GraphQL API: %s", err.Error())
	}

	s := &server.Server{}
	port, _ := strconv.Atoi(conf.Port)
	if err := s.Run(port, api.Handler); err != nil {
		log.Fatalf("Got an error while trying to start the server: %s\n", err.Error())
	}
}
