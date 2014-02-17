package main

import (
  "fmt"
  "io/ioutil"
  "bytes"
  "dffiles"
  "os"
  "path/filepath"
)

func testFile(file dffiles.File, currentPath string) {

}

func testDir(dirList []os.FileInfo, path string, fileList []dffiles.File) {
     var buffer bytes.Buffer
     buffer.WriteString(path)
     buffer.WriteString("/")
     buffer.WriteString(value.Name())
     fmt.Println(buffer.String())
     currentPath = buffer.String()
     subDir, err := ioutil.ReadDir(buffer.String())
}

/*
Loop through
If file process
If directory recurse until no files
pass current directory hierarchy 

*/


func recurse(dirList []os.FileInfo, basePath string, fileList []dffiles.File) []os.FileInfo {
  /*currentPath := basePath
  for _, value := range dirList {
   if value.IsDir() {
     fmt.Println("isDir")
     fmt.Println(value.Name())
     //recurse
     subDir, err := ioutil.ReadDir(buffer.String())
     for _, subDirItem range subDir {
       if subDirItem.IsDir() {
       } else {
       } 
     }
     if err != nil {
         fmt.Println(err)
     } else if (len(subDir) > 0) {
         dirList = append(dirList,subDir...)
     } else {
       //empty dir
       fmt.Println("empty dir")
     }
   } else {
     var testFile dffiles.File
     testFile.Path = currentPath
     fmt.Println(value.Name())
     testFile.Name = value.Name()
     testFile.CheckFile()
     fmt.Println(testFile.HashValue)
     fileList = append(fileList,testFile)
   }
  }*/
  return dirList
}
func main() {
  basePath := "./save"
  filepath.Walk(basePath)
/*  fileList := make([]dffiles.File,0) 
  dirList, err := ioutil.ReadDir(basePath)
  if err != nil {
  }
  xx := reflect.TypeOf(dirList).Kind()
  fmt.Printf("Type = %s\n",xx)
 
  for (len(dirList) > 0) {
    dirList = recurse(dirList,basePath,fileList)
  }

 fmt.Println(fileList)
 */
}
