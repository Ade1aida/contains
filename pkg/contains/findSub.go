package contains

import (
	"bytes"
	"fmt"
	"os"
)

func Contain(ourFile string, text string) (bool, error) {

	file, err := os.ReadFile(ourFile)
	if err != nil {
		return false, fmt.Errorf("FILE DIDN`T EXIST!")
	}

	return bytes.Contains(file, []byte(text)), nil
}
