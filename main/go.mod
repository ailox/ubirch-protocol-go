module main

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.2
	github.com/google/uuid v1.1.1
	github.com/paypal/go.crypto v0.1.0
	github.com/ubirch/ubirch-protocol-go/ubirch v0.0.0
)

replace github.com/ubirch/ubirch-protocol-go/ubirch => ../ubirch
