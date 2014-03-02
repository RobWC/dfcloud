package main

import (
    "path/filepath"
    "os"
    "dffiles"
    "fmt"
   // "encoding/json"
    "runtime"
    "sync"
)

var fileList []dffiles.File
var wg sync.WaitGroup

func testFile(f *dffiles.File) {
    defer wg.Done()
    f.HashValue = f.CheckFile()
}

func printFinal() {
    fmt.Printf("Total File Count: %d\n",len(fileList))
//    b, _ := json.Marshal(fileList)
  //  os.Stdout.Write(b)
}

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())
  filepath.Walk("./save", func(path string, info os.FileInfo, err error) error {
    var newFile dffiles.File
    if (!info.IsDir()) {
      newFile.Name = info.Name()
      newFile.Path = path
      fileList = append(fileList,newFile)
    }
    return nil
  })
  for key := range fileList {
      wg.Add(1)
      go testFile(&fileList[key])
  }
  wg.Wait()
  printFinal()
}
