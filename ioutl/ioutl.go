package ioutl

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/britojr/kbn/utl/errchk"
)

// OpenFile returns a pointer to an open file
func OpenFile(name string) *os.File {
	f, err := os.Open(name)
	errchk.Check(err, fmt.Sprintf("Can't open file %v", name))
	return f
}

// CreateFile returns a pointer to an new file
func CreateFile(name string) *os.File {
	f, err := os.Create(name)
	errchk.Check(err, fmt.Sprintf("Can't create file %v", name))
	return f
}

// Sprintc returns the default formats in a comma-separated string
func Sprintc(a ...interface{}) string {
	s := fmt.Sprintln(a...)
	s = strings.Trim(s, "\n")
	s = strings.Replace(s, " ", ",", -1)
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	return s
}

// ReadYaml reads a yaml file and store it in a string map
func ReadYaml(name string) map[string]string {
	m := make(map[string]string)
	data, err := ioutil.ReadFile(name)
	errchk.Check(err, "")
	errchk.Check(yaml.Unmarshal([]byte(data), &m), "")
	return m
}
