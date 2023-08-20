package util

import "os"

func CreateDir(dirname string) error {
	if _, err := os.Stat(dirname); err != nil {
		err = os.Chmod(dirname, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
