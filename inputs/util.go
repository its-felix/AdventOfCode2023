package inputs

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
)

//go:embed *.txt
var inputs embed.FS

func GetInput(day int) fs.File {
	f, err := inputs.Open(fmt.Sprintf("day%d.txt", day))
	if err != nil {
		panic(err)
	}

	return f
}

func GetInputLines(day int) <-chan string {
	ch := make(chan string)
	go func() {
		f := GetInput(day)
		defer f.Close()
		defer close(ch)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}
