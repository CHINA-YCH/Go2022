package read_line

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

// DoWriteFile 写文件
func DoWriteFile(filePath string) *os.File {
	//_filePath := "./test.txt"
	_filePath := filePath
	_file, _err := os.OpenFile(_filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if _err != nil {
		log.Errorf("打开文件错误=%v\n", _err)
		return nil
	}
	//提前关闭文件
	//defer func(_file *os.File) {
	//	err := _file.Close()
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}(_file)
	return _file
}

func Do(msg string, _file *os.File) error {
	//写入文件
	_writer := bufio.NewWriter(_file)
	_, err := _writer.WriteString(msg + "\n")
	if err != nil {
		return err
	}
	err = _writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
