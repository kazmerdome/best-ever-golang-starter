package main

import (
	"os"
	"os/signal"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kazmerdome/best-ever-golang-starter/cmd/simple-blog-graphql/graph"
	"github.com/kazmerdome/best-ever-golang-starter/internal/actor/db/mongodb"
	db "github.com/kazmerdome/best-ever-golang-starter/internal/actor/db/sql"
	"github.com/kazmerdome/best-ever-golang-starter/internal/actor/server"
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	"github.com/kazmerdome/best-ever-golang-starter/internal/util/config"
	"github.com/kazmerdome/best-ever-golang-starter/internal/util/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	// Init Config
	//
	c := config.NewConfig()
	err := c.LoadConfigFile(".", "env", ".env")
	if err != nil {
		log.
			Error().
			Msg(err.Error())
	}

	// Init Logger
	//
	logger.InitLogger(c.GetString("LOG_LEVEL"), c.GetString("ENVIRONMENT"))

	// Init Databases
	//
	mdb := mongodb.
		NewMongodb(
			c.GetString("MONGO_URI"),
			c.GetString("MONGO_DATABASE"),
			c.GetBool("MONGO_RETRYWRITES"),
		).
		Connect()
	defaultMdb := mdb.GetDatabase()
	defer mdb.Disconnect()

	pdb := db.NewPostgresDB(
		c.GetString("POSTGRES_DATABASE"),
		c.GetString("POSTGRES_URI"),
		c.GetBool("POSTGRES_IS_SSL_DISABLED"),
	).Connect()
	defer pdb.Disconnect()

	// Initialize Domain Modules
	//
	categoryModule := category.NewCategoryModule(defaultMdb)
	postModule := post.NewPostModule(pdb)

	// Init Gql
	//
	resolver := graph.Resolver{
		CategoryService: categoryModule.GetService(),
		PostService:     postModule.GetService(),
	}
	config := graph.Config{Resolvers: &resolver}
	queryHandler := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	// Init Http Server
	//
	es := server.NewEchoServer(9099)
	so := es.GetOperations()
	so.UseRecover()
	so.UseCors()
	so.Get("/", playground.Handler("Best Ever Golang Starter 👊🏻 Simple Blog", "/query"))
	so.Post("/query", queryHandler.ServeHTTP)

	// Start server
	go es.Start()
	// Stop server gracefully
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	es.Stop()
}
