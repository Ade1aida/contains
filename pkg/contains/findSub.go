package contains

import (
	"bytes"
	"fmt"
	"os"
)

func Contain(ourFile string, text string) (bool, error) {

	file, err := os.ReadFile(ourFile) //, os.O_RDONLY, 0666)
	if err != nil {
		return false, fmt.Errorf("FILE DIDN`T EXIST!")
	}

	return bytes.Contains(file, []byte(text)), nil
}
