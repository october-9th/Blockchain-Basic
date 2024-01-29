# Simple implementation of blockchain network in golang

### method

- mine_block: mine new block
- get_chain: view blockchain information
- is_valid: validate chains

### how to run

- first run:

```bash
go run main.go
```

- wait til gin finish initialize and run main.go file
- then open cmd and run following command

1. get_chain

```bash
curl localhost:8080/get_chain
```

2. mine_block

```bash

curl -XPOST localhost:8080/mine_block
```

3. is_valid

```bash
curl localhost:8080/is_valid
```
