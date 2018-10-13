package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type fsItem struct {
	name string
	path []string
	last bool
	dir bool
	size uint
}

func Tree(out *bytes.Buffer, path string, flag bool) error {
	items, createError := createTree(path, flag)
	if createError != nil {
		fmt.Printf("error while creating dir tree: %v", createError)
		return createError
	}
	treeString, printError := printTree(items, flag)
	if printError != nil {
		fmt.Printf("error while creating dir tree: %v", printError)
		return printError
	}
	out.WriteString(treeString)
	return nil
}

func createTree(path string, flag bool) ([]fsItem, error) {
	items := make([]fsItem, 0, 20)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		item := strings.Split(path, "/")
		if flag || info.IsDir() {
			items = append(items, fsItem{item[len(item)-1], item, false, info.IsDir(),uint(info.Size())})
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	items[0].last = true

	last := map[int]bool{1: true}
	for i := len(items) - 1; i > 0; i-- {
		pathLen := len(items[i].path)
		if !last[pathLen] {
			items[i].last = true
			last[pathLen] = true
		}
		if len(items[i-1].path) > pathLen {
			newLast := map[int]bool{pathLen: true}
			for j := 0; j < pathLen; j++ {
				newLast[j] = last[j]
			}
			last = newLast
		}
	}
	return items, nil
}

func printTree(items []fsItem, flag bool) (string, error) {
	result := fmt.Sprintf("└───%s \n", items[0].name)
	indent := ""
	for i := 1; i < len(items); i++ {
		prevPathLen := len(items[i-1].path)
		if prevPathLen < len(items[i].path) {
			if items[i-1].last {
				indent += "   \t"
			} else {
				indent += "│\t"
			}
		} else if prevPathLen > len(items[i].path) {
			nestedLevel := prevPathLen - len(items[i].path)
			indent = indent[:(len(indent) - (4 * nestedLevel))]
		}

		pref := ""
		if items[i].last {
			pref += "└───"
		} else {
			pref += "├───"
		}
		size := "empty"
		if items[i].size > 0 {
			size = strconv.Itoa(int(items[i].size)) + "b"
		}

		if flag && !items[i].dir {
			result += fmt.Sprintf("%s%s%v  (%s) \n", indent, pref, items[i].name, size)
		} else {
			result += fmt.Sprintf("%s%s%v \n", indent, pref, items[i].name)
		}
	}
	return result, nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	buf := new(bytes.Buffer)
	err := Tree(buf, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintln(out, buf.String())
}
