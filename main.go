package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getValue(inputPrompt string, defaultValue int32) int32 {
	log.SetFlags(0)
	log.SetPrefix("InvalidInput: ")

	resourceCount := defaultValue
	validInput := false

	// Input parsing loop
	for !validInput {
		fmt.Print(inputPrompt)

		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Reading Error")
			continue
		}

		trimmedInput := strings.TrimSuffix(input, "\n")

		inputValue, err := strconv.ParseInt(trimmedInput, 10, 32)
		if err != nil {
			log.Println("Please enter an integer (32 bit) value!")
		}

		resourceCount = int32(inputValue)
	}

	return resourceCount
}

// How much of a resource should be reserved for what ranges is provided as
// integers (percentage decimal * 10'000).
func GetResourceToReserve(totalResource, percentage int32) int32 {
	return totalResource * percentage / 100 / 100
}

func main() {
	cpuReservationValues := [][2]int32{
		// mCPU value: reservation percentage * 100
		{0, 600},    // 6.00%
		{1000, 100}, // 1.00%
		{2000, 50},  // 0.50%
		{4000, 25},  // 0.25%
	}

	ramReservationValues := [][2]int32{
		// RAM in MiB: reservation percentage * 100
		{0, 2500},     // 25%
		{4096, 2000},  // 20%
		{8192, 1000},  // 10%
		{16384, 600},  //  6%
		{131072, 200}, //  2%
	}

	var cpuReservation int32
	var ramReservation int32

	cpuCount := getValue("Enter the amount of CPUs (Cores) per node [2]: ", 2)
	ramCount := getValue("Enter the amount of RAM (GiB) per node [4]: ", 4)
	hardEviction := getValue("Enter hard eviction threshold (MiB) [100]: ", 100)

	mCpus := cpuCount * 1000
	mibRam := ramCount * 1024

	for rangeIndex, rangeValues := range cpuReservationValues {
		if mCpus > rangeValues[0] {
		} else {
			break
		}
	}

	fmt.Println(
		cpuCount,
		ramCount,
		hardEviction,
		mCpus,
		mibRam,
		cpuReservation,
		ramReservation,
	)
}
