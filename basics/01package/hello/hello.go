package main
import (
  "fmt"
  "os"
)

func main() {
  // go语言不支持main参数和main返回值
  // 获取参数
  if (len(os.Args) > 1) {
    fmt.Println("The first arg:", os.Args[1])
  }
  os.Exit(0)
}
