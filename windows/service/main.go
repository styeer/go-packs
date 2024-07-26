package main

import (
	"golang.org/x/sys/windows/svc"
)

func main() {

}

func RunService(name string, start, stop func() error) error {
	if flag, _ := svc.IsWindowsService(); flag {
		return svc.Run(name, &WinService{Start: start, Stop: stop})
	}
	return nil
}

type WinService struct {
	Start func() error
	Stop  func() error
}

func (ws *WinService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}

	if err := ws.Start(); err != nil {
		return true, 1
	}

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

loop:
	for {
		c := <-r
		switch c.Cmd {
		case svc.Interrogate:
			changes <- c.CurrentStatus
			//time.Sleep(100 * time.Millisecond)
		case svc.Stop, svc.Shutdown:
			changes <- svc.Status{State: svc.StopPending}
			if err := ws.Stop(); err != nil {
				return true, 2
			}

			break loop
			// case svc.Pause:
			// 	changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
			// case svc.Continue:
			// 	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
		default:
			continue loop
		}
	}

	return false, 0 //正常启动
}
