package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/marcuswestin/go-errs"
)

func Cmd(name string, args ...string) (string, errs.Err) {
	if strings.Contains(name, " ") {
		panic("")
	}
	cmd := exec.Command(name, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		stdErrMsg := stderr.String()
		return "", errs.Wrap(err, errs.Info{"Cmd": name, "Args": args}, stdErrMsg)
	}
	return strings.TrimSpace(string(stdout.Bytes())), nil
}

func CmdPrintOutput(name string, args ...string) errs.Err {
	return CmdPrintOutputEnv(nil, name, args...)
}
func CmdPrintOutputEnv(env []string, name string, args ...string) errs.Err {
	if strings.Contains(name, " ") {
		panic("")
	}
	fmt.Println(append([]string{"Cmd:", name}, args...))
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	for _, val := range env {
		cmd.Env = append(cmd.Env, val)
	}
	stdErr := cmd.Run()
	if stdErr != nil {
		return errs.Wrap(stdErr, errs.Info{"Cmd": name, "Args": args})
	}
	return nil
}
