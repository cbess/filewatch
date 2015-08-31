package watcher

import (
	"fmt"
	"os"
	"time"
)

type FileWatcher struct {
	Path string
	// true if the file contents have been changed
	Changed bool
	// true if the file watching for this file is active
	Active bool
	// active watch channel
	// sends false to stop watching once active
	watchChan chan bool
	// the last known timestamp for a modification
	lastModUnixTime int64
}

func New(filePath string) *FileWatcher {
	fw := &FileWatcher{Path: filePath}
	fw.Changed = false
	fw.Active = false
	fw.watchChan = make(chan bool)
	return fw
}

func (fw *FileWatcher) Watch() error {
	return fw.WatchPath(fw.Path)
}

func (fw *FileWatcher) WatchPath(filePath string) error {
	if fi, err := os.Stat(filePath); err == nil {
		if fi.IsDir() {
			return fmt.Errorf("Should be using a file path, directory found at %q", fi.Name())
		}

		fw.lastModUnixTime = fi.ModTime().Unix()
	}

	fw.Active = true
	go fw.monitorFile()
	return nil
}

func (fw *FileWatcher) StopWatch() {
	fw.watchChan <- false
	fw.Active = false
}

func (fw *FileWatcher) monitorFile() {
	fmt.Println("Monitoring file:", fw.Path)

	for {
		select {
		case watch := <-fw.watchChan:
			if !watch {
				fmt.Println("Stop watch path:", fw.Path)
			}
			return

		default:
			// check file mod date
			if fi, err := os.Stat(fw.Path); err == nil {
				modTime := fi.ModTime().Unix()
				if modTime > fw.lastModUnixTime {
					fw.Changed = true
					fmt.Println("-- Changed:", fw.Path)
					fw.lastModUnixTime = modTime
				}
			} else {
				fmt.Println("fail:", err)
				fw.StopWatch()
			}

			time.Sleep(time.Second)
			break
		}
	}
}
