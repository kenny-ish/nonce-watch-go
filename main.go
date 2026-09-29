package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func nonce(url, addr string) (uint64, error) {
	body, _ := json.Marshal(map[string]any{"jsonrpc": "2.0", "id": 1, "method": "eth_getTransactionCount", "params": []any{addr, "latest"}})
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var r struct{ Result string }
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}
	return strconv.ParseUint(strings.TrimPrefix(r.Result, "0x"), 16, 64)
}

func main() {
	url := flag.String("rpc", "https://ethereum-rpc.publicnode.com", "JSON-RPC endpoint")
	interval := flag.Duration("interval", 12*time.Second, "poll interval")
	flag.Parse()
	addrs := flag.Args()
	if len(addrs) == 0 {
		log.Fatal("usage: nonce-watch-go [-rpc URL] ADDRESS...")
	}
	last := map[string]uint64{}
	seen := map[string]bool{}
	for {
		for _, a := range addrs {
			n, err := nonce(*url, a)
			if err != nil {
				log.Printf("%s: %v", a, err)
				continue
			}
			if d, ok := Delta(last[a], n, seen[a]); ok {
				fmt.Printf("%s  %s sent %d tx (nonce %d -> %d)\n", time.Now().Format("15:04:05"), a, d, last[a], n)
			} else if !seen[a] {
				fmt.Printf("watching %s (nonce %d)\n", a, n)
			}
			last[a], seen[a] = n, true
		}
		time.Sleep(*interval)
	}
}
