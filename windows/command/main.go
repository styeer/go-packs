package command

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CpuName() (string, error) {
	cmd := exec.Command("wmic", "cpu", "get", "name")
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	s := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	if len(s) < 2 {
		return "", errors.New("get cpu infor failed")
	}
	return s[1], nil
}

func ExecServer() (string, error) {
	path, _ := os.Executable()
	cmd1 := exec.Command("wmic", "service", "get", "name,pathname")
	cmd2 := exec.Command("findstr", "/I", "/C:"+path)
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	var out bytes.Buffer
	cmd2.Stdout = &out
	cmd2.Stderr = os.Stderr
	cmd2.Start()
	cmd1.Run()
	cmd2.Wait()
	service := strings.Fields(strings.TrimSpace(out.String()))
	if len(service) == 0 {
		//fmt.Println("Service does not exist")
		log.Fatal("Service not fund")
		return "", errors.New("service not found")
	}
	return service[0], nil
}
