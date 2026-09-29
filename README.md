# nonce-watch-go

Watches EVM addresses and prints a line whenever one of them sends a transaction. An account's nonce
goes up by one for every transaction it sends, so polling the nonce is enough, without logs or an
indexer.

```bash
go run . 0xAddr1 0xAddr2
go run . -interval 6s -rpc https://mainnet.base.org 0xAddr
```

Each line shows how many transactions the address sent since the previous poll.

```bash
go test ./...
```
