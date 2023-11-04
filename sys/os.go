package sys

import (
	"bytes"
	"os"
	"os/exec"
)

type ExecuteConfig struct {
	Exec string
	Args []string
	Dir string
}

func (ec *ExecuteConfig) Run() (stdout, stderr string, err error) {
	cmd := exec.Command(ec.Exec)
	cmd.Dir = ec.Dir
	var stdout_, stderr_ bytes.Buffer
    cmd.Stdout = &stdout_  // 标准输出
    cmd.Stderr = &stderr_  // 标准错误
	err = cmd.Run()

	stdout = stdout_.String()
	stderr = stderr_.String()
	return
}

func Execute(bin string, args ...string) (string, error) {
	cmd := exec.Command(bin, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
