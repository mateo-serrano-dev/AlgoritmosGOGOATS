package tp2

import (
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
	}
}
