package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"

  "gopkg.in/yaml.v2"
)

type Redirect struct {
  To     string `yaml:"to"`
  Status int    `yaml:"status"`
}

type Config struct {
  Server struct {
    Address string `yaml:"address"`
    Port    int    `yaml:"port"`
  } `yaml:"server"`
  Redirects map[string]Redirect `yaml:"redirects"`
}

func main() {
  data, err := ioutil.ReadFile("config.yaml")
  if err != nil {
    log.Fatal(err)
  }
  var cfg Config
  if err := yaml.Unmarshal(data, &cfg); err != nil {
    log.Fatal(err)
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    key := r.Host + r.URL.Path
    if rd, ok := cfg.Redirects[key]; ok {
      http.Redirect(w, r, rd.To, rd.Status)
      return
    }
    if rd, ok := cfg.Redirects[r.Host]; ok {
      http.Redirect(w, r, rd.To, rd.Status)
      return
    }
    http.NotFound(w, r)
  })

  addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)
  log.Printf("Listening on %s", addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}
