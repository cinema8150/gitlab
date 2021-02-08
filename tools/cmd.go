package tools

import (
	"bytes"
	"os/exec"
	"strings"
)

func init() {

}

// Exec 执行指令
func Exec(cmd string) (string, error) {
	params := strings.Split(cmd, " ")
	CMD := exec.Command(params[0], params[1:]...)
	var stdout, stderr bytes.Buffer
	CMD.Stdout = &stdout
	CMD.Stderr = &stderr
	err := CMD.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		return errStr, err
	}
	return outStr, nil
}
