/*
	Write a program that will count the maximum number of simultaneous views of a video stream.
	Video stream sessions are saved in a large array of structures {start time, end time}.
*/
package main

import "sort"

type (
	viewSession struct {
		start int64
		end   int64
	}
	viewEvent struct {
		isIncrease bool
		timePoint  int64
	}
	viewSessions []viewSession
	viewEvents   []viewEvent

	viewSessionsSortStart []viewSession
	viewSessionsSortEnd   []viewSession
)

func (e viewEvents) Len() int {
	return len(e)
}

func (e viewEvents) Less(i, j int) bool {
	if e[i].timePoint == e[j].timePoint {
		return e[j].isIncrease
	}
	return e[i].timePoint < e[j].timePoint
}

func (e viewEvents) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// normalizeViewActions will prepare the data for processing.
// separates the start and end of the session and classifies them as different events with the opposite meaning.
// sorts the collected data by time
//  before:
//    {{start: 1, end: 5},{start: 1, end: 8},{start: 5, end: 8}}
//  after:
//    {
//      {isIncrease: true,  timePoint: 1},
//      {isIncrease: true,  timePoint: 1},
//      {isIncrease: false, timePoint: 5},
//      {isIncrease: true,  timePoint: 5},
//      {isIncrease: false, timePoint: 8},
//      {isIncrease: false, timePoint: 8},
//    }
//
//  O(n*log(n))
func normalizeViewActions(sessions viewSessions) viewEvents {
	var events = make(viewEvents, 0, len(sessions)*2)
	for _, session := range sessions {
		events = append(
			events,
			viewEvent{
				isIncrease: true,
				timePoint:  session.start,
			},
			viewEvent{
				isIncrease: false,
				timePoint:  session.end,
			},
		)
	}
	sort.Sort(events) // O(n*log(n))
	return events
}

// calculateMaxViewers1 takes an array of normalized events as an argument.
// traverses the timeline by counting the number of sessions
//  O(n)
func calculateMaxViewers1(events viewEvents) int {
	var (
		maxViewers  = 0
		currViewers = 0
	)
	for _, event := range events {
		if event.isIncrease {
			currViewers++
		} else {
			currViewers--
		}
		if maxViewers < currViewers {
			maxViewers = currViewers
		}
	}
	return maxViewers
}

func (e viewSessionsSortStart) Len() int {
	return len(e)
}

func (e viewSessionsSortStart) Less(i, j int) bool {
	return e[i].start < e[j].start
}

func (e viewSessionsSortStart) Swap(i, j int) {
	e[i].start, e[j].start = e[j].start, e[i].start
}

func (e viewSessionsSortEnd) Len() int {
	return len(e)
}

func (e viewSessionsSortEnd) Less(i, j int) bool {
	return e[i].end < e[j].end
}

func (e viewSessionsSortEnd) Swap(i, j int) {
	e[i].end, e[j].end = e[j].end, e[i].end
}

// calculateMaxViewers2 counts the number of sessions without copying memory,
// but causes the data in the original array to become inconsistent
//
// the logic is to sort the session start and session end events into two separate arrays and
// moving in parallel count the number of sessions based on each next event that comes earlier (start or end)
//  O(n*log(n))
func calculateMaxViewers2(sessions viewSessions) int {
	sort.Sort(viewSessionsSortStart(sessions))
	sort.Sort(viewSessionsSortEnd(sessions))
	var (
		sPos, ePos              = 0, 0
		currViewers, maxViewers = 0, 0
	)
	for {
		if !(ePos < len(sessions) && sPos < len(sessions)) {
			break
		}
		if sessions[ePos].end <= sessions[sPos].start {
			currViewers--
			ePos++
		} else if sessions[sPos].start < sessions[ePos].end {
			currViewers++
			sPos++
		}
		if maxViewers < currViewers {
			maxViewers = currViewers
		}
	}
	return maxViewers
}

func main() {
	println("1. calculateMaxViewers1(normalizeViewActions(session))")
	println("2. calculateMaxViewers2(session)")
}
