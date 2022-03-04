package main

import (
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

func main() {
	directory := flag.String("d", "", "Root directory to search.")
	flag.Parse()
	if *directory == "" {
		flag.Usage()
		os.Exit(1)
	}
	fileList := list.New()
	getFilesInDirBySize(fileList, *directory)

	for elem := fileList.Front(); elem != nil; elem = elem.Next() {
		fmt.Printf("%d ", elem.Value.(FileNode).Info.Size())
		fmt.Printf("%s\n", elem.Value.(FileNode).FullPath)
	}
}

func getFilesInDirBySize(fileList *list.List, path string) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Error reading directory: " + err.Error())
	}

	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			getFilesInDirBySize(fileList, filepath.Join(path, dirFile.Name()))
		} else if dirFile.Mode().IsRegular() {
			insertSorted(fileList, FileNode{
				FullPath: fullpath,
				Info:     dirFile,
			})
		}
	}
}

func insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		fileList.PushFront(fileNode)
		return
	}
	for elem := fileList.Front(); elem != nil; elem = elem.Next() {
		if fileNode.Info.Size() < elem.Value.(FileNode).Info.Size() {
			fileList.InsertBefore(fileNode, elem)
			return
		}
	}
	fileList.PushBack(fileNode)
}
