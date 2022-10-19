package utils

import "os"

/*
定义 utils 工具函数
新建 utils/directory.go 文件，编写 PathExists 函数，用于判断路径是否存在
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
