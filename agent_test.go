package main

import (
    "os"
    "testing"
)

func TestNewProcData(t *testing.T) {
    i := uint64(0)
    got := NewProcData(0.0, &i, 0)

    if got.io_wait != float64(0.0) {
        t.Errorf("NewProcData(0, 0, 0) = %f; want 0.0", got.io_wait)
    }

    if got.memory != &i {
        t.Errorf("NewProcData(0, 0, 0) = %d; want 0", got.memory)
    }

    if got.processes != 0 {
        t.Errorf("NewProcData(0, 0, 0) = %d; want 0", got.processes)
    }
}

func TestgetProc(t *testing.T) {
    got1 := getProc()

    if got1 != "/proc/" {
        t.Errorf("getProc() = %s; want /proc/", got1)
    }

    m := "/host/mounted/proc"
    os.Setenv("PROMETHEUS_AGENT_PROC", m)

    got2 := getProc()
    if got2 != m {
        t.Errorf("getProc() = %s; want %s", got2, m)
    }

}
