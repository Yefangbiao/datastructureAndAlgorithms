package e8_9

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type dire struct {
	name string
	size int64
}

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan dire)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	files := make(map[string]int64)
	bytes := make(map[string]int64)

loop:
	for {
		select {
		case file, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}

			files[file.name]++
			bytes[file.name] += file.size
		case <-tick:
			go func() {
				for fileName := range files {
					fmt.Print(fileName + ":")
					printDiskUsage(files[fileName], bytes[fileName])
				}
			}()

		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(rootName, dir string, n *sync.WaitGroup, fileSizes chan<- dire) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(rootName, subdir, n, fileSizes)
		} else {
			fileSizes <- dire{rootName, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
