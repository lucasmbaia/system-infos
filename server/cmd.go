package server

import (
  "net"

  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  monitoring "gitlab-devops.totvs.com.br/lucas.martins/monitoring/monitoring"
)

const (
  port  = ":9090"
)

func Run() <-chan error {
  var (
    lis		      net.Listener
    err		      error
    s		      *grpc.Server
    errc	      = make(chan error, 2)
    monitoringServer  = NewMonitoringServer()
  )

  if lis, err = net.Listen("tcp", port); err != nil {
    errc <- err
    return errc
  }

  s = grpc.NewServer()
  //monitoring.RegisterMonitoringServiceServer(s, monitoringServer)
  monitoring.RegisterMonitoringServiceServer(s, monitoringServer)

  reflection.Register(s)

  go func() {
    errc <- s.Serve(lis)
  }()

  go func() {
    errc <- Gateway()
  }()

  return errc
}
