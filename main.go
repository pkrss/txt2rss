package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func handlerRss(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		fileName = "1"
	}
	fileName = "html/txt/" + fileName + ".txt"

	var ss []string
	if bContent, err := ioutil.ReadFile(fileName); err != nil {
		w.Write([]byte(err.Error()))
		return
	} else {
		content := string(bContent)
		bContent = nil
		content = strings.Replace(content, "\r", "", -1)
		ss = strings.Split(content, "\n")
		content = ""
	}

	if r.URL.Query().Get("random") == "1" {
		idx := randInt(0, len(ss))
		ss = append(ss[idx:], ss[0:idx]...)
	}

	w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>
	<rss version="2.0">
	<channel>
		<title><![CDATA[ 1 ]]></title>
		<language>zh-cn</language>
		<ttl>60</ttl>`))

	for i := range ss {
		w.Write([]byte("<item><title><![CDATA[ "))
		w.Write([]byte(ss[i]))
		w.Write([]byte(" ]]> </title></item>"))
	}

	w.Write([]byte(`</channel></rss>`))
}

func main() {
	addr := flag.String("addr", "127.0.0.1:28081", "bind addrress, example: 127.0.0.1:28081")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("html/")))
	http.HandleFunc("/rss/", handlerRss)
	log.Printf("start bind: " + *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
