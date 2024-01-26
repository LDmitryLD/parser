package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/LDmitryLD/parser/app/config"
	"projects/LDmitryLD/parser/app/internal/db"
	"projects/LDmitryLD/parser/app/internal/infrastructure/router"
	"projects/LDmitryLD/parser/app/internal/modules"
	"projects/LDmitryLD/parser/app/internal/parser"
	"projects/LDmitryLD/parser/app/internal/storages"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

// TODO: сделать type switch, проверить обработку ошибок, попробовать заменить селениум парсер на интерфейс, почистить код

func main() {
	confDB := config.NewAppConfig().DB

	_, sqlAdapter, err := db.NewSqlDB(confDB)
	if err != nil {
		log.Fatal("ошибка при подключении к БД", err)

	}

	parser := parser.NewParser()

	storages := storages.NewStorages(sqlAdapter)

	services := modules.NewServices(storages, parser)

	controllers := modules.NewControllers(services)

	r := router.NewRouter(controllers)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
