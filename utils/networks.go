package utils

import (
  //"log"
  "strconv"
  "strings"
  "time"

  net "github.com/shirou/gopsutil/net"
)

type Interfaces struct {
  Name	string
  Mtu	int
  Mac	string
  Addr	string
  Mask	string
}

type Consume struct {
  Interface   string
  Sent	      float64
  Recv	      float64
  PacketsSent uint64
  PacketsRecv uint64
  ErrorRecv   uint64
  ErrorSent   uint64
  DropRecv    uint64
  DropSent    uint64
}

func ListInterfaces() ([]Interfaces, error) {
  var (
    interfaces	[]net.InterfaceStat
    err		error
    addr	[]string
    addrs	[]Interfaces
  )

  if interfaces, err = net.Interfaces(); err != nil {
    return []Interfaces{}, err
  }

  for _, value := range interfaces {
    if len(value.Addrs) > 0 {
      addr = strings.Split(value.Addrs[0].Addr, "/")
    } else {
      addr = []string{"", ""}
    }

    addrs = append(addrs, Interfaces{
      Name: value.Name,
      Mtu:  value.MTU,
      Mac:  value.HardwareAddr,
      Addr: addr[0],
      Mask: addr[1],
    })
  }

  return addrs, nil
}

func ShowConsume() ([]Consume, error) {
  var (
    IOAfter   []net.IOCountersStat
    IOBefore  []net.IOCountersStat
    err	      error
    consume   []Consume
    sent      float64
    recv      float64
  )

  if IOBefore, err = net.IOCounters(true); err != nil {
    return []Consume{}, err
  }

  time.Sleep(1 * time.Second)

  if IOAfter, err = net.IOCounters(true); err != nil {
    return []Consume{}, err
  }

  for _, before := range IOBefore {
    for _, after := range IOAfter {
      if before.Name == after.Name {
	sent = (float64(after.BytesSent) - float64(before.BytesSent))
	recv = (float64(after.BytesRecv) - float64(before.BytesRecv))

	if sent > 0 {
	  sent = toFixed(sent / 1024, 2)
	}

	if recv > 0 {
	  recv = toFixed(recv / 1024, 2)
	}

	consume = append(consume, Consume{
	  Interface:	before.Name,
	  Sent:		sent,
	  Recv:		recv,
	  PacketsSent:	after.PacketsSent - before.PacketsSent,
	  PacketsRecv:	after.PacketsRecv - before.PacketsRecv,
	  ErrorRecv:	after.Errin - before.Errin,
	  ErrorSent:	after.Errout - before.Errout,
	  DropRecv:	after.Dropin - before.Dropin,
	  DropSent:	after.Dropout - before.Dropout,
	})
      }
    }
  }

  return consume, nil
}

func stringToFloat(v string) float64 {
  var (
    f float64
  )

  f, _ = strconv.ParseFloat(v, 64)

  return f
}
