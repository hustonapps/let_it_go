package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err();

	file.Close();

	if err != nil {
		return  nil, errors.New("failed to read line in file")
	}

	return lines, nil
}