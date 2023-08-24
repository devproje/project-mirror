package util

import "os"

func CreateDir(dirname string) error {
	if _, err := os.Stat(dirname); err != nil {
		err = os.Mkdir(dirname, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
