package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution(file string) (int64, error) {
	g, err := ForEachLine(file)
	if err != nil {
		return -1, err
	}

	var minLocation int64 = math.MaxInt64
	for _, seed := range g.seeds {
		location, err := g.getLocation(seed, soilCategory)
		if err != nil {
			return -1, err
		}
		if minLocation == math.MaxInt64 || location < minLocation {
			minLocation = location
		}
	}
	return minLocation, nil
}

func ForEachLine(file string) (*garden, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)
	g := newGarden()
	var currCat string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, seedCategory) {
			out := strings.Split(line, " ")
			seeds := []int64{}
			for _, s := range out[1:] {
				sInt64, err := strconv.Atoi(s)
				if err != nil {
					fmt.Println("s")
					return nil, err
				}
				seeds = append(seeds, int64(sInt64))
			}
			g.seeds = seeds
			continue
		}
		cat := g.getCategory(line)
		if cat != "" {
			currCat = cat
			continue
		}

		m, err := lineToMapping(line)
		if err != nil {
			return nil, err
		}
		g.mappings[currCat] = append(g.mappings[currCat], m)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return g, nil
}

func lineToMapping(line string) (mapping, error) {
	x := strings.Split(line, " ")
	if len(x) < 3 {
		fmt.Println(x)
		return mapping{}, nil
	}
	destStart, sourceStart, size := x[0], x[1], x[2]
	d, err := strconv.Atoi(destStart)
	if err != nil {
		fmt.Println("4")

		return mapping{}, err
	}
	s, err := strconv.Atoi(sourceStart)
	if err != nil {
		fmt.Println("5")

		return mapping{}, err
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("6	")

		return mapping{}, err
	}
	return mapping{
		// size:   s,
		source: interval{start: int64(s), end: int64(s + sz - 1)},
		dest:   interval{start: int64(d), end: int64(d + sz - 1)},
	}, nil
}

var (
	seedCategory        = "seeds:"
	soilCategory        = "-soil"
	fertilizerCategory  = "-fertilizer"
	waterCategory       = "-water"
	lightCategory       = "-light"
	temperatureCategory = "-temperature"
	humidityCategory    = "-humidity"
	locationCategory    = "-location"
)

var categories = []string{
	seedCategory,
	soilCategory,
	fertilizerCategory,
	waterCategory,
	lightCategory,
	temperatureCategory,
	humidityCategory,
	locationCategory,
}

type interval struct {
	start, end int64
}

type mapping struct {
	//size   int
	source interval
	dest   interval
}

func (m *mapping) findMappedNum(num int64) (int64, error) {
	if num < m.source.start || num > m.source.end {
		return -1, nil
	}
	// fmt.Println("num:", num)
	// fmt.Printf("map=%#v\n", m)
	//destStart + size

	idx := num - m.source.start
	// fmt.Println("idx=", idx)

	return int64(m.dest.start + idx), nil
}

type garden struct {
	seeds    []int64
	mappings map[string][]mapping
}

func newGarden() *garden {
	return &garden{
		mappings: map[string][]mapping{},
	}
}

func (g *garden) getCategory(line string) string {
	var currCategory string
	switch {
	case strings.Contains(line, soilCategory):
		currCategory = soilCategory
	case strings.Contains(line, fertilizerCategory):
		currCategory = fertilizerCategory
	case strings.Contains(line, waterCategory):
		currCategory = waterCategory
	case strings.Contains(line, lightCategory):
		currCategory = lightCategory
	case strings.Contains(line, temperatureCategory):
		currCategory = temperatureCategory
	case strings.Contains(line, humidityCategory):
		currCategory = humidityCategory
	case strings.Contains(line, locationCategory):
		currCategory = locationCategory
	default:
		return ""
	}
	g.mappings[currCategory] = []mapping{}
	return currCategory
}

func (g *garden) getLocation(num int64, category string) (location int64, err error) {
	// fmt.Println("num:", num, ", cat=", category, "mappings:", g.mappings[category])
	var nextNum int64 = 0

	for _, m := range g.mappings[category] {
		nextNum, err = m.findMappedNum(num)
		if err != nil {
			return -1, err
		}
		if nextNum != -1 {
			break
		}
	}
	if nextNum == -1 || nextNum == 0 {
		nextNum = num
	}

	switch cat := category; cat {
	case soilCategory:
		return g.getLocation(nextNum, fertilizerCategory)
	case fertilizerCategory:
		return g.getLocation(nextNum, waterCategory)
	case waterCategory:
		return g.getLocation(nextNum, lightCategory)
	case lightCategory:
		return g.getLocation(nextNum, temperatureCategory)
	case temperatureCategory:
		return g.getLocation(nextNum, humidityCategory)
	case humidityCategory:
		return g.getLocation(nextNum, locationCategory)
	case locationCategory:
		return nextNum, nil
	default:
		return -1, errors.New("not valid category")
	}
}
