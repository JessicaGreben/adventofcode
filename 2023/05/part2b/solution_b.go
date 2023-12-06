package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution(file string) (int64, error) {
	g, err := parseInput(file)
	if err != nil {
		return -1, err
	}
	var minLocation int64 = math.MaxInt64
	for _, seed := range g.seedIntervals {
		ivs := g.getLocationIntervals(seed, 1)
		for _, iv := range ivs {
			if minLocation == math.MaxInt64 || iv.start < minLocation {
				minLocation = iv.start
			}
		}
	}

	return minLocation, nil
}

func parseInput(file string) (*garden, error) {
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
			g.seedIntervals = make([]interval, 0, len(seeds)/2)
			for i := 0; i < len(seeds)-1; i += 2 {
				s, size := seeds[i], seeds[i+1]
				iv := interval{start: s, end: s + size - 1}
				g.seedIntervals = append(g.seedIntervals, iv)
			}
			continue
		}
		cat := g.getCategory(line)
		if cat != "" {
			currCat = cat
			continue
		}

		m, err := lineTomapping(line)
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

func lineTomapping(line string) (mapping, error) {
	x := strings.Split(line, " ")
	if len(x) < 3 {
		return mapping{}, nil
	}
	dstStart, srcStart, size := x[0], x[1], x[2]
	d, err := strconv.Atoi(dstStart)
	if err != nil {
		return mapping{}, err
	}
	s, err := strconv.Atoi(srcStart)
	if err != nil {
		return mapping{}, err
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		return mapping{}, err
	}
	return mapping{
		src: interval{start: int64(s), end: int64(s + sz - 1)},
		dst: interval{start: int64(d), end: int64(d + sz - 1)},
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
	src interval
	dst interval
}

func (m *mapping) contains(num int64) bool {
	return num >= m.src.start && num <= m.src.end
}
func (m *mapping) convert(num int64) int64 {
	idx := abs(num - m.src.start)
	return int64(m.dst.start + idx)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

type garden struct {
	seedIntervals []interval
	mappings      map[string][]mapping
}

func newGarden() *garden {
	return &garden{
		mappings: map[string][]mapping{},
	}
}

func (g *garden) getLocationIntervals(iv interval, catIdx int) []interval {
	if catIdx >= len(categories) {
		return []interval{iv}
	}
	locationIvs := []interval{}
	maps := g.mappings[categories[catIdx]]
	ivs := convertInterval(iv, maps)
	for _, nextIv := range ivs {
		nextIvs := g.getLocationIntervals(nextIv, catIdx+1)
		locationIvs = append(locationIvs, nextIvs...)
	}
	return locationIvs
}

// if map.src contains both iv.start and iv.end then convert to new interval convert to map.dst
// if map.src not contains iv.start and contains iv.end then split iv into 2 new intervals and call g.convert on both
// if map.src contains iv.start and not contains iv.end then split iv into 2 new intervals and call g.convert on both
// if map.src not contains both iv.start and iv.end then convert to new interval with same values as iv
func convertInterval(iv interval, maps []mapping) []interval {
	result := []interval{}
	for _, m := range maps {
		if m.contains(iv.start) && m.contains(iv.end) {
			newIv := interval{
				start: m.convert(iv.start),
				end:   m.convert(iv.end),
			}
			return []interval{newIv}
		}

		if !m.contains(iv.start) && m.contains(iv.end) {
			newIv := interval{
				start: m.convert(m.src.start),
				end:   m.convert(iv.end),
			}
			remainingIv := interval{
				start: iv.start,
				end:   m.src.start - 1,
			}
			result = append(result, newIv)
			result = append(result, convertInterval(remainingIv, maps)...)
			return result
		}
		if m.contains(iv.start) && !m.contains(iv.end) {
			newIv := interval{
				start: m.convert(iv.start),
				end:   m.convert(m.src.end),
			}

			remainingIv := interval{
				start: m.src.end + 1,
				end:   iv.end,
			}

			result = append(result, newIv)
			result = append(result, convertInterval(remainingIv, maps)...)
			return []interval{newIv}
		}
	}
	if len(result) == 0 {
		return []interval{iv}
	}
	return result
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
