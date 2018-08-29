package main

import (
  "log"
  "testing"
  utils "gitlab-devops.totvs.com.br/lucas.martins/monitoring/utils"
)

func TestListPartitions(t *testing.T) {
  var (
    partitions	[]utils.Partitions
    err		error
  )

  if partitions, err = utils.ListPartitions(); err != nil {
    log.Fatalf("Error to list partitions: ", err)
  }

  log.Println(partitions)
}

func TestShowIODisk(t *testing.T) {
  var (
    io	[]utils.DiskIO
    err	error
  )

  for {
    if io, err = utils.IO(); err != nil {
      log.Fatalf("Error to get disk io: ", err)
    }


    for _, value := range io {
      if value.Device == "sda" {
	log.Println(value)
      }
    }
  }
}
