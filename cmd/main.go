package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"forum/app/config"
	"forum/app/handlers"
	"forum/app/repository"
	"forum/app/service/post"
	"forum/app/service/session"
	"forum/app/service/user/auth"
	"forum/app/service/user/user"
)

func main() {
	cfg, err := config.InitConfig("./config/config.json")
	if err != nil {
		log.Fatalln(err)
		return
	}
	db, err := repository.NewDB(cfg.Database)
	if err != nil {
		log.Fatalln(err)
		return
	}

	repo := repository.NewRepo(db)
	authservice := auth.NewAuthService(repo)
	userservice := user.NewUserService(repo)

	sessionService := session.NewSessionService(repo)

	postService := post.NewPostService(repo)

	app := handlers.NewAppService(authservice, sessionService, postService, userservice, cfg)
	server := app.Run(cfg.Http)

	go app.ClearSession()

	go func() {
		log.Printf("server started at %s", cfg.ServerAddress)
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("listen %s ", err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down servers ...")
	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("server shut down:%s", err)
	}
	log.Println("server stopped")
}
