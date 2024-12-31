package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/xlzd/gotp"
)

var Otps map[string]string
var Keys string = "/Users/dg/t/golearn/cotp/keys.json"

func JsonWrite(mp *map[string]string) {
	jsonData, err := json.Marshal(*mp)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(Keys, jsonData, 0644)
}

func PrintOtp(otpName *string, mp *map[string]string) {
	if *otpName == "" {
		for k, v := range *mp {
			otp, eTS := gotp.NewDefaultTOTP(string(v)).NowWithExpiration()
			fmt.Println("ET(sec):", eTS-time.Now().Unix(), "- OTP:", otp, "- OTP Name:", k)
		}
	} else {
		if (*mp)[*otpName] == "" {
			log.Fatal("Not found otp name: " + *otpName)
		} else {
			otp, eTS := gotp.NewDefaultTOTP((*mp)[*otpName]).NowWithExpiration()
			fmt.Println("ET(sec):", eTS-time.Now().Unix(), "- OTP:", otp, "- OTP Name:", *otpName)
		}
	}
}

func AddNewOtp(otpName *string, otpKey *string, mp *map[string]string) {
	if *otpName == "" || *otpKey == "" {
		log.Fatal("The name or key must not be empty.")
	} else {
		(*mp)[*otpName] = *otpKey
		JsonWrite(mp)
	}
}

func DelOtp(otpName *string, mp *map[string]string) {
	if *otpName == "" {
		log.Fatal("The name must not be empty.")
	} else {
		delete(*mp, *otpName)
		JsonWrite(mp)
	}
}

func main() {
	// Command-line flags
	otpName := flag.String("name", "", "otp name")
	otpKey := flag.String("key", "", "otp key")
	otpAct := flag.String("act", "print", "action: add, del or print")
	flag.Parse()

	// Load and unmarshall json with otp keys
	content, err := ioutil.ReadFile(Keys)
	if err != nil {
		Otps = make(map[string]string)
		JsonWrite(&Otps)
	}
	json.Unmarshal([]byte(content), &Otps)

	// Performing an action on the keys
	switch *otpAct {
	case "print":
		PrintOtp(otpName, &Otps)
	case "add":
		AddNewOtp(otpName, otpKey, &Otps)
	case "del":
		DelOtp(otpName, &Otps)
	}
}
