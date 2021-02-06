package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {

  started := time.Now()
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(200)
      w.Write([]byte("Hey Now, this is Manish's ContainerService\n"))
  })

  http.HandleFunc("/started", func(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(200)
      data := (time.Now().Sub(started)).String()
      w.Write([]byte(data + "\n"))
  })

  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
      duration := time.Now().Sub(started)
      if duration.Seconds() > 10 {
          w.WriteHeader(500)
          w.Write([]byte(fmt.Sprintf("error: %v\n", duration.Seconds())))
      } else {
          w.WriteHeader(200)
          w.Write([]byte("ok\n"))
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
