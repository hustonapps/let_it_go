package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJSON(filename string, data interface{}) error {
	file, err := os.Create(filename)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	file.Close()

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil;
}