# info2parser

A simple Go library and CLI tool to parse INFO2 `$I` files, from the Windows Recycle Bin. 

It extracts metadata such as the original file path, deletion time, and original size.

[DISCLAIMER] Only works for files from Windows 10 and 11

## Instalation

``` go install github.com/gobelinor/info2parser/info2parser@latest ```

## Usage 

### CLI
``` info2parser -file [file] ```
Output: 
```
==== $I FILE ====
Header          : 1
FileSize        : 102400 octets
DeletionTime    : 2025-04-06T14:31:20Z
FileNameLength  : 88
OriginalPath    : C:\Users\JohnDoe\Desktop\secret.txt
```

### As a Go library
```
import "github.com/gobelinor/info2parser"

func main() {
    data, err := os.ReadFile("path/to/$I123ABC")
    if err != nil {
        log.Fatal(err)
    }

    info, err := info2parser.Parse(data)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Original file path:", info.OriginalPath)
    fmt.Println("Deleted at:", info2parser.FiletimeToTime(info.DeletionTime))
}
```

### Author
Developed by @gobelinor and chatGPT 

