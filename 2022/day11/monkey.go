package day11

type monkey struct {
	inspectionCount int
	items           []int
	op              string
	opInt           int
	test            int
	testTrue        int
	testFalse       int
}

type throw struct {
	item int
	to   int
}
