package utils

import "os"
import "strings"
import "fmt"
import "sync"
import "errors"

func checkErrorLogger(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateDumpFileDirectory(path string) {
	if ok := CheckFileIsExist(path); !ok {
		os.Mkdir(path, 0777)
	}
}

var _dumpFileSaveLock sync.Mutex = sync.Mutex{}

func DumpFileSave(pathName, fileName string, data []string) (err error) {

	CreateDumpFileDirectory(pathName)

	defer func() {
		_dumpFileSaveLock.Unlock()
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error occur on dump file; unknow error %s", err)
		}
	}()

	_dumpFileSaveLock.Lock()

	_file_name := strings.Split(fileName, ".")
	if len(_file_name) != 2 {
		err = errors.New("Error handle file; Does not comply with the rule file name")
	}

	_tmpfile_name := "tmp_" + _file_name[0]

	if err = WriteFile(pathName, _tmpfile_name, data); err != nil {
		checkErrorLogger(err)
	}

	_tmpfile_name = pathName + string(os.PathSeparator) + _tmpfile_name

	_new_file := pathName + string(os.PathSeparator) + fileName

	os.Rename(_tmpfile_name, _new_file)

	return err
}

func WriteFile(pathName, fileName string, data interface{}) (err error) {

	file := pathName + string(os.PathSeparator) + fileName
	if CheckFileIsExist(file) {
		os.Remove(file)
	}

	var tmp interface{}
	switch data.(type) {
	case string:
		tmp = data.(string)
	case int:
		tmp = Str(data)
	case []byte:
		tmp = Str(data.([]byte)[:])
	case []string:
		tmp = strings.Join(data.([]string), "\n")
	default:
	}

	f, err := os.Create(file)
	defer f.Close()
	checkErrorLogger(err)

	_, err = f.WriteString(tmp.(string))
	checkErrorLogger(err)

	return err
}
