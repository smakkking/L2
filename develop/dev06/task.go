import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"flag"
)

func main() {
	f := flag.Int("fields", 'f', 0, "выбрать поля (колонки)")
	d := flag.String("delimiter", 'd', "\t", "использовать другой разделитель")
	s := flag.Bool("separated", 's', "использовать другой разделитель")

	flag.Parse()
	if *f <= 0 {
		log.Fatal("f <= 0")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		splitTxt := strings.Split(txt, *d)
		if *s && !strings.Contains(txt, *d) {
			fmt.Println("")
		} else if len(splitTxt) < *f {
			fmt.Println(txt)
		} else {
			fmt.Println(splitTxt[*f-1])
		}
	}
}