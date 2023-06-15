package TTools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func GetAnsysCSPAgentManagerServiceAppPathByAppName(osServiceManagerAppName string) string {
	fmt.Println("GetAnsysCSPAgentManagerServiceAppPathByAppName - start")

	var appDataByAppNamePath string
	switch runtime.GOOS {
	case "linux":
		appDataByAppNamePath = filepath.Join("/usr/local/go", osServiceManagerAppName)
	case "windows":
		appDataByAppNamePath = filepath.Join("C:\\go", osServiceManagerAppName)
	case "darwin":
		appDataByAppNamePath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", osServiceManagerAppName)
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}

	fmt.Println("GetAnsysCSPAgentManagerServiceAppPathByAppName:", appDataByAppNamePath)

	return appDataByAppNamePath
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func ReadPIDFromFile(filepath string) (int, error) {
	pidBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(pidBytes))
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func WritePIDToFile(filepath string, pid int) (bool, error) {
	pidBytes := []byte(strconv.Itoa(pid))

	// This function overwrites the existing file or creates a new file
	err := ioutil.WriteFile(filepath, pidBytes, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
