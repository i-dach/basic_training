# 問題
echoプログラムを修正して、そのプログラムを起動したコマンド名であるos.Args[0]も表示するようにしなさい

# echoプログラム

```
func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
```