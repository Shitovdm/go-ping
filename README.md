# go-ping

## Getting
`go get gitlab-e.artepay.ru/common/go-ping/Ping`

## Usage  
```go

import "gitlab-e.artepay.ru/common/go-ping/Ping"

func main() {
    go Ping.Serve(8081)
}
```