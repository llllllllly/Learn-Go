	package main

	import (
		"time"
		"os"
		"os/signal"
		"log"
		"net/http"
	)

	func main() {
		server := &http.Server{
			Addr: ":9090",
			WriteTimeout: 2 * time.Second,
		}
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		mux := http.NewServeMux()
		mux.Handle("/", &myHandler{})
		mux.HandleFunc("/hello", sayHello)
		server.Handler = mux

		go func() {
			<-quit

			if err := server.Close(); err != nil {
				log.Fatal("Close Server:", err)
			}
		}()

		log.Println("Start serve ....... v3")
		err := server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server closed under request")
			} else {
				log.Println("Server closed unexpected")
			}
		}
		log.Println("Server eixt")
	}

	type myHandler struct {
	}

	func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server version 3 Request URL is:" + r.URL.String()))
	}

	func sayHello(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ....... v3"))
	}