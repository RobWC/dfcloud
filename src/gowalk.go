package main

import (
    "path/filepath"
    "os"
    "dffiles"
    "fmt"
    "encoding/json"
)

var fileList []dffiles.File

func main(){
  filepath.Walk("./save", func(path string, info os.FileInfo, err error) error {
    var testFile dffiles.File
    if (!info.IsDir()) {
      testFile.Name = info.Name()
      testFile.Path = path
      testFile.HashValue = testFile.CheckFile()
      fileList = append(fileList,testFile)
    }
    return nil
  })
  fmt.Printf("Total File Count: %d\n",len(fileList))
  b, _ := json.Marshal(fileList)
  os.Stdout.Write(b)
}
