package dffiles

import (
    "bufio"
    "encoding/hex"
    "bytes"
    "crypto/sha1"
    "os"
    "io"
)

type File struct {
  Name string
  Path string
  HashValue string
}


func (f File) CheckFile() string {
    var buffer bytes.Buffer
    buffer.WriteString(f.Path)
    h := sha1.New()

    file, err := os.Open(buffer.String())
    if err != nil {
    }

    defer file.Close()
    reader :=bufio.NewReader(file)
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
    f.HashValue = hex.EncodeToString(h.Sum(nil))
    return f.HashValue 
}
