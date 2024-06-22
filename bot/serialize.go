package bot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"GoSnipeFun/config"
)

func parseFileToMemory() (records map[string]Record, err error) {
	records = make(map[string]Record)
	file, err := os.Open(config.FileName)
	if err != nil {
		return records, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		address := parts[0]
		record := Record{Address: address}
		if len(parts) > 1 {
			record.Timestamp, _ = strconv.ParseInt(parts[1], 10, 64)
			if record.Timestamp == 0 {
				record.Timestamp = time.Now().Unix()
			}
		}
		if len(parts) >= 3 {
			record.MarketCap, _ = strconv.ParseFloat(parts[2], 64)
		}
		records[address] = record
	}

	if err = scanner.Err(); err != nil {
		return records, fmt.Errorf("error reading file: %v", err)
	}

	return records, nil
}

func saveRecordsToFile(records map[string]Record) error {
	if err := os.Remove(config.FileName); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove file: %v", err)
	}

	file, err := os.OpenFile(config.FileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	for _, record := range records {
		line := fmt.Sprintf("%s %d %.2f\n", record.Address, record.Timestamp, record.MarketCap)
		if _, err := file.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	return nil
}

func addLineToFile(record Record) error {
	file, err := os.OpenFile(config.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	line := fmt.Sprintf("%s %d %.2f\n", record.Address, record.Timestamp, record.MarketCap)
	if _, err := file.WriteString(line); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}
