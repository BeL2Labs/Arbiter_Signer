// Copyright (c) 2025 The bel2 developers

package events

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

func CreateConfirmDir(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("create dir ", filePath)
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			fmt.Println(" createConfirmDir filePaht ", err)
			return err
		}
	}
	return nil
}

func SaveContractEvent(path string, event *ContractLogEvent) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(event)

	dir := filepath.Dir(path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	err = os.WriteFile(path, buffer.Bytes(), 0644)
	return err
}

func UpdateCurrentBlock(datadir string, block uint64) error {
	fielPath := datadir + "/" + "listened_block.txt"
	dir := filepath.Dir(fielPath)
	_, err := os.Stat(fielPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	height := big.NewInt(0).SetUint64(block)
	err = os.WriteFile(fielPath, height.Bytes(), 0644)
	return err
}

func GetCurrentBlock(datadir string) (uint64, error) {
	fielPath := datadir + "/" + "listened_block.txt"

	fileContent, err := os.ReadFile(fielPath)
	if err != nil {
		return 0, err
	}
	block := big.NewInt(0).SetBytes(fileContent)
	return block.Uint64(), nil
}
