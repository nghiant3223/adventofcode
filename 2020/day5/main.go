package day1

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

const rowLength = 7

func mySeatID(boardingPasses []string) int {
	//lookupTable := make(map[int]int)
	//for _, pass := range boardingPasses {
	//	id := binaryBoarding(pass)
	//	direction, ok := lookupTable[id]
	//	log.Println(id, direction, ok)
	//	if ok {
	//		if direction == -2 {
	//			return id + 1
	//		}
	//		if direction == 2 {
	//			return id - 1
	//		}
	//		panic("invalid direction")
	//	}
	//	lookupTable[id+2] = 2
	//	lookupTable[id-2] = -2
	//}
	//return -1
	var ids []int
	for _, pass := range boardingPasses {
		id := binaryBoarding(pass)
		fmt.Println(pass, id)
		ids = append(ids, id)
	}
	sort.Ints(ids)
	fmt.Println(ids)
	return 0
}

func maxSeadID(boardingPasses []string) int {
	maxSeatID := math.MinInt64
	for _, pass := range boardingPasses {
		seadID := binaryBoarding(pass)
		if seadID > maxSeatID {
			maxSeatID = seadID
		}
	}
	return maxSeatID
}

func binaryBoarding(boardingPass string) int {
	col := boardingPass[rowLength:]
	row := boardingPass[:rowLength]
	rowBin := rowToBin(row)
	colBin := colToBin(col)
	return binToDec(rowBin)*8 + binToDec(colBin)
}

func rowToBin(row string) string {
	bin := ""
	for i := range row {
		if row[i] == 'B' {
			bin += "1"
		} else if row[i] == 'F' {
			bin += "0"
		} else {
			log.Panicf("invalid character: %s\n", string(row[i]))
		}
	}
	return bin
}

func colToBin(row string) string {
	bin := ""
	for i := range row {
		if row[i] == 'R' {
			bin += "1"
		} else if row[i] == 'L' {
			bin += "0"
		} else {
			log.Panicf("invalid character: %s\n", string(row[i]))
		}
	}
	return bin
}

func binToDec(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(dec)
}
