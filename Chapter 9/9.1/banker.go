/*
Exercise 9.1: Add a functionWithdraw(amount int) boolto thegopl.io/ch9/bank1program.
The result should indicate whether the transaction succeeded or failed due to insufficient funds.
The message sent to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back toWithdraw.
*/
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan int) //receive amount to withdraw
var result = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraw <- amount
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case plan := <-withdraw:
			if balance >= plan {
				balance -= plan
				result <- true
			} else {
				result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
