package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	EventRegex      = `\[(?P<timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2})] (?P<message>.*)`
	GuardNumRegex   = `#(\d+)`
	TimestampFormat = `2006-01-02 15:04`
	StartShift      = iota
	WakeUp          = iota
	FallAsleep      = iota
)

type Event struct {
	Timestamp time.Time
	Type      int
	GuardId   uint
}

func main() {
	input, err := shared.LoadInputFile("day4/input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	// Parse input
	events := parseInput(lines)

	// Sort events
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Sub(events[j].Timestamp) < 0
	})

	part1(events)
	fmt.Println("---")
	part2(events)
}

func part1(events []*Event) {
	// Get heatmap
	heatmap := createSleepingHeatmap(events)

	maxSleepingMins := 0
	var maxSleepingGuard uint

	// Loop over every guard
	for guardId, timeline := range heatmap {

		// Calculate total of sleeping minutes
		total := 0
		for _, count := range timeline {
			total += int(count)
		}

		// Keep track of guard with most sleeping minutes
		if total > maxSleepingMins {
			maxSleepingMins = total
			maxSleepingGuard = guardId
		}
	}

	// Get timeline of guard with max sleeping minutes
	timeline := heatmap[maxSleepingGuard]

	currentMaxMin := 0
	currentMax := 0

	// Get minute with most sleep
	for min, count := range timeline {
		if int(count) > currentMax {
			currentMax = int(count)
			currentMaxMin = min
		}
	}

	fmt.Printf("Guard with most sleep: %d\n", maxSleepingGuard)
	fmt.Printf("Minute with most amount of sleep of this guard: %d\n", currentMaxMin)
	fmt.Printf("Checksum: %d\n", int(maxSleepingGuard)*currentMaxMin)
}

func part2(events []*Event) {
	// Get heatmap
	heatmap := createSleepingHeatmap(events)

	var currentMax uint
	var currentMaxGuard uint
	currentMaxMinute := 0

	// Loop over every guard
	for guardId, timeline := range heatmap {
		// Get minute and amount of sleeping
		for min, count := range timeline {

			// If minute has more sleeping
			if count > currentMax {
				// Note current minute and guard
				currentMax = count
				currentMaxGuard = guardId
				currentMaxMinute = min
			}
		}
	}

	fmt.Printf("Guard with minute with most sleep: %d\n", currentMaxGuard)
	fmt.Printf("Minute with most sleep: %d\n", currentMaxMinute)
	fmt.Printf("Checksum: %d\n", currentMaxMinute*int(currentMaxGuard))
}

func createSleepingHeatmap(events []*Event) map[uint][]uint {
	// Create timelines map
	heatmap := make(map[uint][]uint)

	// Guard which is currently guarding
	var currGuard uint

	// Which minute did guard last fall asleep
	startedSleeping := 0

	for _, event := range events {
		switch event.Type {
		case StartShift:
			{
				// Keep track of guard
				currGuard = event.GuardId
			}
		case FallAsleep:
			{
				// Keep track of minute guard falls asleep
				startedSleeping = event.Timestamp.Minute()
			}
		case WakeUp:
			{
				// Get timeline of guard
				timeline, ok := heatmap[currGuard]
				if !ok {
					// If the guard has no timeline yet, make one
					newTimeline := make([]uint, 60)
					heatmap[currGuard] = newTimeline
					timeline = newTimeline
				}

				// Get current minute
				currMinute := event.Timestamp.Minute()

				// Add 1 to every minute since guard started sleeping
				for i := startedSleeping; i < currMinute; i++ {
					timeline[i]++
				}
			}
		}
	}

	return heatmap
}

func parseInput(lines []string) []*Event {
	eventMatcher := regexp.MustCompile(EventRegex)
	guardNumMatcher := regexp.MustCompile(GuardNumRegex)

	events := make([]*Event, len(lines))

	for i, line := range lines {

		// Skip empty lines
		if line == "" {
			continue
		}

		match := shared.RegexMatch(eventMatcher, line)

		timestamp, err := time.Parse(TimestampFormat, match["timestamp"])

		if err != nil {
			panic("Could not parse timestamp")
		}

		message := match["message"]

		eventType := 0

		var guardId uint

		switch message {
		case "falls asleep":
			{
				eventType = FallAsleep
			}
		case "wakes up":
			{
				eventType = WakeUp
			}
		default:
			{
				eventType = StartShift
				num, err := strconv.ParseInt(guardNumMatcher.FindStringSubmatch(message)[1], 10, 64)

				if err != nil {
					panic("Could not convert guard id")
				}

				guardId = uint(num)
			}
		}
		events[i] = &Event{
			Timestamp: timestamp,
			Type:      eventType,
			GuardId:   guardId,
		}
	}

	return events
}
