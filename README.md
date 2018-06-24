# market

market is a go library meant to facilitate customers' purchase of Products offered at a market, applying various Specials upon Checkout.

to run,

`git clone https://github.com/coip/market.git`  
`cd market`  
`docker build -t market . && docker-compose up`  

includes rudimentary integration tests, found in [*market_test.go*](https://github.com/coip/market/blob/master/market_test.go), which should pass once the market_server is up.

verify success via `go test -v` once the market_server is up locally.

use [*godemon.sh*](https://github.com/coip/market/blob/master/godemon.sh) to rebuild/compose on {`"*.go"`,`"*/*.go"`} filechanges. (NOTE: just a crufty ~nodemon tool cause I'm a fan.)

should be fine as a standalone executable, though Docker and various Linux/Unix tools may improve the development lifecycle experience.