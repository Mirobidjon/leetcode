package medium

import (
	"strconv"
)

func NumberOfRounds(startTime string, finishTime string) int {
	hStart, _ := strconv.Atoi(startTime[:2])
	mStart, _ := strconv.Atoi(startTime[3:])
	hFinish, _ := strconv.Atoi(finishTime[:2])
	mFinish, _ := strconv.Atoi(finishTime[3:])

	if mStart%15 > 0 {
		mStart = (mStart / 15) + 1
	} else {
		mStart = (mStart / 15)
	}
	mFinish = (mFinish / 15)

	if hStart == hFinish && mStart > mFinish {
		hFinish += 24
	}

	if hStart > hFinish {
		hFinish += 24
	}

	return (hFinish-hStart)*4 + (mFinish - mStart)
}
