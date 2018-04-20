package cmdsh

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var ErrTimeout error = errTimeOut{}

type errTimeOut struct{}

func (errTimeOut) Error() string { return "command timed out" }

func ExecPrint(cmdstr string, timeout int) ([]byte, error) {
	log.Println(cmdstr)
	out, err := Exec(cmdstr, timeout)
	if err != nil {
		if err == ErrTimeout {
			log.Println(err)
		} else {
			log.Printf("command errored: %v\n", err)
			log.Println(string(out))
		}
	}
	return out, err
}

func Exec(cmdstr string, timeout int) ([]byte, error) {
	if timeout <= 0 {
		return execCmd(cmdstr)
	}
	return execCmd(fmt.Sprintf("timeout %v %s", timeout, cmdstr))
}

func execCmd(cmdstr string) ([]byte, error) {
	args := strings.Fields(cmdstr)
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	return out, err
}
