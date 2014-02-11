package main

import (
  "fmt"
  "io/ioutil"
  "io"
  "crypto/sha1"
  "encoding/hex"
  "bufio"
  "bytes"
  "os"
)

type File struct {
  Name string
  Path string
  HashValue string
}

func (file File) Happy() string {
  return "Happy"
}

func checkFile(path string, fileName string) (fileTest File) {
  fileTest.Name = fileName
  fileTest.Path = path

  var buffer bytes.Buffer
  buffer.WriteString(path)
  buffer.WriteString("/")
  buffer.WriteString(fileName)
  h := sha1.New()

  file, err := os.Open(buffer.String())
  if err != nil {
  }
  defer file.Close()
  reader := bufio.NewReader(file)
  buf := make([]byte, 8096)
  for {
    n, err := reader.Read(buf)
    if err != nil && err != io.EOF {
    }
    if n == 0 {
      break
    }
    h.Write(buf[:n])
  }
  fileTest.HashValue = hex.EncodeToString(h.Sum(nil))
  fmt.Println(fileTest.HashValue)
  fmt.Println(fileTest)
  fmt.Println(fileTest.Happy())
  return fileTest
}

func main() {
  fileList := make([]File,0) 
  dirList, err := ioutil.ReadDir("./save")
  if err != nil {
  }
  
  for _, value := range dirList {
    if value.IsDir() {
    } else {
      fileList = append(fileList,checkFile("./save",value.Name()))
    }
  }
 fmt.Println(fileList)
}
