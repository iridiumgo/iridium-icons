// cmd/fetch_genicons/main.go
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	zipURL = "https://github.com/lucide-icons/lucide/archive/refs/heads/main.zip"
)

// sanitizeName converts "arrow-right.svg" → "ArrowRight"
func sanitizeName(name string) string {
	name = strings.TrimSuffix(name, ".svg")
	var b strings.Builder
	upperNext := true
	for _, r := range name {
		if r == '-' || r == '_' {
			upperNext = true
			continue
		}
		if upperNext {
			b.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// processSVGContent converts SVG content to templ-compatible format
func processSVGContent(svgContent string) string {
	// Remove the opening <svg> tag and closing </svg> tag to extract inner content
	content := strings.TrimSpace(svgContent)

	// Find the end of the opening <svg> tag
	svgStart := strings.Index(content, "<svg")
	if svgStart == -1 {
		return content // Return as-is if no <svg> tag found
	}

	svgTagEnd := strings.Index(content[svgStart:], ">")
	if svgTagEnd == -1 {
		return content
	}
	svgTagEnd += svgStart + 1

	// Find the closing </svg> tag
	svgCloseStart := strings.LastIndex(content, "</svg>")
	if svgCloseStart == -1 {
		return content
	}

	// Extract the inner content between <svg...> and </svg>
	innerContent := strings.TrimSpace(content[svgTagEnd:svgCloseStart])

	// Clean up the inner content for templ
	// Convert self-closing tags to proper format if needed
	innerContent = strings.ReplaceAll(innerContent, "/>", "></path>")
	innerContent = strings.ReplaceAll(innerContent, "<path", "<path")
	innerContent = strings.ReplaceAll(innerContent, "<polyline", "<polyline")
	innerContent = strings.ReplaceAll(innerContent, "<circle", "<circle")
	innerContent = strings.ReplaceAll(innerContent, "<rect", "<rect")
	innerContent = strings.ReplaceAll(innerContent, "<line", "<line")

	// Fix self-closing tags properly
	innerContent = strings.ReplaceAll(innerContent, "></path>", "/>")
	innerContent = strings.ReplaceAll(innerContent, "></polyline>", "/>")
	innerContent = strings.ReplaceAll(innerContent, "></circle>", "/>")
	innerContent = strings.ReplaceAll(innerContent, "></rect>", "/>")
	innerContent = strings.ReplaceAll(innerContent, "></line>", "/>")

	return innerContent
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	outDir := "./icon/icons"
	tempDir := "./temp"
	zipFile := filepath.Join(tempDir, "lucide.zip")

	// Create directories
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("cannot create output directory: %v", err)
	}
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		log.Fatalf("cannot create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up temp directory

	log.Println("Downloading Lucide repository archive...")
	if err := downloadFile(zipURL, zipFile); err != nil {
		log.Fatalf("Failed to download zip: %v", err)
	}

	// Open the zip file
	log.Println("Extracting icons from archive...")
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatalf("Failed to open zip: %v", err)
	}
	defer r.Close()

	iconCount := 0
	for _, f := range r.File {
		// Look for files in the icons directory
		if !strings.Contains(f.Name, "/icons/") || !strings.HasSuffix(f.Name, ".svg") {
			continue
		}

		// Extract just the filename
		parts := strings.Split(f.Name, "/")
		fileName := parts[len(parts)-1]

		// Open file in zip
		rc, err := f.Open()
		if err != nil {
			log.Printf("Error opening %s in zip: %v", fileName, err)
			continue
		}

		// Read SVG content
		svgData, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			log.Printf("Error reading %s: %v", fileName, err)
			continue
		}

		// Process SVG content for templ
		svgContent := string(svgData)
		innerContent := processSVGContent(svgContent)

		// Generate templ component
		componentName := sanitizeName(fileName)
		templFile := filepath.Join(outDir, strings.ToLower(componentName)+".templ")

		templContent := fmt.Sprintf(
			`package icons

import (
    "github.com/iridiumgo/iridium-icons/icon"
)

var %s = icon.NewIcon(%sComponent)

templ %sComponent(i *icon.Icon) {
	<svg
		xmlns="http://www.w3.org/2000/svg"
        width={i.Width}
        height={i.Height}
        viewBox={i.ViewBox}
        fill={i.Fill}
        stroke={i.Stroke}
        stroke-width={i.StrokeWidth}
        stroke-linecap={i.StrokeLineCap}
        stroke-linejoin={i.StrokeLineJoin}
        {i.Attributes...}
	>
		%s
	</svg>
}
`, componentName, componentName, componentName, innerContent,
		)

		if err := os.WriteFile(templFile, []byte(templContent), 0644); err != nil {
			log.Printf("Failed to write templ file %s: %v", templFile, err)
			continue
		}

		iconCount++
		if iconCount%50 == 0 {
			log.Printf("Processed %d icons...", iconCount)
		}
	}

	fmt.Printf("Successfully generated %d templ icon components in %s\n", iconCount, outDir)
	fmt.Println("Run 'templ generate' to compile the templates")
}
