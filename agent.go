package main

import (
    "fmt"
    "github.com/prometheus/procfs"
    "github.com/gin-gonic/gin"
    "os"
)

const Proc = "/proc/"
const Defaultlisten = "127.0.0.1:9091"

type ProcData struct {
    io_wait float64
    memory *uint64
    processes int
}

func NewProcData(w float64, m *uint64, p int) *ProcData {
    z := ProcData{io_wait: w, memory: m, processes: p}

    return &z
}

func getProc() string {
    p := os.Getenv("PROMETHEUS_AGENT_PROC")

    if len(p) > 0 {
        return p
    }

    return Proc
}

func getListen() string {
    p := os.Getenv("PROMETHEUS_AGENT_LISTEN")

    if len(p) > 0 {
        return p
    }

    return Defaultlisten

}

func gatherProcData() (*ProcData, error) {
    fs, err := procfs.NewFS(getProc())
    p, err := fs.AllProcs()
    s, err := fs.Stat()
    m, err := fs.Meminfo()

    if err != nil {
        mem := uint64(0)
        return NewProcData(0.0, &mem, 0), err
    }

    return NewProcData(s.CPUTotal.Iowait, m.MemFree, p.Len()), err

}

func main() {
    r := gin.Default()

    r.GET("/metrics", func(c *gin.Context) {
        gpd, err := gatherProcData()

        if err != nil {
            c.JSON(500, gin.H{
                "error": true,
                "message": fmt.Sprintf(
                    "Application error: %s",
                    err.Error(),
                ),
            })

            return
        }

        c.JSON(200, gin.H{
            "error": false,
            "result": gin.H{
                "free_memory_kb": gpd.memory,
                "io_wait_pct": gpd.io_wait,
                "running_process_count": gpd.processes,
            },
        })
    })

    r.Run(getListen())

}
