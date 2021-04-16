package main

import (
	"encoding/csv"
	"fmt"
	"github.com/hsqlu/comet-backup/traffic"
	flag "github.com/spf13/pflag"
	"os"
	"strconv"
	"strings"
)

var (
	north = flag.Int32("n", 0, "The North CPM of a Intersection, default 0")
	east  = flag.Int32("e", 0, "The East CPM of a Intersection, default 0")
	south = flag.Int32("s", 0, "The South CPM of a Intersection, default 0")
	west  = flag.Int32("w", 0, "The West CPM of a Intersection, default 0")

	input = flag.String("input", "", "The input file ")

	help = flag.BoolP("help", "h", false, "Display available options")
)

func loadFromFile(path string) ([]traffic.Intersection, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	var intersections []traffic.Intersection
	for _, line := range lines {
		cpms := parseLine(line)
		intersections = append(intersections, traffic.New(cpms[0], cpms[1], cpms[2], cpms[3]))
	}
	return intersections, nil
}

func parseLine(line []string) [4]int32 {
	var ans [4]int32

	for i, s := range line {
		s = strings.TrimSpace(s)
		if v, err := strconv.Atoi(s); err == nil {
			ans[i] = int32(v)
		}
	}
	return ans
}

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if len(strings.TrimSpace(*input)) > 0 {
		out, _ := loadFromFile(*input)
		fmt.Println(out)
	} else {
		inter := traffic.New(*north, *east, *south, *west)
		fmt.Println(inter)
	}
}
