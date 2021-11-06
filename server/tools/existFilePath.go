package tools

import "os"

func IsExistFile(path string) bool {
	_,err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		//if os.IsNotExist(err) {
		//	return false
		//}
		return false
	}
	return true
}