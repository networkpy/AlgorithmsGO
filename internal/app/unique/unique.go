package unique

import (
	"bufio"
	"fmt"
	"io"
)

// Получает на вход файл, где должны быть отсортированы цифры
func Unique(input io.Reader, output io.Writer) error {
	reader := bufio.NewScanner(input)
	var prev string
	for reader.Scan() {
		text := reader.Text()
		if text == prev {
			continue
		}
		if text < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = text
		fmt.Fprintln(output, text)
	}
	return nil
}
