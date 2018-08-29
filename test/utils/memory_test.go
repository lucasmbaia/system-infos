package main

import (
  "log"
  "testing"
  utils "gitlab-devops.totvs.com.br/lucas.martins/monitoring/utils"
)

func TestShowInfosMemory(t *testing.T) {
  var (
    memory  utils.Memory
    err	    error
  )

  if memory, err = utils.InfosMemory(); err != nil {
    log.Fatalf("Error to get infos of memory: ", err)
  }

  log.Println(memory)
}
