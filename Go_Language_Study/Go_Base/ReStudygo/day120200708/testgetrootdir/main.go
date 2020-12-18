package main
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

)
func GetChaosErisRootDir() (chaoseris_root string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	chaoseris_root = substr(dir, 0, strings.LastIndex(dir, "\\"))
	return
}
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
func main(){
fmt.Println(GetChaosErisRootDir())
fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
}