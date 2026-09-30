package main

// import stuff
import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var gifURLs = []string{
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/AttackingTheCatBuritto.gif",
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/Bearodynamic.gif",
	"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/BigMackTennisMatch.gif",
}

// when I receive a path for naked path, return it or return an error

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		count := 1
		if requestedCount, err := strconv.Atoi(r.URL.Query().Get("count")); err == nil && requestedCount >= 1 && requestedCount <= 3 {
			count = requestedCount
		}

		var page strings.Builder
		page.WriteString(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Random GIF activity</title>
</head>
<body style="background-color: lightblue;">
  <h1>Hello, world!</h1>
		<form method="get" action="/">
  <label for="count">Number of GIFs:</label>
  <select id="count" name="count">
    <option value="1">1</option>
    <option value="2">2</option>
    <option value="3">3</option>
  </select>
  <button type="submit">Show GIFs</button>
</form>`)
		for range count {
			gifURL := gifURLs[rand.Intn(len(gifURLs))]
			fmt.Fprintf(&page, `
<img src="%s" alt="A cute animal">`, gifURL)
		}
		page.WriteString(`
</body>
</html>`)
		if _, err := w.Write([]byte(page.String())); err != nil {
			log.Printf("write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
