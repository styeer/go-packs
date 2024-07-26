package command

import (
	"errors"
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

// func main() {
// 	name, _ := CpuName()
// 	fmt.Println(name)
// }
