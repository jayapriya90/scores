[![Build Status](https://travis-ci.com/jayapriya90/scores.svg?branch=master)](https://travis-ci.com/jayapriya90/scores)

## Usage

```
go get github.com/jayapriya90/scores
```

```
scores -fpath=<path to input file>
```

## Sample Input/Output
1. Default input parameters
```
$ ./scores
2020/06/03 19:38:24 /work/scores/scores.go:79: Top N Scores: [
 {
  "score": 13214012,
  "id": "085a11e1b82b441184f4a193a3c9a13c"
 },
 {
  "score": 11446512,
  "id": "84a0ccfec7d1475b8bfcae1945aea8f0"
 },
 {
  "score": 11269569,
  "id": "7ec85fe3aa3c4dd599e23111e7abf5c1"
 }
]
Time taken: 888.457µs
```

2. Absolute file path in `-fpath` flag (Input file outside of current directory)
```
$ ./scores -fpath=/work/input.txt -n=1
2020/06/03 19:39:01 /work/scores/scores.go:79: Top N Scores: [
 {
  "score": 13214012,
  "id": "085a11e1b82b441184f4a193a3c9a13c"
 }
]
Time taken: 1.862121ms
```

3. `n` greater than number of records in the input
```
$ ./scores -fpath=input.txt -n=10
2020/06/03 19:39:58 /work/scores/scores.go:79: Top N Scores: [
 {
  "score": 13214012,
  "id": "085a11e1b82b441184f4a193a3c9a13c"
 },
 {
  "score": 11446512,
  "id": "84a0ccfec7d1475b8bfcae1945aea8f0"
 },
 {
  "score": 11269569,
  "id": "7ec85fe3aa3c4dd599e23111e7abf5c1"
 },
 {
  "score": 11027069,
  "id": "f812d487de244023a6a713e496a8427d"
 },
 {
  "score": 10622876,
  "id": "3c867674494e4a7aac9247a9d9a2179c"
 }
]
Time taken: 919.224µs
```

4. Input with invalid JSON in the result
```
$ ./scores -fpath=input_with_error.txt -n=5
2020/06/03 19:41:23 /work/scores/scores.go:75: invalid JSON format. JSON decode error
```

5. Run the command with memory allocation stats
```
$ ./scores -fpath=input.txt -n=10 -memusage -persist
Heap Allocation = 86856 bytes  Total Allocation = 86856 bytes
System space = 71387144 bytes
Mallocs = 196  Frees = 4
2020/06/03 19:43:46 /work/scores/scores.go:79: Top N Scores: [
 {
  "score": 13214012,
  "id": "085a11e1b82b441184f4a193a3c9a13c"
 },
 {
  "score": 11446512,
  "id": "84a0ccfec7d1475b8bfcae1945aea8f0"
 },
 {
  "score": 11269569,
  "id": "7ec85fe3aa3c4dd599e23111e7abf5c1"
 },
 {
  "score": 11027069,
  "id": "f812d487de244023a6a713e496a8427d"
 },
 {
  "score": 10622876,
  "id": "3c867674494e4a7aac9247a9d9a2179c"
 }
]
2020/06/03 19:43:46 /work/scores/scores.go:222: Wrote 357 bytes to output file
Time taken: 2.34051ms
Heap Allocation = 178320 bytes   Total Allocation = 178320 bytes
System space = 71387144 bytes
Mallocs = 599  Frees = 63
```
