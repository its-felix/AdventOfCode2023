package day5

import (
	"bufio"
	"github.com/its-felix/AdventOfCode2023/util"
	"slices"
	"strings"
	"sync"
)

type almanac struct {
	initialSeeds []int
	seeds        *ranges[[2]int]
	soil         *ranges[[2]int]
	fertilizer   *ranges[[2]int]
	water        *ranges[[2]int]
	light        *ranges[[2]int]
	temperature  *ranges[[2]int]
	humidity     *ranges[[2]int]
	location     *ranges[[2]int]
}

func SolvePart1(input string) int {
	a := parse(input)
	return solve(a, 1, func(i int, s []int) [2]int {
		seed := s[i]
		return [2]int{seed, seed}
	})
}

func SolvePart2(input string) int {
	a := parse(input)
	return solve(a, 2, func(i int, s []int) [2]int {
		start, l := s[i], s[i+1]
		return [2]int{start, start + l}
	})
}

func solve(a almanac, step int, seedFn func(i int, s []int) [2]int) int {
	r := make([][2]int, 0)
	for i := 0; i < len(a.initialSeeds); i += step {
		r = append(r, seedFn(i, a.initialSeeds))
	}

	r = chunkRanges(dedupRanges(r), 5000)

	var wg sync.WaitGroup
	ch := make(chan int, 5000)
	for _, v := range r {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			for seed := start; seed <= end; seed++ {
				ch <- findLowestLocation(a, seed)
			}
		}(v[0], v[1])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	low := <-ch
	for v := range ch {
		if v < low {
			low = v
			println(low)
		}
	}

	return low
}

func chunkRanges(r [][2]int, csize int) [][2]int {
	chunked := make([][2]int, 0)

	for _, v := range r {
		start, end := v[0], v[1]
		for start <= end {
			chunked = append(chunked, [2]int{start, min(start+csize, end)})
			start += csize
		}
	}

	return chunked
}

func dedupRanges(r [][2]int) [][2]int {
	deduped := make([][2]int, 0)
	seen := &ranges[bool]{}

	for _, v := range r {
		if start, end, ok := seen.subrange(v[0], v[1], true); ok {
			deduped = append(r, [2]int{start, end})
		}
	}

	return deduped
}

func findLowestLocation(a almanac, seed int) int {
	chain := []*ranges[[2]int]{
		a.seeds,
		a.soil,
		a.fertilizer,
		a.water,
		a.light,
		a.temperature,
		a.humidity,
		a.location,
	}

	ids := []int{seed}

	for _, r := range chain {
		idsNext := make([]int, 0)

		for _, id := range ids {
			values := r.find(id)
			if len(values) > 0 {
				for _, v := range values {
					srcStart, dstStart := v[0], v[1]
					offset := id - srcStart
					idsNext = append(idsNext, dstStart+offset)
				}
			} else {
				idsNext = append(idsNext, id)
			}
		}

		ids = idsNext
	}

	return slices.Min(ids)
}

func parse(input string) almanac {
	r := almanac{
		seeds:       &ranges[[2]int]{},
		soil:        &ranges[[2]int]{},
		fertilizer:  &ranges[[2]int]{},
		water:       &ranges[[2]int]{},
		light:       &ranges[[2]int]{},
		temperature: &ranges[[2]int]{},
		humidity:    &ranges[[2]int]{},
		location:    &ranges[[2]int]{},
	}

	input, r.initialSeeds = util.ReadInts(util.TrimNonDigit(input))

	for mappingName, mapping := range readMappings(bufio.NewScanner(strings.NewReader(input))) {
		switch mappingName {
		case "seed-to-soil":
			updateMapping(r.seeds, mapping)
		case "soil-to-fertilizer":
			updateMapping(r.soil, mapping)
		case "fertilizer-to-water":
			updateMapping(r.fertilizer, mapping)
		case "water-to-light":
			updateMapping(r.water, mapping)
		case "light-to-temperature":
			updateMapping(r.light, mapping)
		case "temperature-to-humidity":
			updateMapping(r.temperature, mapping)
		case "humidity-to-location":
			updateMapping(r.humidity, mapping)
		}
	}

	return r
}

func updateMapping(src *ranges[[2]int], mapping [][3]int) {
	for _, v := range mapping {
		srcStart := v[1]
		dstStart := v[0]
		l := v[2]

		src.add(srcStart, srcStart+l, [2]int{srcStart, dstStart})
	}
}

func readMappings(s *bufio.Scanner) map[string][][3]int {
	mappings := make(map[string][][3]int)
	currMappingName := ""
	currMapping := make([][3]int, 0)

	for s.Scan() {
		line := s.Text()
		if strings.HasSuffix(line, " map:") {
			if currMappingName != "" {
				mappings[currMappingName] = currMapping
			}

			currMappingName = strings.TrimSuffix(line, " map:")
			currMapping = make([][3]int, 0)
		} else {
			_, values := util.ReadInts(line)
			if len(values) != 3 {
				continue
			}

			currMapping = append(currMapping, [3]int{values[0], values[1], values[2]})
		}
	}

	if currMappingName != "" {
		mappings[currMappingName] = currMapping
	}

	return mappings
}
