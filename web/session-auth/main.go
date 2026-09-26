package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"session-auth/config"
	"session-auth/controller"
	"session-auth/entity"
	"session-auth/pkg/db"
	"session-auth/repo/persistent"
	"session-auth/service/user"
	"session-auth/view"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error load config: %v", err)
	}
	database, err := db.NewDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("error new db: %v", err)
	}
	err = database.AutoMigrate(entity.User{}, entity.Session{})
	if err != nil {
		log.Fatalf("error auto migrate db: %v", err)
	}
	v, err := view.New()
	if err != nil {
		log.Fatalf("error load templates: %v", err)
	}
	userRepo := persistent.NewUserRepo(database)
	sessionRepo := persistent.NewSessionRepo(database)
	userService := user.New(userRepo, sessionRepo)
	router := controller.NewRouter(cfg, userService, v)
	mux := http.NewServeMux()
	router.Register(mux)
	svr := &http.Server{Addr: cfg.Addr, Handler: controller.Sleep(mux)}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server: %v", err)
		}
	}()
	<-ctx.Done()
	stop()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := svr.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
	}
	if sqlDB, err := database.DB(); err == nil {
		sqlDB.Close()
	}
}
