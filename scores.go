package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	// Define command line flags
	// The value of the flags will be stored in corresponding variables at runtime
	filepath = flag.String("fpath", "input.txt", "input file path")
	topn     = flag.Int("n", 3, "top n scores")
	persist  = flag.Bool("persist", false, "persist result to output file")
	memusage = flag.Bool("memusage", false, "memory allocation stats")

	// ErrOpeningInputFile error is occurs when processInputFile is unable to
	// open the input file
	ErrOpeningInputFile = errors.New("failed opening input file")

	// ErrFileIsADirectory error occurs when the filepath (-fpath) is a directory
	// instead of a file
	ErrFileIsADirectory = errors.New("file is a directory and cannot be opened")

	// ErrScoreNotInteger error occurs when the score is not a valid 32-bit
	// integer in the input JSON
	ErrScoreNotInteger = errors.New("score is not a valid integer")

	// ErrIDMissing error occurs when ID field is missing in the JSON document
	ErrIDMissing = errors.New("id is missing. JSON format error")

	// ErrInvalidJSONFormat error occurs when JSON document is missing for
	// a score
	ErrInvalidJSONFormat = errors.New("invalid JSON format. JSON decode error")
)

// Record struct contains a product ID and it's corresponding score
// Struct fields must be exported (capitalized) for JSON encoding
type Record struct {
	Score int    `json:"score"`
	ID    string `json:"id"`
}

// Init function is executed before main as part of the program initialization
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	start := time.Now()
	flag.Parse()

	if *memusage {
		printMemUsage()
	}

	sm, err := processInputFile(*filepath)
	if err != nil {
		// Fatal writes a log message and then terminates the program using
		// the os.Exit(1) system call
		log.Fatal(err)
	}

	jsonResult, err := FindHighScores(sm, *topn)
	if err != nil {
		log.Fatal(err)
	}

	// Write the result to console and (optionally) persist result
	log.Printf("Top N Scores: %s\n", jsonResult)
	if *persist {
		err := writeOutputFile(jsonResult)
		if err != nil {
			log.Fatal(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s \n", elapsed)

	if *memusage {
		printMemUsage()
	}
}

func processInputFile(f string) (map[int]string, error) {
	file, err := os.Open(f)
	fileInfo, _ := os.Stat(f)

	if err != nil {
		return nil, ErrOpeningInputFile
	}
	if err == nil && fileInfo.IsDir() {
		return nil, ErrFileIsADirectory
	}

	// Defer a call to file.Close() so that the file is closed before the
	// main() function exits.
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Parse the input data file
	scanner := bufio.NewScanner(file)
	scoreMap := make(map[int]string)
	var scoreRecord string

	// Set Buffer maximum token size to 1 MB to handle large JSON documents
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	// Scan() method reads the next line of the file and the string that is read
	// is available through the Text() method
	for scanner.Scan() {
		scoreRecord = scanner.Text()
		// Split the record to <score> : <json> format
		keyValue := strings.SplitN(scoreRecord, ":", 2)
		key, err := strconv.Atoi(keyValue[0])
		// exit if score is not an integer
		if err != nil {
			return nil, ErrScoreNotInteger
		}

		// Unmarshall the JSON into a map variable and fetch the `id` field
		value := keyValue[1]
		var c map[string]interface{}
		err = json.Unmarshal([]byte(value), &c)
		if err != nil {
			scoreMap[key] = "error"
			continue
		}
		// log and exit if id is missing in the JSON
		if c["id"] == nil {
			return nil, ErrIDMissing
		}
		// As the value of the map is an interface cast `id` to specific type
		id := c["id"].(string)
		// Store <score> : <id> in the hashmap if JSON decoding is successful
		// Store <score> : <"error"> in hashmap if the JSON format is invalid
		scoreMap[key] = id
	}
	// Err() returns the first non-EOF error that was encountered by the Scanner
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// For debugging
	// fmt.Println(scoreMap)
	return scoreMap, nil
}

// FindHighScores returns n high scores and their corresponding ids from
// the hashmap of <score> : <id>
func FindHighScores(sm map[int]string, n int) ([]byte, error) {
	// As entries in map are unordered, maintain a separate data structure
	// to maintain sorted order of keys
	var keys []int
	for k := range sm {
		keys = append(keys, k)
	}
	// Sort slice of keys in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	// For debugging
	// for _, k := range keys {
	// 	fmt.Println("Key: ", k, "Value: ", sm[k])
	// }

	inputSize := len(keys)
	var result []Record
	for i := 0; i < n; i++ {
		if i >= inputSize {
			break
		}
		k := keys[i]
		v := sm[keys[i]]
		// Check if JSON decoding was successful
		if v == "error" {
			return nil, ErrInvalidJSONFormat
		}
		r := Record{
			Score: k,
			ID:    v,
		}
		result = append(result, r)
	}

	// Marshal result into JSON
	jsonResult, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return nil, err
	}
	return jsonResult, nil
}

func writeOutputFile(jsonResult []byte) error {
	// Create a file to persist the response
	rfile, err := os.Create("output.txt")
	if err != nil {
		return err
	}

	defer func() {
		if err = rfile.Close(); err != nil {
			log.Println(err)
		}
	}()

	bytes, err := rfile.Write(jsonResult)
	if err != nil {
		return err
	}
	log.Printf("Wrote %d bytes to output file \n", bytes)
	return nil
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// HeapAlloc is bytes of allocated heap objects
	fmt.Printf("Heap Allocation = %v bytes", m.HeapAlloc)
	// TotalAlloc is cumulative bytes allocated for heap objects. TotalAlloc increases
	// as heap objects are allocated, but unlike HeapAlloc, it does not decrease when
	// objects are freed
	fmt.Printf("\t Total Allocation = %v bytes \n", m.TotalAlloc)
	// Sys is the space obtained from the OS (reserved by the Go runtime) for heap,
	// stacks and other internal data structures
	fmt.Printf("System space = %v bytes \n", m.Sys)
	// Mallocs is the cumulative count of heap objects allocated
	fmt.Printf("Mallocs = %v", m.Mallocs)
	// Frees is the cumulative count of heap objects freed
	fmt.Printf("\t Frees = %v \n", m.Frees)
}
