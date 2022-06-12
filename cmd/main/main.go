package main

import (
	"MakeAnAPI/internal/cartoon"
	cartoon2 "MakeAnAPI/internal/cartoon/db"
	"MakeAnAPI/internal/config"
	"MakeAnAPI/pkg/client/postgesql"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	log.Println("create new client postgresql")
	postgreSQLClient, err := postgesql.NewClient(context.TODO(), cfg.Storage, 3)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Println("create repository new db")
	repository := cartoon2.NewDB(postgreSQLClient)

	log.Println("find all cartoons")
	_, err = repository.FindAll(context.TODO())
	if err != nil {
		log.Fatalf("Error :%v", err)
	}

	log.Println("register cartoon handler")
	handler := cartoon.NewHandler(repository)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			log.Fatal(err)
		}

		socketPath := path.Join(appDir, "app.sock")
		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindID, cfg.Listen.Port))
	}

	if listenErr != nil {
		log.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
