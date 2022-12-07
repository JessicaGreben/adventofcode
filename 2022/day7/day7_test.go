package day7

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 1555642, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := 5974547, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestCreateFS(t *testing.T) {
	testCases := []struct {
		name       string
		input      [][]string
		currentDir string
		fileCount  int
		dirCount   int
	}{

		{
			"0",
			[][]string{},
			"/",
			0,
			0,
		},
		{
			"1",
			[][]string{{"$", "cd", "/"}},
			"/",
			0,
			0,
		},
		{
			"2",
			[][]string{{"$", "ls"}, {"$", "cd", "/"}},
			"/",
			0,
			0,
		},
		{
			"3",
			[][]string{{"$", "ls"}, {"dir", "a"}, {"$", "cd", "a"}},
			"a",
			0,
			0,
		},
		{
			"4",
			[][]string{{"$", "ls"}, {"dir", "a"}, {"14848514", "b.txt"}},
			"/",
			1,
			1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotfs, err := createFsFromCommands(tc.input)
			if err != nil {
				t.Error(err)
			}

			if want, got := rootDirName, gotfs.root.name; want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
			if want, got := tc.currentDir, gotfs.currentDir.name; want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
			if want, got := tc.fileCount, len(gotfs.currentDir.files); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
			if want, got := tc.dirCount, len(gotfs.currentDir.dirs); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
func TestCreateFSTree(t *testing.T) {
	case1 := [][]string{
		{"$", "cd", "/"},
		{"$", "ls"},
		{"dir", "a"},
		{"14848514", "b.txt"},
		{"8504156", "c.dat"},
		{"dir", "d"},
		{"$", "cd", "a"},
		{"$", "ls"},
		{"dir", "e"},
		{"29116", "f"},
		{"2557", "g"},
		{"62596", "h.lst"},
		{"$", "cd", "e"},
		{"$", "ls"},
		{"584", "i"},
		{"$", "cd", ".."},
		{"$", "cd", ".."},
		{"$", "cd", "d"},
		{"$", "ls"},
		{"4060174", "j"},
		{"8033020", "d.log"},
		{"5626152", "d.ext"},
		{"7214296", "k"},
	}
	testCases := []struct {
		name   string
		input  [][]string
		output int
	}{

		{
			"1",
			case1,
			7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotfs, err := createFsFromCommands(tc.input)
			if err != nil {
				t.Error(err)
			}
			//gotfs.tree()

			if want, got := 95437, sumDirSizes(gotfs); want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
