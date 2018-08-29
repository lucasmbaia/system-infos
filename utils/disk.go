package utils

import (
  "io/ioutil"
  "strings"
  "time"
  "strconv"

  disk "github.com/shirou/gopsutil/disk"
)

const(
  PROC_DISK = "/proc/diskstats"
)

type Partitions struct {
  Device	    string
  Mounted	    string
  FSType	    string
  Size		    uint64
  Free		    uint64
  Used		    uint64
  UsedPercent	    float64
  InodesTotal	    uint64
  InodeUsed	    uint64
  InodesFree	    uint64
  InodesUsedPercent float64
}

type DiskIO struct {
  Device	  string
  ReadsComplete	  float64
  ReadsMerged	  uint64
  SectorsRead	  uint64
  TimeSpentRead	  uint64
  WritesCompleted  float64
  WritesMerged	  uint64
  SectorsWrite	  uint64
  TimeSpentWrite  uint64
  IOs		  uint64
  TimeIO	  uint64
  WeightedTimeIO  uint64
  Wait		  float64
}

func ListPartitions() ([]Partitions, error) {
  var (
    partitions	[]disk.PartitionStat
    err		error
    p		[]Partitions
    usage	*disk.UsageStat
  )

  if partitions, err = disk.Partitions(false); err != nil {
    return p, err
  }

  for _, partition := range partitions {
    if usage, err = disk.Usage(partition.Mountpoint); err != nil {
      return p, err
    }

    p = append(p, Partitions{
      Device:		  partition.Device,
      Mounted:		  partition.Mountpoint,
      FSType:		  partition.Fstype,
      Size:		  (usage.Total / 1024) / 1024,
      Free:		  (usage.Free / 1024) / 1024,
      Used:		  (usage.Used / 1024) / 1024,
      UsedPercent:	  toFixed(usage.UsedPercent, 2),
      InodesTotal:	  (usage.InodesTotal / 1024) / 1024,
      InodeUsed:	  (usage.InodesUsed / 1024) / 1024,
      InodesFree:	  (usage.InodesFree / 1024) / 1024,
      InodesUsedPercent:  toFixed(usage.InodesUsedPercent, 2),
    })
  }

  return p, nil
}

func IO() ([]DiskIO, error) {
  var (
    ioAfter   []string
    ioBefore  []string
    err	      error
    after     []string
    before    []string
    diskIO    []DiskIO
    write     float64
    read      float64
    //wait      float64
  )

  if ioAfter, err = procDisk(); err != nil {
    return []DiskIO{}, err
  }

  time.Sleep(1 * time.Second)

  if ioBefore, err = procDisk(); err != nil {
    return []DiskIO{}, err
  }

  for i, _ := range ioBefore {
    after = strings.Split(strings.TrimSpace(ioAfter[i]), " ")
    before = strings.Split(strings.TrimSpace(ioBefore[i]), " ")

    write = (stringToFloat(before[15]) - stringToFloat(after[15])) * 512
    read = (stringToFloat(before[11]) - stringToFloat(after[11])) * 512

    if write > 0 {
      write = toFixed(write / 1024, 2)
    }

    if read > 0 {
      read = toFixed(read / 1024, 2)
    }

    diskIO = append(diskIO, DiskIO{
      Device:		after[8],
      ReadsComplete:	read,
      ReadsMerged:	stringToUint(before[10]) - stringToUint(after[10]),
      SectorsRead:	stringToUint(before[11]) - stringToUint(after[11]),
      TimeSpentRead:	stringToUint(before[12]) - stringToUint(after[12]),
      WritesCompleted:	write,
      WritesMerged:	stringToUint(before[14]) - stringToUint(after[14]),
      SectorsWrite:	stringToUint(before[15]) - stringToUint(after[15]),
      TimeSpentWrite:	stringToUint(before[16]) - stringToUint(after[16]),
      IOs:		stringToUint(before[17]) + stringToUint(after[17]),
      TimeIO:		stringToUint(before[18]) - stringToUint(after[18]),
      WeightedTimeIO:	stringToUint(before[19]) - stringToUint(after[19]),
    })

    if write > 0 || read > 0 {
      //diskIO.Wait = (float64(diskIO.TimeSpentRead) + float64(diskIO.TimeSpentWrite)) / ((write / 1024) + (read / 1024))
    }
  }

  return diskIO, nil
}

func procDisk() ([]string, error) {
  var (
    procsDisk []byte
    err	      error
    disks     []string
  )

  if procsDisk, err = ioutil.ReadFile(PROC_DISK); err != nil {
    return disks, err
  }

  disks = strings.Split(string(procsDisk), "\n")

  return disks[:len(disks) - 1], nil
}

func stringToUint(v string) uint64 {
  var (
    u uint64
  )

  u, _ = strconv.ParseUint(v, 10, 64)

  return u
}
