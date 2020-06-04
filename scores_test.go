package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_TopNHighestScores(t *testing.T) {
	m := map[int]string{
		10622876: "3c867674494e4a7aac9247a9d9a2179c",
		11027069: "f812d487de244023a6a713e496a8427d",
		11269569: "7ec85fe3aa3c4dd599e23111e7abf5c1",
		11446512: "84a0ccfec7d1475b8bfcae1945aea8f0",
		13214012: "085a11e1b82b441184f4a193a3c9a13c",
	}
	recs := []Record{
		Record{
			Score: 13214012,
			ID:    "085a11e1b82b441184f4a193a3c9a13c",
		},
		Record{
			Score: 11446512,
			ID:    "84a0ccfec7d1475b8bfcae1945aea8f0",
		},
		Record{
			Score: 11269569,
			ID:    "7ec85fe3aa3c4dd599e23111e7abf5c1",
		},
	}
	er, err := json.MarshalIndent(recs, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	expectedResponse := string(er)
	r, err := FindTopNHighestScores(m, 3)
	response := string(r)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResponse, response)

}

func Test_TopNHighestScoresNGreaterThanInputSize(t *testing.T) {
	m := map[int]string{
		10622876: "3c867674494e4a7aac9247a9d9a2179c",
		11027069: "f812d487de244023a6a713e496a8427d",
		11269569: "7ec85fe3aa3c4dd599e23111e7abf5c1",
		11446512: "84a0ccfec7d1475b8bfcae1945aea8f0",
		13214012: "085a11e1b82b441184f4a193a3c9a13c",
	}
	recs := []Record{
		Record{
			Score: 13214012,
			ID:    "085a11e1b82b441184f4a193a3c9a13c",
		},
		Record{
			Score: 11446512,
			ID:    "84a0ccfec7d1475b8bfcae1945aea8f0",
		},
		Record{
			Score: 11269569,
			ID:    "7ec85fe3aa3c4dd599e23111e7abf5c1",
		},
		Record{
			Score: 11027069,
			ID:    "f812d487de244023a6a713e496a8427d",
		},
		Record{
			Score: 10622876,
			ID:    "3c867674494e4a7aac9247a9d9a2179c",
		},
	}
	er, err := json.MarshalIndent(recs, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	expectedResponse := string(er)
	r, err := FindTopNHighestScores(m, 10)
	response := string(r)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResponse, response)
}

func Test_TopNHighestScoresInvalidJSON(t *testing.T) {
	m := map[int]string{
		10622876: "3c867674494e4a7aac9247a9d9a2179c",
		11025835: "error",
		11027069: "f812d487de244023a6a713e496a8427d",
		11269569: "7ec85fe3aa3c4dd599e23111e7abf5c1",
		11446512: "84a0ccfec7d1475b8bfcae1945aea8f0",
		13214012: "085a11e1b82b441184f4a193a3c9a13c",
	}

	// expectedResponse := invalid JSON format. JSON decode error
	_, err := FindTopNHighestScores(m, 10)
	assert.NotEqual(t, nil, err)
}
