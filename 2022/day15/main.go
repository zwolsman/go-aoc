package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"io"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in, 2_000_000))
	fmt.Println(part2(in))
}

func part1(in []byte, y int) int {

	sensors := make(map[common.Vector]int)
	beacons := make(map[common.Vector]int)

	reader := bytes.NewReader(in)
	maxX, minX := 0, 0
	for {
		var sx, sy, bx, by int
		_, err := fmt.Fscanf(reader, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sx, &sy, &bx, &by)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			// TODO: workaround
			if err == io.ErrUnexpectedEOF {
				break
			}

			panic(err)
		}

		sensor := common.Vector{X: sx, Y: sy}
		beacon := common.Vector{X: bx, Y: by}
		dist := sensor.Dist(beacon)
		sensors[sensor] = dist

		if n := sensor.X + dist; n > maxX {
			maxX = n
		}

		if n := sensor.X - dist; n < minX {
			minX = n
		}

		beacons[beacon]++
	}

	var possibilities int
	for x := minX; x < maxX; x++ {
		current := common.Vector{X: x, Y: y}
		if _, ok := sensors[current]; ok {
			possibilities--
			continue
		}

		if _, ok := beacons[current]; ok {
			continue
		}

		for sensor, maxDist := range sensors {
			if sensor.Dist(current) <= maxDist {
				possibilities--
				break
			}
		}
	}

	return possibilities * -1
}

func part2(in []byte) any {
	return nil
}
