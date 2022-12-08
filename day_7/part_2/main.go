package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File interface {
	name() string
	isDir() bool
	size() int
}

type file struct {
	_name string
	_size int
}

type dir struct {
	file
	children []File
}

func (f *file) isDir() bool {
	return false
}

func (d *dir) isDir() bool {
	return true
}

func (f *file) size() int {
	return f._size
}

func (f *file) name() string {
	return f._name
}

func (d dir) findChild(target string) File {
	var child File
	for _, c := range d.children {
		if c.name() == target {
			return c
		}
	}
	return child
}

func walk(d *dir, depth int) {
	fmt.Printf("%*s (%d)\n", depth, d.name(), d.size())
	new_depth := depth + 1
	for _, c := range d.children {
		if c.isDir() {
			walk(c.(*dir), new_depth)
		} else {
			fmt.Printf("%*s (%d)\n", new_depth, c.name(), c.size())
		}
	}
}

func populateSizes(d *dir) {
	my_size := 0
	for _, c := range d.children {
		if c.isDir() {
			populateSizes(c.(*dir))
		}
		my_size += c.size()
	}
	d._size = my_size
}

func sumSmallDirs(d *dir) int {
	sum := 0
	fmt.Printf("Size of %s is %d\n", d.name(), d.size())
	if d.size() < 100000 {
		fmt.Println("It is")
		sum += d.size()
	}
	for _, c := range d.children {
		if c.isDir() {
			sum += sumSmallDirs(c.(*dir))
		}
	}
	return sum
}

func smallestSuitable(d *dir, needed int, smallest *int) {
	fmt.Printf("My size %d\n", d.size())
	if d.size() >= needed && d.size() < *smallest {
		fmt.Printf("Found better one %s\n", d.name())
		*smallest = d.size()
	}
	for _, c := range d.children {
		if c.isDir() {
			smallestSuitable(c.(*dir), needed, smallest)
		}
	}
}

func main() {

	root := dir{file: file{_name: "/"}}

	stack := list.New()
	stack.PushFront(root)

	curr := &root

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ ") {
			line := line[2:]
			if strings.HasPrefix(line, "cd") {
				dir_name := line[3:]
				if dir_name == "/" {
					curr = &root
					stack := list.New()
					stack.PushFront(curr)
				} else if dir_name == ".." {
					parent := stack.Front()
					stack.Remove(parent)
					curr = parent.Value.(*dir)
				} else {
					child := curr.findChild(dir_name)
					stack.PushFront(curr)
					curr = child.(*dir)
				}
			}
		} else {
			// we are listing
			if strings.HasPrefix(line, "dir ") {
				name := line[4:]
				curr.children = append(curr.children, &dir{file: file{_name: name}})
			} else {
				parts := strings.Split(line, " ")
				fsize, _ := strconv.Atoi(parts[0])
				name := parts[1]
				curr.children = append(curr.children, &file{_name: name, _size: fsize})
			}
		}
	}
	populateSizes(&root)
	walk(&root, 0)
	res := root.size()
	smallestSuitable(&root, root.size()-40000000, &res)
	fmt.Printf("Smallest suitable: %d\n", res)
}
