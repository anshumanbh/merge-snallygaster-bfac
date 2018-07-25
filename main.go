package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/xtgo/set"
)

type config struct {
	snallygasterFile string
	bfacFile         string
	outFile          string
}

type snallygaster struct {
	Cause string `json:"cause"`
	URL   string `json:"url"`
	Misc  string `json:"misc"`
}

type bfac struct {
	URL           string `json:"url"`
	StatusCode    int    `json:"status_code"`
	ContentLength int    `json:"content_length"`
}

var (
	cfg        config
	sgarray    []snallygaster
	bfarray    []bfac
	backupurls []string
)

func loadConfig() {
	snallygasterFile := flag.String("snallygasterFile", "snallygaster.json", "Snallygaster scan file in JSON")
	bfacFile := flag.String("bfacFile", "bfac.json", "Bfac scan file in JSON")
	outFile := flag.String("outFile", "results.txt", "Final output merged file in .txt")

	flag.Parse()

	cfg = config{
		snallygasterFile: *snallygasterFile,
		bfacFile:         *bfacFile,
		outFile:          *outFile,
	}

}

func exists(path string) (bool, int64, error) {
	fi, err := os.Stat(path)
	if err == nil {
		return true, fi.Size(), nil
	}
	if os.IsNotExist(err) {
		return false, int64(0), nil
	}
	return false, int64(0), err
}

func ensureFilePathExists(filepath string) error {
	value := false
	fsize := int64(0)

	for (value == false) || (fsize == int64(0)) {
		i, s, err := exists(filepath)
		if err != nil {
			log.Println("Failed to determine if the file exists or not..")
		}
		value = i
		fsize = s
	}

	log.Println(filepath+" File exists:", value)
	log.Println(filepath+" File size:", fsize)

	return nil
}

func loopsgfile(sgfile string) error {

	sf, err := os.Open(sgfile)
	if err != nil {
		log.Printf("Couldn't open the Snallygaster Scan file: %v", err)
		return err
	}
	defer sf.Close()

	sfbytearray, err := ioutil.ReadAll(sf)
	if err != nil {
		log.Printf("Couldn't read the opened Snallygaster Scan file: %v", err)
		return err
	}
	json.Unmarshal(sfbytearray, &sgarray)

	for _, sgurls := range sgarray {
		backupurls = append(backupurls, sgurls.URL)
	}
	return nil
}

func loopbffile(bffile string) error {

	bf, err := os.Open(bffile)
	if err != nil {
		log.Printf("Couldn't open the Bfac Scan file: %v", err)
		return err
	}
	defer bf.Close()

	bfbytearray, err := ioutil.ReadAll(bf)
	if err != nil {
		log.Printf("Couldn't read the opened Bfac Scan file: %v", err)
		return err
	}
	json.Unmarshal(bfbytearray, &bfarray)

	for _, bfurls := range bfarray {
		backupurls = append(backupurls, bfurls.URL)
	}
	return nil
}

func writeResultsToCsv(results []string, outputFilePath string) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Couldn't create the output file: %v", err)
		return err
	}
	defer outputFile.Close()

	if len(results) != 0 {
		for _, str := range results {
			outputFile.WriteString(str + "\n")
		}
	} else {
		outputFile.WriteString("NA")
	}
	return nil
}

func main() {

	loadConfig()

	err := ensureFilePathExists(cfg.snallygasterFile)
	if err != nil {
		log.Fatalf("Couldn't ensure whether the Snallygaster file exists or not: %v", err)
	}

	err = ensureFilePathExists(cfg.bfacFile)
	if err != nil {
		log.Fatalf("Couldn't ensure whether the Bfac file exists or not: %v", err)
	}

	err = loopsgfile(cfg.snallygasterFile)
	if err != nil {
		log.Fatalf("Couldn't loop the snallygaster scan file: %v", err)
	}

	err = loopbffile(cfg.bfacFile)
	if err != nil {
		log.Fatalf("Couldn't loop the bfac scan file: %v", err)
	}

	data := sort.StringSlice(backupurls)
	sort.Sort(data)
	n := set.Uniq(data) // Uniq returns the size of the set
	data = data[:n]     // trim the duplicate elements

	fmt.Println(data)

	// writing the results to the outfile once everything is done
	err = writeResultsToCsv(data, cfg.outFile)
	if err != nil {
		log.Fatalf("Couldn't write to the out file: %v", err)
	}

	fmt.Println("=======================================")
	fmt.Println("Results saved to: " + cfg.outFile)
}
