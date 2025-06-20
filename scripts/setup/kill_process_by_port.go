package setup

import (
	"errors"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"time"
)

const shutdownRetries = 20
const shutdownRetryWaitTime = 50 * time.Millisecond

func KillProcessesListeningOnPort(port uint32) error {
	for i := 0; i < shutdownRetries; i++ {
		if running, err := hasPidListeningOnPort(port); err != nil {
			return err
		} else if !running {
			return nil
		}

		pids, err := pidsListeningOnPort(port)
		if err != nil {
			return err
		}

		killPids(pids)

		time.Sleep(shutdownRetryWaitTime)
	}

	return errors.New("could not shutdown process")
}

func hasPidListeningOnPort(port uint32) (bool, error) {
	if stillRunning, err := pidsListeningOnPort(port); err != nil {
		return false, err
	} else {
		return len(stillRunning) > 0, nil
	}

}

func pidsListeningOnPort(port uint32) ([]int32, error) {
	connections, err := net.Connections("tcp")
	if err != nil {
		return nil, err
	}

	var result []int32
	for _, conn := range connections {
		if conn.Status != "LISTEN" {
			continue
		}
		if conn.Laddr.Port != port {
			continue
		}
		if conn.Laddr.IP == "0.0.0.0" || conn.Laddr.IP == "127.0.0.1" || conn.Laddr.IP == "::" || conn.Laddr.IP == "*" {
			result = append(result, conn.Pid)
		}
	}

	return result, nil
}

func killPids(pids []int32) {
	for _, p := range pids {
		proc, err := process.NewProcess(p)
		if err != nil {
			continue
		}

		_ = proc.Kill()
	}
}
