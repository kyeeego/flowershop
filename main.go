package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kyeeego/flowershop/delivery/gql"
	"github.com/kyeeego/flowershop/repository"
	"github.com/kyeeego/flowershop/server"
	"github.com/kyeeego/flowershop/service"
	"log"
	"net/url"
)

func main() {
	//conf, err := config.Init()
	//if err != nil {
	//	log.Fatalf("Unable to read environment variables: %e", err)
	//}

	dsn := url.URL{
		Scheme: "postgres",
		//User:     url.UserPassword(conf.DbUsername, conf.DbPassword),
		//Host:     fmt.Sprintf("%s:%d", conf.DbHost, conf.DbPort),
		//Path:     conf.DbName,
		User:     url.UserPassword("admin", "admin"),
		Host:     "localhost:5432",
		Path:     "upsflower",
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
	if err := s.Run(8081, api.Handler); err != nil {
		log.Fatalf("Got an error while trying to start the server: %s\n", err.Error())
	} else {
		log.Printf("Server is active on port :%d\n", 8081)
	}

}
