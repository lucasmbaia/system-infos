package utils

import (
  mem "github.com/shirou/gopsutil/mem"
)

type Memory struct {
  Total	      uint64
  Free	      uint64
  Available   uint64
  Used	      uint64
  Cached      uint64
  UsedPercent float64
}

func InfosMemory() (Memory, error) {
  var (
    infos   *mem.VirtualMemoryStat
    err	    error
    memory  Memory
  )

  if infos, err = mem.VirtualMemory(); err != nil {
    return memory, err
  }

  memory = Memory{
    Total:	  (infos.Total / 1024) / 1024,
    Free:	  (infos.Free / 1024) / 1024,
    Available:	  (infos.Available / 1024) / 1024,
    Used:	  (infos.Used / 1024) / 1024,
    UsedPercent:  toFixed(infos.UsedPercent, 2),
    Cached:	  (infos.Cached / 1024) / 1024,
  }

  return memory, nil
}
