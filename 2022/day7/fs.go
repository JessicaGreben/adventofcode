package day7

import (
	"fmt"
	"strings"
)

type dir struct {
	name   string
	parent *dir
	// dirs is a map of names to child directories
	dirs  map[string]*dir
	files []*file
}

func newDir(name string) *dir {
	return &dir{
		name:  name,
		dirs:  map[string]*dir{},
		files: []*file{},
	}
}

func (d *dir) addDir(childDir *dir) {
	if _, ok := d.dirs[childDir.name]; ok {
		fmt.Println("dir already exists", childDir.name)
	}
	childDir.parent = d
	d.dirs[childDir.name] = childDir
}

func (d *dir) addFile(file *file) {
	d.files = append(d.files, file)
}

// totalSize sums the size of all files in dir.
func (d *dir) sumFileSizes() int {
	sum := 0
	for _, file := range d.files {
		sum += file.size
	}
	return sum
}

// path returns the path to the dir from the root dir of the filesystem
func (d *dir) path() string {
	path := []string{}
	curr := d
	for curr != nil && curr.name != rootDirName {
		path = append(path, curr.name)
		curr = curr.parent
	}

	reverse(path)
	return "/" + strings.Join(path, "/")
}

func reverse(input []string) {
	i, j := 0, len(input)-1
	for i < j {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}
}

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

type filesystem struct {
	root       *dir
	currentDir *dir
}

func newFilesystem() *filesystem {
	r := newDir("/")
	return &filesystem{
		root:       r,
		currentDir: r,
	}
}

func (fs *filesystem) sumDirSizes(dir *dir, sums map[string]int) {
	sum := dir.sumFileSizes()
	for _, childDir := range dir.dirs {
		child := childDir
		fs.sumDirSizes(child, sums)
		sum += sums[child.path()]
	}
	sums[dir.path()] = sum
}

// tree walks the filesystem and prints the name of each dir and file
func (fs *filesystem) tree() {
	dfsPrint(fs.root, "-")
}

func dfsPrint(dir *dir, prefix string) {
	fmt.Println(prefix, dir.name, "(dir)")
	p := fmt.Sprintf("  %s", prefix)
	for _, f := range dir.files {
		suffix := fmt.Sprintf("(file, size=%d)", f.size)
		fmt.Println(p, f.name, suffix)
	}
	for _, d := range dir.dirs {
		dd := d
		dfsPrint(dd, p)
	}
}
