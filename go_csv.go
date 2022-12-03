package main

import (
    "encoding/csv"
    "log"
    "os"
    "fmt"
    "unsafe"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func arraySize(arr [][]string) int {
    totalSize := unsafe.Sizeof(arr)
    for _, subArr := range arr {
        totalSize += unsafe.Sizeof(subArr)
        for _, elem := range subArr {
            totalSize += unsafe.Sizeof(elem)
        }
    }
    return int(totalSize)
}

func main() {
    // /usr/bin/time -lp go run go_csv.go
    // ./harnesssh "go run go_csv.go"
    records := readCsvFile("data/importer_contacts10M.csv")
    fmt.Println("Array Size is", float64(arraySize(records)) / (1024 * 1024), "MBs")

}