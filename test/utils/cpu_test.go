package main

import (
  "log"
  "testing"
  utils "gitlab-devops.totvs.com.br/lucas.martins/monitoring/utils"
)

func TestShowInfosCpusOfFileStat(t *testing.T) {
  var (
    //procs []string
    err	  error
  )

  if _, err = utils.ProcCpus(); err != nil {
    log.Fatalf("Error to list procs of file /proc/stat: ", err)
  }
}

func TestShowCpusUsagePercent(t *testing.T) {
  var (
    percent utils.Cpus
    err	    error
  )

  for {
    if percent, err = utils.CpuUsagePercent(); err != nil {
      log.Fatalf("Error to get usage percent cpu: ", err)
    }
  log.Println(percent)
  }


}
