package utils

import (
  "io/ioutil"
  "strings"
  "strconv"
  "time"
  "math"
)

const(
  PROC_STAT = "/proc/stat"
)

type Cpus struct {
  UsagePerCpus	[]float64
  TotalUsage	float64
}

func ProcCpus() ([]string, error) {
  var (
    procsCpus []byte
    err	      error
    cpus      []string
  )

  if procsCpus, err = ioutil.ReadFile(PROC_STAT); err != nil {
    return cpus, err
  }

  cpus = strings.Split(string(procsCpus), "\n")

  return cpus[1:len(cpus) - 1], nil
}

func CpuUsagePercent() (Cpus, error) {
  var (
    cpusBefore	    []string
    cpusAfter	    []string
    err		    error
    icb		    []string
    ica		    []string
    percentUsage    []float64
    totalCpu	    = 0.0
  )

  if cpusAfter, err = ProcCpus(); err != nil {
    return Cpus{}, err
  }

  time.Sleep(1 * time.Second)

  if cpusBefore, err = ProcCpus(); err != nil {
    return Cpus{}, err
  }

  for i, _ := range cpusAfter {
    if strings.Contains(cpusAfter[i], "cpu") && strings.Contains(cpusBefore[i], "cpu") {
      var (
	idleA   int
	idleB   int
	systemA	int
	systemB	int
	idle	int
	total	int
	percent float64
      )

      ica = strings.Split(strings.TrimSpace(cpusAfter[i]), " ")
      icb = strings.Split(strings.TrimSpace(cpusBefore[i]), " ")

      idleA = stringToInt(ica[4]) + stringToInt(ica[5])
      idleB = stringToInt(icb[4]) + stringToInt(icb[5])

      systemA = stringToInt(ica[1]) + stringToInt(ica[2]) + stringToInt(ica[3]) + stringToInt(ica[6]) + stringToInt(ica[7]) + stringToInt(ica[8]) + idleA
      systemB = stringToInt(icb[1]) + stringToInt(icb[2]) + stringToInt(icb[3]) + stringToInt(icb[6]) + stringToInt(icb[7]) + stringToInt(icb[8]) + idleB

      idle = idleB - idleA
      total = systemB - systemA

      percent = 100 * (float64(total) - float64(idle)) / float64(total)
      totalCpu = totalCpu + percent
      percentUsage = append(percentUsage, toFixed(percent, 2))
    }
  }

  totalCpu = totalCpu / float64(len(percentUsage))
  //percentUsage = append(percentUsage, toFixed(totalCpu, 2))

  return Cpus{UsagePerCpus: percentUsage, TotalUsage: toFixed(totalCpu, 2)}, nil
}

func toFixed(value float64, fixed int) float64 {
  var (
    out	= math.Pow(10, float64(fixed))
    num	= int((value * out) + math.Copysign(0.5, (value * out)))
  )

  return float64(num) / out
}

func stringToInt(value string) int {
  var (
    i int
  )

  i, _ = strconv.Atoi(value)
  return i
}

func hexToDec(value string) (uint64, error) {
  var (
    dec uint64
    err error
  )

  if dec, err = strconv.ParseUint(value, 16, 32); err != nil {
    return dec, err
  }

  return dec, nil
}

