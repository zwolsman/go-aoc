package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2019/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var asteroids []Vector
	for y, row := range strings.Split(string(data), "\n") {
		for x, c := range row {
			if c != '#' {
				continue
			}

			asteroids = append(asteroids, Vector{float64(x), float64(y)})
		}
	}

	laser := part1(asteroids)
	part2(laser, asteroids)
}

func part1(asteroids []Vector) (pos Vector) {
	best := 0
	for i := 0; i < len(asteroids); i++ {
		base := asteroids[i]
		angles := make(map[float64]bool)
		for j := 0; j < len(asteroids); j++ {
			if j == i {
				continue
			}

			asteroid := asteroids[j]
			dist := asteroid.min(base)
			angles[dist.Angle()] = true
		}

		if val := len(angles); val > best {
			pos = base
			best = val
		}
	}
	fmt.Printf("Best position is %v where we can see %d other astroids in direct line of sight.\n", pos, best)
	return
}

func part2(laser Vector, asteroids VectorArray) {

	asteroidMap := make(map[float64]VectorArray)
	var angles []float64
	for _, asteroid := range asteroids {
		if asteroid == laser {
			continue
		}
		dist := laser.min(asteroid)
		angle := dist.Angle()
		asteroidMap[angle] = append(asteroidMap[angle], dist)
	}

	// sort by distance
	for angle, v := range asteroidMap {
		angles = append(angles, angle)
		sort.Sort(v)

		asteroidMap[angle] = v
	}
	sort.Sort(sort.Float64Slice(angles))
	curr := sort.SearchFloat64s(angles, Vector{x: 0, y: 1}.Angle())

	for i := 0; i < 200; i++ {
		var acquireTarget func() (float64, VectorArray)

		acquireTarget = func() (float64, VectorArray) {
			angle := angles[curr%len(angles)]
			fmt.Printf("using angle %v\n", angle)
			targets, ok := asteroidMap[angle]
			curr++
			if !ok {
				log.Fatal("getting asteroid is not ok.")
			}
			if len(targets) == 0 { //no more asteroids left..
				return acquireTarget()
			}
			return angle, targets
		}

		angle, targets := acquireTarget()
		fmt.Printf("Round %d -> Shooting from %v to %v; all targets: %v\n\n", i+1, laser, laser.min(targets[0]), targets)
		asteroidMap[angle] = targets[1:]
	}
}

type Vector struct {
	x, y float64
}

func (v Vector) min(o Vector) Vector {
	return Vector{
		x: v.x - o.x,
		y: v.y - o.y,
	}
}

func (v Vector) Angle() float64 {
	return math.Atan2(v.y, v.x)
}

func (p Vector) Length() int {
	return int(math.Abs(p.x) + math.Abs(p.y))
}

type VectorArray []Vector

func (p VectorArray) Len() int           { return len(p) }
func (p VectorArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p VectorArray) Less(i, j int) bool { return p[i].Length() < p[j].Length() }

func (v Vector) String() string {
	return fmt.Sprintf("Vector(%.f, %.f)", v.x, v.y)
}
