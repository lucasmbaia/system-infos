package server

import (
  monitoring "gitlab-devops.totvs.com.br/lucas.martins/monitoring/monitoring"
  utils "gitlab-devops.totvs.com.br/lucas.martins/monitoring/utils"
)

type MonitoringServer struct {}

func NewMonitoringServer() *MonitoringServer {
  return new(MonitoringServer)
}

func (m *MonitoringServer) CpuInfos(emp *monitoring.Empty, infos monitoring.MonitoringService_CpuInfosServer) error {
  var (
    cpus  utils.Cpus
    err	  error
  )

  for {
    if cpus, err = utils.CpuUsagePercent(); err != nil {
      return err
    }

    if err = infos.Send(&monitoring.InfosCpu{PerCpu: cpus.UsagePerCpus, TotalCpu: cpus.TotalUsage}); err != nil {
      return err
    }
  }

  return nil
}

func (m *MonitoringServer) ConsumeInterfaces(emp *monitoring.Empty, infos monitoring.MonitoringService_ConsumeInterfacesServer) error {
  var (
    netConsume	[]utils.Consume
    err		error
    consume	monitoring.ShowConsume
  )

  for {
    consume = monitoring.ShowConsume{}

    if netConsume, err = utils.ShowConsume(); err != nil {
      return err
    }

    for _, value := range netConsume {
      consume.Consume = append(consume.Consume, &monitoring.InfosConsume{
	Interface:    value.Interface,
	Sent:	      value.Sent,
	Recv:	      value.Recv,
	PacketsSent:  value.PacketsSent,
	PacketsRecv:  value.PacketsRecv,
	ErrorRecv:    value.ErrorRecv,
	ErrorSent:    value.ErrorSent,
	DropRecv:     value.DropRecv,
	DropSent:     value.DropSent,
      })
    }

    if err = infos.Send(&consume); err != nil {
      return err
    }
  }

  return nil
}

func (m *MonitoringServer) ConsumeDisk(emp *monitoring.Empty, infos monitoring.MonitoringService_ConsumeDiskServer) error {
  var (
    diskConsume	[]utils.DiskIO
    err		error
    consume	monitoring.ShowConsumeDisk
  )

  for {
    consume = monitoring.ShowConsumeDisk{}

    if diskConsume, err = utils.IO(); err != nil {
      return err
    }

    for _, value := range diskConsume {
      consume.DiskIo = append(consume.DiskIo, &monitoring.InfosConsumeDisk{
	Device:		  value.Device,
	ReadsComplete:	  value.ReadsComplete,
	ReadsMerged:	  value.ReadsMerged,
	SectorsRead:	  value.SectorsRead,
	TimeSpentRead:	  value.TimeSpentRead,
	WritesCompleted:  value.WritesCompleted,
	WritesMerged:	  value.WritesMerged,
	SectorsWrite:	  value.SectorsWrite,
	TimeSpentWrite:	  value.TimeSpentWrite,
      })
    }

    if err = infos.Send(&consume); err != nil {
      return err
    }
  }

  return nil
}

/*func (m *MonitoringServer) AllInfos(emp *monitoring.Empty, infos monitoring.MonitoringService_AllInfosServer) error {
  return nil
}*/
