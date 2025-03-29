package ambulance_wl

import (
	"time"

	"slices"
)

func (a *Ambulance) reconcileWaitingList() {
	slices.SortFunc(a.WaitingList, func(left, right WaitingListEntry) int {
		if left.WaitingSince.Before(right.WaitingSince) {
			return -1
		} else if left.WaitingSince.After(right.WaitingSince) {
			return 1
		} else {
			return 0
		}
	})

	// we assume the first entry EstimatedStart is the correct one (computed before previous entry was deleted)
	// but cannot be before current time
	// for sake of simplicity we ignore concepts of opening hours here

	if a.WaitingList[0].EstimatedStart.Before(a.WaitingList[0].WaitingSince) {
		a.WaitingList[0].EstimatedStart = a.WaitingList[0].WaitingSince
	}

	if a.WaitingList[0].EstimatedStart.Before(time.Now()) {
		a.WaitingList[0].EstimatedStart = time.Now()
	}

	nextEntryStart :=
		a.WaitingList[0].EstimatedStart.
			Add(time.Duration(a.WaitingList[0].EstimatedDurationMinutes) * time.Minute)
	for _, entry := range a.WaitingList[1:] {
		if entry.EstimatedStart.Before(nextEntryStart) {
			entry.EstimatedStart = nextEntryStart
		}
		if entry.EstimatedStart.Before(entry.WaitingSince) {
			entry.EstimatedStart = entry.WaitingSince
		}

		nextEntryStart =
			entry.EstimatedStart.
				Add(time.Duration(entry.EstimatedDurationMinutes) * time.Minute)
	}
}
