package main

import (
  "log"
  "testing"
  utils "gitlab-devops.totvs.com.br/lucas.martins/monitoring/utils"
)

func TestListInterfaces(t *testing.T) {
  var (
    interfaces	[]utils.Interfaces
    err		error
  )

  if interfaces, err = utils.ListInterfaces(); err != nil {
    log.Fatalf("Error to list interfaces: ", err)
  }

  log.Println(interfaces)
}

func TestShowConsumeInterfaces(t *testing.T) {
  var (
    consume   []utils.Consume
    err	      error
    totalSent float64
    totalRecv float64
  )

  for {
    if consume, err = utils.ShowConsume(); err != nil {
      log.Fatalf("Error to get consume interfaces: ", err)
    }

    for _, value := range consume {
      if value.Interface == "ens33" {
	log.Println(value)

	totalSent = totalSent + value.Sent
	totalRecv = totalRecv + value.Recv
      }
    }

    log.Println("Total Sent: ", totalSent)
    log.Println("Total Recv: ", totalRecv)
  }
}
