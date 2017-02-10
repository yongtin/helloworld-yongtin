package main

import (
  "fmt"
  "net/http"
  "github.com/urfave/negroni"
  "github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
  log.Formatter = new(logrus.JSONFormatter)
  log.Formatter = new(logrus.TextFormatter) // default

  // file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
  // if err == nil {
  //  log.Out = file
  // } else {
  //  log.Info("Failed to log to file, using default stderr")
  // }

  log.Level = logrus.DebugLevel
}

func main() {
  var port = 3000
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  n := negroni.Classic() // Includes some default middlewares
  n.UseHandler(mux)

  log.WithFields(logrus.Fields{
    "port": port,
  }).Info(fmt.Sprintf("Starting API"))

  http.ListenAndServe(fmt.Sprintf(":%d",port), n)
}
