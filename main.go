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

func PrintOtp(otpName *string, mp *map[string]string) {
	if *otpName == "" {
		for k, v := range *mp {
			otp, eTS := gotp.NewDefaultTOTP(string(v)).NowWithExpiration()
			fmt.Println("ET(sec):", eTS-time.Now().Unix(), "- OTP:", otp, "- OTP Name:", k)
		}
	} else {
		otp, eTS := gotp.NewDefaultTOTP((*mp)[*otpName]).NowWithExpiration()
		fmt.Println("ET(sec):", eTS-time.Now().Unix(), "- OTP:", otp, "- OTP Name:", *otpName)
	}
}

func AddNewOtp(otpName *string, otpKey *string, mp *map[string]string) {
	if otpName == "" || otpKey == "" {
		log.Fatal("The name or key must not be empty.")
	} else {
		fmt.Println(*mp)
	}
}

func main() {
	// Command-line flags
	otpName := flag.String("name", "", "otp name")
	otpKey := flag.String("key", "", "otp key")
	otpAct := flag.String("act", "print", "action: add, del or print")

	flag.Parse()
	fmt.Println("name:", *otpName)
	fmt.Println("key:", *otpKey)
	fmt.Println("act:", *otpAct)

	// Load and unmarshall json with otp keys
	content, err := ioutil.ReadFile("./keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	json.Unmarshal([]byte(content), &Otps)

	fmt.Println(Otps)

	switch *otpAct {
	case "print":
		PrintOtp(otpName, &Otps)
	case "add":
		AddNewOtp(otpName, otpKey, &Otps)
	default:
		fmt.Println("hello")
	}
}
