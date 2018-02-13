package main

import (
	"crypto/rand"
	"log"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"bytes"
	"math/big"
	"time"
)

func IntToHex(num uint64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Generate some random block data
/*

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}
 */

const targetBits = 16

func main() {

	blockData := make([]byte, 100)


	if _, err := rand.Read(blockData); err != nil {
		log.Fatalln("Random Token Generator failed")
	}

	sha_256 := sha256.New()

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))



	var nonce uint64 = 0

	before := time.Now()

	for {

		data := bytes.Join([][]byte{IntToHex(nonce), blockData}, []byte{}) // append(nonce, token...)

		sha_256.Write([]byte(data))

		var hashInt big.Int

		hash := sha_256.Sum(nil)

		fmt.Printf("\rHashing: %x", hash)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			now := time.Now()
			fmt.Printf("\nBlock Data:\t%X", blockData)
			fmt.Printf("\nHash:\t%X", hash)
			fmt.Printf("\nNonce:\t%v", nonce)

			fmt.Printf("\nHashing time:\t%.3f\n\r sec", now.Sub(before).Seconds())
			break
		} else {
			nonce++
		}

	}
}
