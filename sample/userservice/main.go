package main

import (
	"github.com/tnyidea/httpserver/go/httpserver"
	"github.com/tnyidea/httpserver/go/sample/userservice/data/models"
	"github.com/tnyidea/httpserver/go/sample/userservice/httpserver/config"
	usersrouter "github.com/tnyidea/httpserver/go/sample/userservice/httpserver/router"
	"log"
)

func main() {
	port := 8080

	ctx, err := config.NewContext()
	if err != nil {
		log.Fatal("fatal: Error initializing server: ", err)
	}

	router := httpserver.NewRouter()
	router = httpserver.AddDefaultRouter(router, ctx)
	router = usersrouter.AddApiV1UsersRouter(router, ctx)

	server := httpserver.HttpServer{
		Port:    port,
		Router:  router,
		Context: ctx,
	}
	defer func() {
		db := server.Context.Value(config.UserServiceContextDatabase).(models.DB)
		err := db.Close()
		if err != nil {
			log.Fatal("fatal: Error while exiting:", err)
		}
	}()

	log.Fatal(server.ListenAndServe())
}
