package main

import (
	//"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, Thank you for testing s3cr3tx with Go")
	var email string = "pk@gratitech.com"
	var APIToken string = "bCrCgRrFoMKyw4PDkyoYF8O9PSXDrsKQRuKAueKAmTbDqsOdacOkw6/Dn8Kzw5/DvRTDjuKAusOyWG3DgG1oUMOBOcObfcOLDwpiwo06CcK2W8OHY3vDmkQzwqTCs1Mzwr/Cvg=="
	var AuthCode string = "w4zCtm8fH1o7wr/DuB3FoeKAmDtfw5ZvwrLihKLigLA9XcOeMcKtHRYBy4Z5GhTDuMWgwqZoUsOnXR4hCsuG4oChV1HigLDDgMK4w6HFoTbDhlksw7gowrXDhsOMUMOKw4BmFg=="
	var baseurl string = "https://s3cr3tx.com/Values"

	req, err := http.NewRequest(http.MethodGet, baseurl, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Email", email)
	req.Header.Set("APIToken", APIToken)
	req.Header.Set("AuthCode", AuthCode)
	req.Header.Set("EorD", "e")
	req.Header.Set("Input", "This is something secret")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Your encrypted text is : " + string(resBody))

	req2, err := http.NewRequest(http.MethodGet, baseurl, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req2.Header.Set("Content-Type", "text/plain")
	req2.Header.Set("Accept", "text/plain")
	req2.Header.Set("Email", email)
	req2.Header.Set("APIToken", APIToken)
	req2.Header.Set("AuthCode", AuthCode)
	req2.Header.Set("EorD", "d")
	req2.Header.Set("Input", string(resBody))
	client2 := http.Client{
		Timeout: 30 * time.Second,
	}

	res2, err := client2.Do(req2)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	resBody2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Your decrypted text is : " + string(resBody2))

	fmt.Println("Done with s3cr3tx Go Test!")
}
