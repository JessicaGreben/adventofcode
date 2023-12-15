package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func solution(file string) (int64, error) {
	g, err := ForEachLine(file)
	if err != nil {
		return -1, err
	}

	//res := sequential(g)
	res := concurrent(g)
	//fmt.Println(res)
	return res, nil
}

func concurrent(g *garden) int64 {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	var minLocation int64 = math.MaxInt64
	sem := make(chan int, 1000)

	start := time.Now()
	for i, seedInterval := range g.seedIntervals {
		for i := seedInterval.start; i <= seedInterval.end; i++ {
			sem <- 1
			wg.Add(1)
			go func(i int64) {
				defer wg.Done()
				loc, err := g.getLocation(i, soilCategory)
				if err != nil {
					fmt.Printf("err=%v, seed=%d", err, i)
					return
				}
				mu.Lock()
				if minLocation == math.MaxInt64 || loc < minLocation {
					minLocation = loc
				}
				mu.Unlock()
				<-sem
			}(i)
		}
		fmt.Printf("time=%v, %d/%d seedInterval=%v, minLocation=%d\n", time.Since(start), i+1, len(g.seedIntervals), seedInterval, minLocation)
	}
	wg.Wait()
	return minLocation
}

func sequential(g *garden) int64 {
	var minLocation int64 = math.MaxInt64
	for _, seedInterval := range g.seedIntervals {
		fmt.Printf(">>> %d - seed interval: %#v\n", seedInterval.start, seedInterval)
		for i := seedInterval.start; i <= seedInterval.end; i++ {
			loc, err := g.getLocation(i, soilCategory)
			if err != nil {
				return -1
			}
			fmt.Printf("seed=%d, location=%d\n", i, loc)
			if minLocation == math.MaxInt64 || loc < minLocation {
				minLocation = loc
			}
		}
	}
	return minLocation
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
					return nil, err
				}
				seeds = append(seeds, int64(sInt64))
			}
			g.seeds = seeds
			g.seedIntervals = make([]interval, 0, len(seeds)/2)
			// var totalSeeds int64
			for i := 0; i < len(seeds)-1; i += 2 {
				s, size := seeds[i], seeds[i+1]
				iv := interval{start: s, end: s + size - 1}
				g.seedIntervals = append(g.seedIntervals, iv)
				// totalSeeds += size
			}
			// fmt.Println("totalSeed=", totalSeeds)
			g.locationIntervals = make([]interval, 0, len(g.seedIntervals))
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
		return mapping{}, nil
	}
	destStart, sourceStart, size := x[0], x[1], x[2]
	d, err := strconv.Atoi(destStart)
	if err != nil {
		return mapping{}, err
	}
	s, err := strconv.Atoi(sourceStart)
	if err != nil {
		return mapping{}, err
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		return mapping{}, err
	}
	return mapping{
		source: interval{start: int64(s), end: int64(s + sz - 1)},
		dest:   interval{start: int64(d), end: int64(d + sz - 1)},
	}, nil
}

const (
	seedCategory        = "seeds:"
	soilCategory        = "-soil"
	fertilizerCategory  = "-fertilizer"
	waterCategory       = "-water"
	lightCategory       = "-light"
	temperatureCategory = "-temperature"
	humidityCategory    = "-humidity"
	locationCategory    = "-location"
)

type interval struct {
	start, end int64
}

type mapping struct {
	source interval
	dest   interval
}

func (m *mapping) findMappedNum(num int64) (int64, error) {
	if num < m.source.start || num > m.source.end {
		return -1, nil
	}
	idx := num - m.source.start
	return int64(m.dest.start + idx), nil
}

type garden struct {
	seeds             []int64
	seedIntervals     []interval
	locationIntervals []interval
	mappings          map[string][]mapping
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
