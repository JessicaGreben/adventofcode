package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	lines, err := inputToLines()
	if err != nil {
		fmt.Println(err)
	}
	fs, err := createFsFromCommands(lines)
	if err != nil {
		fmt.Println(err)
	}
	return sumDirSizes(fs)
}

func sumDirSizes(fs *filesystem) int {
	dirSizes := map[string]int{}
	fs.sumDirSizes(fs.root, dirSizes)
	sum := 0
	for _, v := range dirSizes {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func inputToLines() ([][]string, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	// TODO: convert output to a list of shell commands objects
	output := [][]string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		output = append(output, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

const (
	beginCommand = "$"
	cdCommand    = "cd"
	lsCommand    = "ls"
	parentDir    = ".."
	rootDirName  = "/"
	dirStr       = "dir"
)

func createFsFromCommands(lines [][]string) (*filesystem, error) {
	fs := newFilesystem()
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		firstArg := line[0]

		switch firstArg {
		case beginCommand:
			command := line[1]
			switch command {
			case cdCommand:
				dirArg := line[2]
				switch dirArg {
				case parentDir:
					if fs.currentDir.parent == nil {
						continue
					}
					fs.currentDir = fs.currentDir.parent
				case rootDirName:
					fs.currentDir = fs.root
				default:
					d, ok := fs.currentDir.dirs[dirArg]
					if !ok {
						fmt.Printf("err current dir (parent): %v, child dir name: %v\n", fs.currentDir.name, dirArg)
						continue
					}
					fs.currentDir = d
				}
			case lsCommand:
				continue
			default:
				fmt.Println("not a valid command:", command)
			}
		case dirStr:
			// if we get here then its part of the output after the lsCommand
			dirName := line[1]
			d := newDir(dirName)
			fs.currentDir.addDir(d)
		default:
			// if we get here then its part of the output after the lsCommand
			filesize, err := strconv.Atoi(firstArg)
			if err != nil {
				fmt.Println("err Atoi", firstArg)
			}
			filename := line[1]
			f := newFile(filename, filesize)
			fs.currentDir.addFile(f)
		}
	}
	return fs, nil
}
