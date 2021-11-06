package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"net"
	"os"
	"path/filepath"
	"time"
)

type ServerInfo struct {
	CpuMode string `json:"cpu_mode"`
	CpuCores int32 `json:"cpu_cores"`
	MemTotal uint64 `json:"mem_total"`
	MemUsed uint64 `json:"mem_used"`
	MemFree uint64 `json:"mem_free"`
	HostName string `json:"host_name"`
	HostPlatform string `json:"host_platform"`
	DiskPath string `json:"disk_path"`
	DiskTotal uint64 `json:"disk_total"`
	DiskFree uint64 `json:"disk_free"`
	DiskUsed uint64 `json:"disk_used"`
	DiskUsedPercent float64 `json:"disk_used_percent"`
	HostIp net.IP `json:"host_ip"`
	StartTime string `json:"start_time"`
}


// 服务器信息
func PostServerInfo(ctx *gin.Context) {
	var serverInfo ServerInfo
	cpuInfo,_ := cpu.Info()
	//for _,info := range cpuInfo {
	//	data,_ := json.MarshalIndent(info,""," ")
	//	fmt.Print(string(data))
	//}
	memInfo,_ := mem.VirtualMemory()
	hostInfo,_ := host.Info()
	exePath,_ := os.Executable()
	runPath,_ := filepath.EvalSymlinks(filepath.Dir(exePath))
	diskInfo,_ := disk.Usage(runPath)
	//conn,_ := net.Dial("udp","8.8.8.8:80")
	//defer conn.Close()
	//localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Printf("ip:%v\n",localAddr.String())
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		serverInfo.HostIp = ipAddr.IP.To4()
	}
	startTimestamp,_ := host.BootTime()
	startTime := time.Unix(int64(startTimestamp), 0)
	serverInfo.StartTime = startTime.Local().Format("2006-01-02 15:04:05")
	serverInfo.CpuMode = cpuInfo[0].ModelName
	serverInfo.CpuCores = cpuInfo[0].Cores
	serverInfo.MemTotal = memInfo.Total
	serverInfo.MemFree = memInfo.Free
	serverInfo.MemUsed = memInfo.Used
	serverInfo.HostName = hostInfo.Hostname
	serverInfo.HostPlatform = hostInfo.Platform
	serverInfo.DiskPath = runPath
	serverInfo.DiskTotal = diskInfo.Total
	serverInfo.DiskFree = diskInfo.Free
	serverInfo.DiskUsed = diskInfo.Used
	serverInfo.DiskUsedPercent = diskInfo.UsedPercent
	common.Success(ctx,serverInfo)
}
