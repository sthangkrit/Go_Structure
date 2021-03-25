package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"Go_Structure/handler"

	// log "github.com/sirupsen/logrus"
)

func main() {

	port := flag.String("port", "8080", "port number")
	configPath := flag.String("config", "configure", "set configs path, default as: 'configure'")

	flag.Parse()

	fmt.Println("port : %+v", *port)
	fmt.Println("configPath directory : %+v", *configPath)

	//connect database
	// InitConnectionDatabase(*configPath)

	//start http server
	r := handler.Routes{} //new object
	handleRoute := r.InitTransactionRoute()
	srv := &http.Server{
		Addr:    fmt.Sprint(":", *port), //":8080"
		Handler: handleRoute,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("transaction listen: %s\n", err)

		} else if err != nil {
			fmt.Println("transaction listen error: %s\n", err)
			
		}
		fmt.Println("transaction listen at: %s", *port)
	}()

	// create channel wait signals
	// จับสัญญาณ ctr+C
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	handler.StartScheduler()
	<-signals // wait for SIGINT
}

// func InitConnectionDatabase(configPath string) {
// 	mssql.InitDB(configPath)
// }
