package main
import "fmt"

//Rewrite lab1 in Go

var TICKS float64   // Number of Ticks
var lambda float64	// Average number of packets generated / arrived (packets per second)
var L float64		// Length of a packet in bits
var C float64		// The service time received by a packet. (Ex: The transmission rate of the output link in bits per second.)
var p float64		// Utilization of the queue (= input rate/service rate = L λ/C)
var E_N float64		// Average number of packets in the buffer/queue
var E_T float64		// Average sojourn time or queuing delay + service time spent by the packet in the system
var P_IDLE float64	// The proportion of time the server is idle
var P_LOSS float64	// The packet loss probability (for M/D/1/k queue). It is the ratio of the total number of packets lost due to buffer full condition to the total number of packets generated.
var M float64		// The number of times you repeat your experiments
var K float64		// Size of the buffer

type Packet struct{
      generated_tick uint32
      served_tick uint32
}

func main(){
	fmt.Println("HW")
}