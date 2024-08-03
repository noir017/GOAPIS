package tools

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type RunRecord struct {
	Name    string
	RunTime time.Time
}

// 写入记录到 JSON 文件
func WriteRecord(record RunRecord, filePath string) error {
	records, err := ReadRecords(filePath)
	if err != nil {
		return err
	}

	records = append(records, record)

	data, err := json.Marshal(records)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 读取记录从 JSON 文件
func ReadRecords(filePath string) ([]RunRecord, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []RunRecord{}, nil
		}
		return nil, err
	}

	var records []RunRecord
	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// 检查函数是否已经运行过
func IsFunctionRun(name string, filePath string) bool {
	records, err := ReadRecords(filePath)
	if err != nil {
		return false
	}

	today := GetTimeToday()
	for _, record := range records {
		if record.Name == name && record.RunTime.Format("2006-01-02") == today {
			return true
		}
	}

	return false
}
