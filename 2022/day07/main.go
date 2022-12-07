package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

const (
	FileSystemDiskSize = 70_000_000
	RequiredDiskSpace  = 30_000_000
)

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	root := buildFileSystem(in)
	root.printNode(0)

	var sum int
	var walk func(f *file)
	walk = func(f *file) {
		for _, c := range f.children {
			if c.size == 0 {
				if c.totalSize() <= 100_000 {
					sum += c.totalSize()
				}
				walk(c)
			}
		}
	}
	walk(root)

	return sum
}

func part2(in []byte) any {
	root := buildFileSystem(in)
	currentlyUsed := FileSystemDiskSize - root.totalSize()
	needed := RequiredDiskSpace - currentlyUsed

	target := math.MaxInt

	var walk func(f *file)
	walk = func(f *file) {
		for _, c := range f.children {
			if c.size == 0 {
				size := c.totalSize()
				if size >= needed && size < target {
					target = size
				}
				walk(c)
			}
		}
	}
	walk(root)

	return target
}

type file struct {
	size     int
	name     string
	parent   *file
	children []*file
}

func buildFileSystem(in []byte) *file {

	var dir *file

	cd := func(param string) {
		switch param {
		case "/":
			dir = &file{
				name: "/",
			}
			break
		case "..":
			dir = dir.parent
			break
		default:
			for _, c := range dir.children {
				if c.name == param {
					dir = c
					break
				}
			}

		}
	}

	output := strings.Split(string(in), "\n")

	readLs := func(i int) (files []*file, read int) {
		for !strings.HasPrefix(output[i+read], "$") {
			split := strings.Split(output[i+read], " ")
			read++

			size, _ := strconv.Atoi(split[0])

			files = append(files, &file{
				size:   size,
				name:   split[1],
				parent: dir,
			})
			if i+read == len(output) {
				break
			}
		}

		return
	}

	for i := 0; i < len(output); i++ {
		line := output[i]

		switch {
		case strings.HasPrefix(line, "$ cd "):
			cd(line[5:])
			break
		case strings.HasPrefix(line, "$ ls"):
			files, read := readLs(i + 1)
			i += read
			dir.children = files
			break
		}
	}

	root := dir.parent
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (f *file) printNode(count int) {
	offset := strings.Repeat(" ", count)
	fmt.Print(offset + f.name + " ")
	if f.size == 0 {
		fmt.Println("(dir)")
	} else {
		fmt.Printf("(file, size=%d)\n", f.size)
	}

	for _, c := range f.children {
		c.printNode(count + 1)
	}
}

func (f *file) totalSize() int {
	sum := f.size
	for _, c := range f.children {
		sum += c.totalSize()
	}
	return sum
}
