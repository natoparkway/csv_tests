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
    // csvReader := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(f))
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func makeColumnar(data [][]string) [][]string {
    var columnarData [][]string
    if len(data) > 0 {
        // Add the base columns
        numColumns := len(data[0])
        for i := 0; i < numColumns; i++ {
            columnarData = append(columnarData, []string {})
        }

        for _, row := range data {
            for i, elem := range row {
                columnarData[i] = append(columnarData[i], elem)
            }
        }    
    }
    
    return columnarData
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
    // ./harness.sh "go run go_csv.go"
    records := readCsvFile("data/importer_contacts10M.csv")

    columnar := true
    if columnar {
        columnarRecords := makeColumnar(records)
        fmt.Println("Num Columns:", len(columnarRecords))
        fmt.Println("Num Rows:", len(columnarRecords[0]))
    } else {
        fmt.Println("Num Columns:", len(records[0]))
        fmt.Println("Num Rows:", len(records))
    }


}