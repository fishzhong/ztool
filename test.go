package ztool

type account struct {
	balance float32
}

func Foreach(accounts []*account) {
	for _, a := range accounts {
		a.balance += 1000
	}
}

var accounts = []*account{
	{balance: 100.},
	{balance: 200.},
	{balance: 300.},
}

var account2 = []account{
	{balance: 100.},
	{balance: 200.},
	{balance: 300.},
}

func Foreach2(accounts2 []account) {
	for i := range accounts2 { // ❶
		accounts[i].balance += 1000
	}
}
