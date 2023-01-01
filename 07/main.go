package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FS struct {
	currentDir *Dir
	root       *Dir
}

func (f *FS) cd(arg string) {
	if arg == "/" {
		f.currentDir = f.root
		return
	}
	if arg == ".." {
		if f.currentDir.parent == nil {
			panic("cannot cd .. while in root")
		}
		f.currentDir = f.currentDir.parent
		return
	}

	for _, dir := range f.currentDir.dirs {
		if dir.name == arg {
			f.currentDir = dir
			return
		}
	}
	fmt.Println("tried to cd into", arg, "but it doesnt exist in current dir", f.currentDir)
	panic("dir doesnt exist, exiting")

}

func NewFS() *FS {
	root := NewDir("/", nil)
	return &FS{
		currentDir: root,
		root:       root,
	}
}

type Dir struct {
	name      string
	files     []*File
	dirs      []*Dir
	parent    *Dir
	totalSize int
}

func (d *Dir) AddFile(file *File) {
	d.files = append(d.files, file)
}

func (d *Dir) AddDir(dir *Dir) {
	d.dirs = append(d.dirs, dir)
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:   name,
		parent: parent,
	}
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func main() {
	inputFile, err := os.Open("input.txt")
	check(err)
	defer inputFile.Close()

	fs := NewFS()

	populateFsFromInputFile(fs, inputFile)

	// printTree(fs.root, 0)

	sumUpToHundredK := sumDirsUpToHundredK(fs.root, 0)
	fmt.Println("sum of dir sizes up to hundred k is", sumUpToHundredK)

}

func populateFsFromInputFile(fs *FS, input *os.File) {

	currentCommand := ""
	var currentCommandOutput []string

	forLineOfFile(input, func(line string) {
		if string(line[0]) == "$" {
			if currentCommand != "" {
				processCurrentCommand(fs, currentCommand, currentCommandOutput)
			}
			currentCommand = line
			currentCommandOutput = []string{}
			return
		}
		currentCommandOutput = append(currentCommandOutput, line)
	})

}

func processCurrentCommand(fs *FS, command string, output []string) {
	args := strings.Fields(command)
	if args[0] != "$" {
		panic("commands must start with %")
	}
	if args[1] == "cd" {
		if len(args) != 3 {
			panic("cd must have only one argument")
		}
		fs.cd(args[2])
	}
	if args[1] == "ls" {
		for _, item := range output {
			itemFields := strings.Fields(item)
			if len(itemFields) != 2 {
				panic("each ls out must have 2 fields!")
			}
			if itemFields[0] == "dir" {
				fs.currentDir.dirs = append(fs.currentDir.dirs, NewDir(itemFields[1], fs.currentDir))
			} else {
				size, err := strconv.Atoi(itemFields[0])
				check(err)
				fs.currentDir.files = append(fs.currentDir.files, NewFile(itemFields[1], size))
				increaseDirSizes(fs.currentDir, size)
			}
		}
	}
}

func increaseDirSizes(dir *Dir, size int) {

	currentDir := dir
	currentDir.totalSize += size

	for currentDir.parent != nil {
		currentDir = currentDir.parent
		currentDir.totalSize += size
	}

}

func printTree(dir *Dir, level int) {
	currentDir := dir
	fmt.Println(strings.Repeat(" ", level), currentDir.name, currentDir.totalSize)
	for _, child := range currentDir.dirs {
		printTree(child, level+1)
	}
}

func sumDirsUpToHundredK(dir *Dir, sum int) int {
	currentDir := dir
	currentSum := sum

	if currentDir.totalSize <= 100000 {
		currentSum += currentDir.totalSize
	}
	for _, child := range currentDir.dirs {
		currentSum = sumDirsUpToHundredK(child, currentSum)
	}
	return currentSum

}

// general utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func forLineOfFile(file *os.File, fn func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
}
