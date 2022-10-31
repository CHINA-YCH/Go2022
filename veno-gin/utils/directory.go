package utils

import "os"

/*
 * @Author: ych
 * @Description: ...
 * @File: directory
 * @Version: ...
 * @Date: 2022-10-31 15:39:04
 */

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
