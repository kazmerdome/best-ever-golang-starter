package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"gitlab.com/kazmerdome/best-ever-golang-starter/cmd/simple-blog-graphql/graph"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/actor/db/mongodb"
	db "gitlab.com/kazmerdome/best-ever-golang-starter/internal/actor/db/sql"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/config"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/logger"
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
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/", echo.WrapHandler(playground.Handler("Best Ever Golang Starter 👊🏻 Simple Blog", "/query")))
	e.POST("/query", func(c echo.Context) error {
		req := c.Request()
		queryHandler.ServeHTTP(c.Response(), req)
		return nil
	})

	// Start server
	go func() {
		if err := e.Start(":" + "9099"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	// Stop server gracefully
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
