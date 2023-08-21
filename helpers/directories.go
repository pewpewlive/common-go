package helpers

import (
	"fmt"
	"os"
)

func EnsureDirectoryExistsAndIsEmpty(directoryPath string) error {

	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if err == nil {
		// Directory exists, remove it
		err := os.RemoveAll(directoryPath)
		if err != nil {
			fmt.Println("Failed to remove directory:", err)
			return err
		}
		fmt.Println("Directory removed successfully.")
	} else if os.IsNotExist(err) {
		// Directory doesn't exist, do nothing
	} else {
		// An error occurred while checking the directory
		fmt.Println("Error:", err)
		return err
	}

	// Create directory
	err = os.MkdirAll(directoryPath, os.ModePerm)
	return err
}
