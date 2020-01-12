package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const levels = "▁▂▃▄▅▆▇█"

var (
	decFlag bool
	incFlag bool
)

func init() {
	flag.BoolVar(&decFlag, "dec", false, "Decrement backlight")
	flag.BoolVar(&incFlag, "inc", false, "Increment backlight")
}

func levelSize(max int) int {
	return max / len([]rune(levels))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type action func(current, levelSize, max int) int

func noop(current, _, _ int) int {
	return current
}

func inc(current, levelSize, max int) int {
	if current < 0 {
		return levelSize
	}
	return minInt(current+levelSize, max)
}

func dec(current, levelSize, max int) int {
	if current > max {
		return max - levelSize
	}
	return maxInt(current-levelSize, levelSize)
}

func set(next int) error {
	brightness := strconv.Itoa(next)
	f, err := os.Create("/sys/class/backlight/intel_backlight/brightness")
	if err != nil {
		return err
	}
	if _, err = f.WriteString(brightness); err != nil {
		_ = f.Close() // Ignore file close error
		return err
	}
	return f.Close()
}

func get(bl string) (int, error) {
	data, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/" + bl)
	if err != nil {
		return -1, err
	}
	value, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return -1, err
	}
	return value, nil
}

func levelIndex(current, levelSize, max int) int {
	if current >= max {
		return len([]rune(levels)) - 1
	}
	if current <= levelSize {
		return 0
	}
	return current/levelSize - 1
}

// LevelGlyph gets the backlight brightness level glyph that represents current brightness
func LevelGlyph(current, levelSize, max int) string {
	if current < levelSize {
		return "‼"
	}
	return string([]rune(levels)[levelIndex(current, levelSize, max)])
}

// PrintStatus prints the current status of the backlight brightness level.
func PrintStatus() {}

func main() {
	flag.Parse()
	var action action = noop
	switch {
	case decFlag && !incFlag:
		action = dec
	case !decFlag && incFlag:
		action = inc
	}
	max, err := get("max_brightness")
	if err != nil {
		log.Fatal(err)
	}
	levelSize := levelSize(max)
	current, err := get("brightness")
	if err != nil {
		log.Fatal(err)
	}
	next := action(current, levelSize, max)
	err = set(next)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s▏%d%%\n", LevelGlyph(next, levelSize, max), int(float64(next)/float64(max)*100.0))
	return
}
