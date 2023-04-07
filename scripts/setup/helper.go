package setup

import (
	"fmt"
	"os"
)

func FailOnError(action string, err error) {
	if err != nil {
		fmt.Printf("error on %s\n", action)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func StartService(s ServerProcess) {
	FailOnError(fmt.Sprintf("stopping %s", s.Name), KillProcessesListeningOnPort(s.Port))
	FailOnError(fmt.Sprintf("starting %s", s.Name), RunCommandInBackground(s.Name, s.Command))
	FailOnError(fmt.Sprintf("waiting for %s to be up", s.Name), AwaitHttpServiceStartup(s.HealthEndpoint))
}
func StopService(s ServerProcess) {
	FailOnError(fmt.Sprintf("stopping %s", s.Name), KillProcessesListeningOnPort(s.Port))
}
