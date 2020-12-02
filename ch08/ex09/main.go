package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type fileSize struct {
	idx         int
	rootDirName string
	nbytes      int64
}

type result struct {
	idx            int
	rootDirName    string
	nfiles, nbytes int64
}

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan *fileSize)

	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(i, root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	results := make([]*result, 0, len(roots))
	for idx, root := range roots {
		results = append(results, &result{idx, root, 0, 0})
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			results[size.idx].nfiles++
			results[size.idx].nbytes += size.nbytes
		case <-tick:
			printDiskUsage(results)
			fmt.Println()
		}
	}

	fmt.Println("Results:")
	printDiskUsage(results) // final totals
}

func printDiskUsage(results []*result) {
	for _, res := range results {
		fmt.Printf("[%s] %d files  %.1f GB\n", res.rootDirName, res.nfiles, float64(res.nbytes)/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(idx int, rootDirName, dir string, n *sync.WaitGroup, fileSizes chan<- *fileSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(idx, rootDirName, subdir, n, fileSizes)
		} else {
			fileSizes <- &fileSize{idx, rootDirName, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
