package main

import (
	"encoding/json"
	"log"
	"sort"
)

type Counter struct {
	mp map[string]int
	ww []WC
}

type WC struct {
	word  string
	count int
}

func (c Counter) Count(word string) {
	if _, ok := c.mp[word]; !ok {
		wc := WC{word: word, count: 1}
		c.ww = append(c.ww, wc)
		log.Print(c.ww)
		c.mp[word] = len(c.ww)-1
		return
	}
	c.ww[c.mp[word]].count++
}

func (c Counter) Top() []byte {
	log.Print(c)
	sort.Slice(c.ww, func(i, j int) bool {
		return c.ww[i].count < c.ww[j].count
	})
	jb, _ := json.Marshal(c.ww)
	return jb
}
