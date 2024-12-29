package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/xlzd/gotp"
)

var Otps map[string]string

func PrintAllOtps(mp *map[string]string) {
	for k, v := range *mp {
		otp, expiredTimestamp := gotp.NewDefaultTOTP(string(v)).NowWithExpiration()
		fmt.Println("ET(sec):", expiredTimestamp-time.Now().Unix(), "- OTP:", otp, "- OTP Name:", k)
	}
}

func main() {
	// Load and unmarshall json with otp keys
	content, err := ioutil.ReadFile("./keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	json.Unmarshal([]byte(content), &Otps)

	fmt.Println(Otps)
	PrintAllOtps(&Otps)

}
