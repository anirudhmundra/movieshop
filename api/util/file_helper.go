package util

import "io/ioutil"

type FileHelper interface{
	ReadJsonFile(path string) ([]byte, error)
}

type fileHelper struct{
}

func NewFileHelper() FileHelper{
	return fileHelper{}
}

func (helper fileHelper) ReadJsonFile(path string) ([]byte, error){
	return ioutil.ReadFile(path)
}
