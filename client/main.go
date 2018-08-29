package main

import (
  "log"
  "io"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  monitoring "gitlab-devops.totvs.com.br/lucas.martins/monitoring/monitoring"
)

func main() {
  var (
    conn		      *grpc.ClientConn
    c			      monitoring.MonitoringServiceClient
    err			      error
    //responseCpu		      monitoring.MonitoringService_CpuInfosClient
    //responseConsumeInterfaces monitoring.MonitoringService_ConsumeInterfacesClient
    responseConsumeDisk	      monitoring.MonitoringService_ConsumeDiskClient
  )

  // Set up a connection to the server.
  if conn, err = grpc.Dial("localhost:9090", grpc.WithInsecure()); err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()

  c = monitoring.NewMonitoringServiceClient(conn)

  /*if responseCpu, err = c.CpuInfos(context.Background(), &monitoring.Empty{}); err != nil {
    log.Fatal(err)
  }

  for {
    cpu, err := responseCpu.Recv()

    if err == io.EOF {
      break
    }

    log.Println(cpu)
  }*/

  /*if responseConsumeInterfaces, err = c.ConsumeInterfaces(context.Background(), &monitoring.Empty{}); err != nil {
    log.Fatal(err)
  }

  for {
    consume, err := responseConsumeInterfaces.Recv()

    if err == io.EOF {
      break
    }

    for _, value := range consume.Consume {
      if value.Interface == "ens37" {
	log.Println(value)
      }
    }
  }*/

  if responseConsumeDisk, err = c.ConsumeDisk(context.Background(), &monitoring.Empty{}); err != nil {
    log.Fatal(err)
  }

  for {
    consume, err := responseConsumeDisk.Recv()

    if err == io.EOF {
      break
    }

    for _, value := range consume.DiskIo {
      if value.Device == "sda" {
	log.Println(value)
      }
    }
  }
}
