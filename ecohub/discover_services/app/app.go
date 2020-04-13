package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/discover_services/app/handler"
	"github.com/discover_services/config"
	"go.mongodb.org/mongo-driver/mongo"
	_"github.com/discover_services/app/apdb"
	"time"
	"github.com/discover_services/logger"
	"log"
	"fmt"
	_"os"

)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *mongo.Client
}
var logs *logger.Logger

var accessLogs *logger.Logger

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {

	//Load Runner configuration
	config.LoadConf()
        logs, err := logger.Initialize(true)
        if err != nil {
                log.Fatalf("log Initialize failed")
		}

	logs = logs.GetLogger("App")

	//Initialize Access log
	config.CONF.Logger.FileLocation = config.CONF.Logger.AccessLog
	accessLogs, err = logger.Initialize(false)
	if err != nil {
		log.Fatalf("Access Logs Initialize failed")
	}	

	a.Router = mux.NewRouter()
	a.setRouters()
	a.Router.Use(RequestLoggingMiddleware)
	handler.Init()
	//apdb.Init()



}

// setRouters sets the all required routers
func (a *App) setRouters() {

	a.Get("/health", a.handleRequest(handler.NotImplemented))
	a.Post("/api/test_connectivity", a.handleRequest(handler.TestConnectivity))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run() {
	port := config.CONF.Port
	host := fmt.Sprintf("%s%d", ":", port)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// start := time.Now()
		// log.Printf("before Method: %s, URL: %s, Duration: %s", r.Method, r.URL, time.Since(start))
		// defer func() {
		// 	log.Printf("after Method: %s, URL: %s, Duration: %s", r.Method, r.URL, time.Since(start))
		// }()
		handler(w, r)
	}
}


// RequestLoggingMiddleware ...
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		alogs := accessLogs.GetLogger("Access")
		start := time.Now()
		alogs.Infof("Request Method: %s, URL: %s, Duration: %s", r.Method, r.URL, time.Since(start))
		
		defer func() {
			alogs.Infof("Response Method: %s, URL: %s, Duration: %s", r.Method, r.URL, time.Since(start))
		}()
		next.ServeHTTP(w, r)
	})
}

// func myLoggingHandler(h http.Handler) http.Handler {
// 	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 	  panic(err)
// 	}
// 	return handlers.LoggingHandler(logFile, h)
//   }
