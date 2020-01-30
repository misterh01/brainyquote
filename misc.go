package main

import (
	"math/rand"
	"net/http"
)

func randomElem(s []Quote) Quote {
	return s[rand.Intn(len(s))]
}

// queryValue gets url query parameter value
func queryValue(r *http.Request, name string) string {
	q := r.URL.Query()[name]

	queryValue := ""
	if len(q) > 0 {
		queryValue = q[0]
	}

	return queryValue
}

// func readTopicFile(f string) []string {
// 	file, err := os.Open(f)
// 	if err != nil {
// 		log.Fatal("Error with file opening")
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, strings.ToLower(scanner.Text()))
// 	}
// 	return lines
// }