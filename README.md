# High Scores
​
[![Build Status](https://travis-ci.org/jayapriya90/Outline.svg?branch=master)](https://travis-ci.org/jayapriya90/Outline) [![Go Report Card](https://goreportcard.com/badge/github.com/jayapriya90/Outline)](https://goreportcard.com/report/github.com/jayapriya90/Outline)


## Usage



## Sample Input/Output
1. Default input parameters
```
$ ./highestnscores
Key:  13214012 Value:  085a11e1b82b441184f4a193a3c9a13c
Key:  11446512 Value:  84a0ccfec7d1475b8bfcae1945aea8f0
Key:  11269569 Value:  7ec85fe3aa3c4dd599e23111e7abf5c1
Key:  11027069 Value:  f812d487de244023a6a713e496a8427d
Key:  10622876 Value:  3c867674494e4a7aac9247a9d9a2179c
2020/05/19 11:06:57 /work/Outline/highestnscores.go:57: Top N Highest Scores: [
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
2020/05/19 11:06:57 /work/Outline/highestnscores.go:61: Time taken: 1.672525ms
```

2. Absolute file path in `-fpath` flag (Input file outside of current directory)
```
$ ./highestnscores -fpath=/work/input.txt -n=1
Key:  13214012 Value:  085a11e1b82b441184f4a193a3c9a13c
Key:  11446512 Value:  84a0ccfec7d1475b8bfcae1945aea8f0
Key:  11269569 Value:  7ec85fe3aa3c4dd599e23111e7abf5c1
Key:  11027069 Value:  f812d487de244023a6a713e496a8427d
Key:  11025835 Value:  error
Key:  10622876 Value:  3c867674494e4a7aac9247a9d9a2179c
2020/05/19 11:09:08 /work/Outline/highestnscores.go:57: Top N Highest Scores: [
 {
  "score": 13214012,
  "id": "085a11e1b82b441184f4a193a3c9a13c"
 }
]
2020/05/19 11:09:08 /work/Outline/highestnscores.go:61: Time taken: 1.35815ms
```

3. `n` greater than number of records in the input
```
$ ./highestnscores -fpath=input.txt -n=10
Key:  13214012 Value:  085a11e1b82b441184f4a193a3c9a13c
Key:  11446512 Value:  84a0ccfec7d1475b8bfcae1945aea8f0
Key:  11269569 Value:  7ec85fe3aa3c4dd599e23111e7abf5c1
Key:  11027069 Value:  f812d487de244023a6a713e496a8427d
Key:  10622876 Value:  3c867674494e4a7aac9247a9d9a2179c
2020/05/19 11:11:45 /work/Outline/highestnscores.go:57: Top N Highest Scores: [
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
2020/05/19 11:11:45 /work/Outline/highestnscores.go:61: Time taken: 1.44767ms
```

4. Input with invalid JSON in the result
```
$ ./highestnscores -fpath=input_with_error.txt -n=5
Key:  13214012 Value:  085a11e1b82b441184f4a193a3c9a13c
Key:  11446512 Value:  84a0ccfec7d1475b8bfcae1945aea8f0
Key:  11269569 Value:  7ec85fe3aa3c4dd599e23111e7abf5c1
Key:  11027069 Value:  f812d487de244023a6a713e496a8427d
Key:  11025835 Value:  error
Key:  10622876 Value:  3c867674494e4a7aac9247a9d9a2179c
2020/05/19 11:13:08 /work/Outline/highestnscores.go:166: invalid JSON format. JSON decode error
exit status 1
```

5. Run the command with memory allocation stats
```
$ ./highestnscores -fpath=input.txt -n=10 -memusage
Heap Allocation = 86424 bytes  Total Allocation = 86424 bytes
System space = 71125000 bytes
Mallocs = 190  Frees = 4
2020/05/21 00:50:52 /work/Outline/highestnscores.go:56: Top N Highest Scores: [
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
Time taken: 891.308µs
Heap Allocation = 177240 bytes   Total Allocation = 177240 bytes
System space = 71125000 bytes
Mallocs = 583  Frees = 62
```
