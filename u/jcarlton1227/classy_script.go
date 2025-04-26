package inner

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url" // To parse the URL for the filename
	"os"
	"path/filepath" // To extract the filename from the path

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

func main() {
	download()

	f, err := os.Open("./demo.dem")
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := demoinfocs.NewParser(f)
	defer p.Close()

	// Register handler on kill events
	p.RegisterEventHandler(func(e events.Kill) {
		var hs string
		if e.IsHeadshot {
			hs = " (HS)"
		}
		var wallBang string
		if e.PenetratedObjects > 0 {
			wallBang = " (WB)"
		}
		fmt.Printf("%s <%v%s%s> %s\n", e.Killer, e.Weapon, hs, wallBang, e.Victim)
	})

	// Parse to end
	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}
}

func download() {
	// --- Configuration ---
	// Replace with the actual URL of the file you want to download
	fileURL := "https://bucket-production-0799.up.railway.app/cs-demos/iem-melbourne-2025/liquid-vs-natus-vincere-m1-anubis.dem" // Example URL

	// Optional: Specify an output path. If empty, it will try to use the
	// filename from the URL in the current directory.
	outputFilePath := "./demo.dem"
	// --- End Configuration ---

	// --- Determine Output Filename ---
	if outputFilePath == "" {
		parsedURL, err := url.Parse(fileURL)
		if err != nil {
			fmt.Printf("Error parsing URL '%s': %v\n", fileURL, err)
			os.Exit(1)
		}
		// Use the base name of the URL path as the filename
		outputFilePath = filepath.Base(parsedURL.Path)
		if outputFilePath == "." || outputFilePath == "/" || outputFilePath == "" {
			fmt.Printf("Could not determine filename from URL '%s'. Please specify outputFilePath.\n", fileURL)
			os.Exit(1)
		}
		fmt.Printf("Output filename not specified, using filename from URL: '%s'\n", outputFilePath)
	}

	fmt.Printf("Attempting to download '%s' to '%s'...\n", fileURL, outputFilePath)

	// --- Create the Output File ---
	// os.Create creates a file if it doesn't exist, or truncates it if it does.
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating file '%s': %v\n", outputFilePath, err)
		os.Exit(1)
	}
	// Ensure the file is closed before main() exits
	defer outFile.Close()

	// --- Make the HTTP GET Request ---
	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Printf("Error making GET request to '%s': %v\n", fileURL, err)
		// Clean up the potentially empty file created earlier
		os.Remove(outputFilePath)
		os.Exit(1)
	}
	// Ensure the response body is closed before main() exits
	defer resp.Body.Close()

	// --- Check Server Response ---
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received non-200 status code: %s\n", resp.Status)
		// Clean up the potentially empty file created earlier
		os.Remove(outputFilePath)
		os.Exit(1)
	}

	// --- Copy Response Body to File ---
	// io.Copy efficiently copies data from the response body (reader)
	// to the output file (writer). It handles large files well.
	bytesCopied, err := io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Printf("Error copying content to file '%s': %v\n", outputFilePath, err)
		// Clean up the potentially partially written file
		os.Remove(outputFilePath)
		os.Exit(1)
	}

	fmt.Printf("Successfully downloaded %d bytes to '%s'.\n", bytesCopied, outputFilePath)
}
