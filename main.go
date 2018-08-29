package main

import (
  "log"
  "gitlab-devops.totvs.com.br/lucas.martins/monitoring/server"
)

func main() {
  err := server.Run()

  select {
  case e := <-err:
    log.Fatal(e)
  }
}
