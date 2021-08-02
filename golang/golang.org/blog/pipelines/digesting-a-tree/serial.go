package main

import (
	"crypto/md5"
	"sync"
)

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5All(root string) (map[string][md5.Size]byte, error) {
	// First implementation
	//
	// m := make(map[string][md5.Size]byte)
	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !info.Mode().IsRegular() {
	// 		return nil
	// 	}
	// 	data, err := ioutil.ReadFile(path)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	m[path] = md5.Sum(data)
	// 	return nil
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return m, nil
	// End first implementation

	// Third implementation
	//
	// MD5All closes the done channel when it returns; it may do so before
	// receiving all the values from c and errc.
	done := make(chan struct{})
	defer close(done)
	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result)
	var wg sync.WaitGroup
	// set a limit
	const numDigesters = 20

	paths, errc := walkFiles(done, root)

	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	// End third implementation

	// Second implementation
	//
	// MD5All closes the done channel when it returns; it may do so before
	// receiving all the values from c and errc.
	// done := make(chan struct{})
	// defer close(done)

	// c, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	// Check whether the Walk failed.
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}
