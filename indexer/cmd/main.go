package main

import (
	"fmt"
	ia "indexer/internal/app/init"
	"os"
	"runtime/pprof"
)

func main() {
	profilerFile, err := os.Create("z_indexer.pprof")
	if err != nil {
		fmt.Println("There was an error creating the profiler", err)
		return
	}
	defer profilerFile.Close()
	err = pprof.StartCPUProfile(profilerFile)
	if err != nil {
		return
	}

	defer pprof.StopCPUProfile()

	ia.Init()
}
