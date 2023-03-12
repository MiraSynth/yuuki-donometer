package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//go:embed assets/meter.html
var meterHTML string

//go:embed assets/alert.html
var alertHTML string

//go:embed assets/headpat-yuuki.gif
var yuukiGIF string

//go:embed assets/banana.mp3
var bananaMP3 string

func main() {
	fmt.Print("\n\n")
	fmt.Println("                        Welcome to the Yuuki-Donometer!")
	fmt.Println("════════════════════════════════════════════════════════════════════════════════")

	fmt.Print("\n")

	fmt.Println("  What is this?:")
	fmt.Println("  ════════════════════════════════════════════════════════════")
	fmt.Println("    A little app that connects to GoFundMe, gets donation data ")
	fmt.Println("    and displays a meter and an alert box with sounds and")
	fmt.Println("    everything!")
	fmt.Print("\n")
	fmt.Println("    Consider donating to Yuuki: https://gofund.me/60c780d1")

	fmt.Print("\n\n")

	fmt.Println("  Pages:")
	fmt.Println("  ════════════════════════════════════════════════════════════")
	fmt.Println("    #Alert Box")
	fmt.Println("     Copy the url into an OBS browser source.")
	fmt.Println("         - http://localhost:8969/alert")
	fmt.Print("\n")
	fmt.Println("     To style add the following to the Custom CSS field")
	fmt.Println("         :root {")
	fmt.Println("             --alert-text-color: #ffffff;")
	fmt.Println("             --alert-text-size: 3em;")
	fmt.Println("         }")

	fmt.Print("\n\n")

	fmt.Println("    #Donation Meter")
	fmt.Println("     Copy the url into an OBS browser source.")
	fmt.Println("         - http://localhost:8969/meter")
	fmt.Print("\n")
	fmt.Println("     To style add the following to the Custom CSS field")
	fmt.Println("         :root {")
	fmt.Println("             --meter-height: 3em;")
	fmt.Println("             --meter-color: #ffb9ea;")
	fmt.Println("             --meter-progress-color: #ff61d0;")
	fmt.Println("             --meter-text-color: #000000;")
	fmt.Println("             --meter-text-size: 0.8em;")
	fmt.Println("         }")

	fmt.Print("\n\n")

	fmt.Println("  Notes:")
	fmt.Println("  ════════════════════════════════════════════════════════════")
	fmt.Println("    Remember to uncheck the following in the OBS Browser Source")
	fmt.Println("      - Shutdown source when not visible")
	fmt.Println("      - Refresh browser when scene becomes active")
	fmt.Print("\n")
	fmt.Println("    Do not close this Application, just minimize it!")

	fmt.Print("\n")
	fmt.Println("════════════════════════════════════════════════════════════════════════════════")
	fmt.Println("                Written by @MiraSynth, for Yuuki with love! <3")
	fmt.Print("\n\n")

	http.HandleFunc("/meter", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, meterHTML)
	})

	http.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, alertHTML)
	})

	http.HandleFunc("/assets/headpat-yuuki.gif", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "image/gif")
		fmt.Fprint(w, yuukiGIF)
	})

	http.HandleFunc("/assets/banana.mp3", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "audio/mpeg")
		fmt.Fprint(w, bananaMP3)
	})

	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		data := getInfo()
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, data)
	})

	http.HandleFunc("/api/donations", func(w http.ResponseWriter, r *http.Request) {
		data := getDonations()
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, data)
	})

	log.Fatal(http.ListenAndServe(":8969", nil))
}

func getInfo() string {
	url := "https://gateway.gofundme.com/web-gateway/v1/feed/63yxmf-helping-yuuki-with-her-new-cancer-treatments/counts"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(body)
}

func getDonations() string {
	url := "https://gateway.gofundme.com/web-gateway/v1/feed/63yxmf-helping-yuuki-with-her-new-cancer-treatments/donations?limit=1&offset=0&sort=recent"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(body)
}
