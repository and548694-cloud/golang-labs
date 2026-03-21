package main

type Currency interface {
	Name() string
	USDRate() float64
	IsCrypto() bool
}

type USD struct{}

func (u USD) Name() string {
	return "USD"
}

func (u USD) USDRate() float64 {
	return 1.0
}

func (u USD) IsCrypto() bool {
	return false
}

type Bitcoin struct{}

func (b Bitcoin) Name() string {
	return "Bitcoin"
}

func (b Bitcoin) USDRate() float64 {
	return 65000.0
}

func (b Bitcoin) IsCrypto() bool {
	return true
}
