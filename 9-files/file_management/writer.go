package fileManagement

import "os"

func WriteToFile(filepath string, data string) error {
	return os.WriteFile(filepath, []byte(data), 0644)
}
