//readlinebyline 逐行读文件
package readlinebyline

import (
	"bufio"
	"os"
)

func readLineByLine(path string) {
	file, err := os.Open(path)
	if err != nil {
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	for scanner.Scan(){
		txtLines = append(txtLines,scanner.Text())
	}

}
