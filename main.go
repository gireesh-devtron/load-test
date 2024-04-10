package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	ch := make(chan struct{})
	fn := func(n int) {
		for i := 0; i < n; i++ {
			s := ""
			m := &sync.Mutex{}
			for j := 0; j < 100; j++ {
				go func() {
					defer m.Unlock()
					m.Lock()
					s = s + fmt.Sprintf("hello world-%d-%d ", i, j)
					fmt.Println(s)
				}()
			}
		}
	}

	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		ns := r.URL.Query().Get("n")
		n, _ := strconv.Atoi(ns)
		fn(n)
		fmt.Fprint(w, "success")
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		log.Fatal(err)
	}()
	<-ch
}
