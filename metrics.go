package main

import (
	"os/exec"
)

type Metrics struct {
	CPU  string
	RAM  string
	Disk string
}

func (m *Metrics) UpdateAll() *Metrics {
	m.GetCPU()
	m.GetRAM()
	m.GetDisk()

	return m
}

func (m *Metrics) GetCPU() string {
	usage, _ := exec.Command("/bin/sh", "-c", "grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print usage}'").Output()
	m.CPU = string(usage)
	return m.CPU
}

func (m *Metrics) GetRAM() string {
	usage, _ := exec.Command("/bin/sh", "-c", "free | grep Mem | awk '{print $3/$2 * 100.0}'").Output()
	m.RAM = string(usage)
	return m.RAM
}

func (m *Metrics) GetDisk() string {
	usage, _ := exec.Command("/bin/sh", "-c", "df --output=pcent / | tr -dc '0-9'").Output()
	m.Disk = string(usage)
	return m.Disk
}
