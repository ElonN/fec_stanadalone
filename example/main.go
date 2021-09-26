package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// kcp receiving is based on packets
	// recvbuf turns packets into stream
	recvbuf []byte
	bufptr  []byte

	// FEC codec
	fecDec *fecDecoder
	fecEnc *fecEncoder

	// default values
	dataShards   = 10
	parityShards = 3

	// packet size range
	packet_size_low      = 300
	packet_size_variance = 1000
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func get_packet() []byte {
	packet_len := packet_size_low + rand.Intn(packet_size_variance)
	packet := make([]byte, packet_len)
	for i := range packet {
		packet[i] = byte(letters[rand.Intn(len(letters))])
	}
	return packet
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// FEC codec initialization
	fecDec = newFECDecoder(dataShards, parityShards)
	fecEnc = newFECEncoder(dataShards, parityShards, 0)

	packet_arr := make([][]byte, 20)
	var parity_shards [][]byte

	for i := 0; i < 20; i++ {
		packet_arr[i] = get_packet()
	}

	for _, p := range packet_arr {
		parity_shards = fecEnc.encode(p)
		fmt.Println("Data: ", string(p)[:100])
		fmt.Println("Parity: ", parity_shards)
	}

}
