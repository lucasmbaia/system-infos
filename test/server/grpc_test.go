package main

import (
  "log"
  "testing"
  monitoring "gitlab-devops.totvs.com.br/lucas.martins/monitoring/monitoring"
  server "gitlab-devops.totvs.com.br/lucas.martins/monitoring/server"
)

func TestStremInfosServer(t *testing.T) {
  var (
    s	  = server.NewMonitoringServer()
    err	  error
    infos monitoring.MonitoringService_AllInfosServer
  )

  if err = s.AllInfos(nil, infos); err != nil {
    log.Fatal(err)
  }
}
