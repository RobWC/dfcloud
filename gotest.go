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

func checkFile(path string, fileName string) {
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
  fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

func main() {
  dirList, err := ioutil.ReadDir("./save")
  if err != nil {
  }
  
  for _, value := range dirList {
    if value.IsDir() {
    } else {
      checkFile("./save",value.Name())
    }
  }
}
