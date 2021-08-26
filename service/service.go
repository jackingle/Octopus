package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func NewService() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "Hello from a HandleFunc #1!\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/stocks/", GetStock)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

///https://cloud.iexapis.com/stable/tops?token=YOUR_TOKEN_HERE&symbols=aapl

func GetStock(w http.ResponseWriter, r *http.Request) {
	s := r.URL.RequestURI()
	fmt.Println(s)
	stock := strings.SplitAfter(s, "/")[2]
	url := fmt.Sprintf("https://cloud.iexapis.com/stable/tops?token=%v&symbols=%v", os.Getenv("IEX_TOKEN"), stock)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
