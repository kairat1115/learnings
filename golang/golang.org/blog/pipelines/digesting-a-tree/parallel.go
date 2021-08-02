package main

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// For each regular file, start a goroutine that sums the file and sends
	// the result on c.  Send the result of the walk on errc.
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()
			// abort the walk if done is closed
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		// walk has returned, so all calls to wg.add are done.
		// start a goroutine to close c once all the sends are done.
		go func() {
			wg.Wait()
			close(c)
		}()

		// no select needed here, since errc is buffered
		errc <- err
	}()
	return c, errc
}
