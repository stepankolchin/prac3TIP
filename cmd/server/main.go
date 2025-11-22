package main

import (
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"example.com/Prac3/internal/api"
	"example.com/Prac3/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
		})

	// Коллекция
	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	// Элемент
	mux.HandleFunc("GET /tasks/", h.GetTask)
	mux.HandleFunc("PATCH /tasks/", h.UpdateTask)
	mux.HandleFunc("DELETE /tasks/", h.DeleteTask)

	// Подключаем CORS
	handler := api.CORS(api.Logging(mux))

	server := &http.Server{
		Addr: ":8080",
		Handler: handler,
	}

	go func() {
		log.Println("listening on", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server could not have started", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server was shut by purpose or not:", err)
	}

	log.Println("Server exited properly")

}

