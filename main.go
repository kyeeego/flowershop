package main

import (
	"github.com/kyeeego/flowershop/config"
	"github.com/kyeeego/flowershop/delivery/gql"
	"github.com/kyeeego/flowershop/server"
	"log"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatalf("Unable to read environment variables: %e", err)
	}

	//dsn := url.URL{
	//	Scheme:   "postgres",
	//	User:     url.UserPassword(conf.DbUsername, conf.DbPassword),
	//	Host:     fmt.Sprintf("%s:%d", conf.DbHost, conf.DbPort),
	//	Path:     conf.DbName,
	//	RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	//}
	//db, err := gorm.Open("postgres", dsn.String())
	//if err != nil {
	//	log.Fatalf("Could not connect to database: %e\n", err)
	//}

	api, err := gql.New(gql.Customer{}, gql.Bouquet{}, gql.Purchase{}, gql.Seller{})
	if err != nil {
		log.Fatalf("Unable to initialize GraphQL API: %s", err.Error())
	}

	s := &server.Server{}
	if err := s.Run(conf.Port, api.Handler); err != nil {
		log.Fatalf("Got an error while trying to start the server: %s\n", err.Error())
	} else {
		log.Printf("Server is active on port :%d\n", conf.Port)
	}

}
