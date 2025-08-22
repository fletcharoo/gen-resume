package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Resume represents the complete resume data structure
type Resume struct {
	Name            string          `json:"name"`
	Title           string          `json:"title"`
	Email           string          `json:"email"`
	Phone           string          `json:"phone"`
	Location        string          `json:"location"`
	LinkedIn        string          `json:"linkedin"`
	GitHub          string          `json:"github"`
	Website         string          `json:"website"`
	Summary         string          `json:"summary"`
	Experience      []Experience    `json:"experience"`
	TechnicalSkills []SkillCategory `json:"technicalSkills"`
	SoftSkills      []SkillCategory `json:"softSkills"`
	Education       []Education     `json:"education"`
	Projects        []Project       `json:"projects"`
	Certifications  []Certification `json:"certifications"`
	Languages       []Language      `json:"languages"`
}

// Experience represents work experience
type Experience struct {
	Title            string   `json:"title"`
	Company          string   `json:"company"`
	StartDate        string   `json:"startDate"`
	EndDate          string   `json:"endDate"`
	Responsibilities []string `json:"responsibilities"`
}

// SkillCategory represents a category of skills
type SkillCategory struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

// Education represents educational background
type Education struct {
	Degree      string `json:"degree"`
	School      string `json:"school"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
}

// Project represents a project
type Project struct {
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
}

// Certification represents a certification
type Certification struct {
	Name   string `json:"name"`
	Issuer string `json:"issuer"`
	Date   string `json:"date"`
}

// Language represents a language skill
type Language struct {
	Name        string `json:"name"`
	Proficiency string `json:"proficiency"`
}

func main() {
	// Define command-line flags
	var (
		jsonPath     = flag.String("json", "", "Path to JSON file containing resume data (required)")
		templatePath = flag.String("template", "template.html", "Path to HTML template file")
		outputPath   = flag.String("output", "resume.pdf", "Path for output PDF file")
		help         = flag.Bool("help", false, "Show help message")
	)

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Gen Resume - Generate professional PDF resumes from JSON data\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s -json <path> [-template <path>] [-output <path>]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -json resume.json -template template.html -output john_doe_resume.pdf\n", os.Args[0])
	}

	flag.Parse()

	// Show help if requested
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Validate required arguments
	if *jsonPath == "" {
		fmt.Fprintf(os.Stderr, "Error: JSON file path is required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Process the resume
	if err := generateResume(*jsonPath, *templatePath, *outputPath); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✓ Resume generated successfully: %s\n", *outputPath)
}

func generateResume(jsonPath, templatePath, outputPath string) error {
	// Read JSON file
	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	// Parse JSON into Resume struct
	var resume Resume
	if err := json.Unmarshal(jsonData, &resume); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Read template file
	templateData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse and execute template
	tmpl, err := template.New("resume").Parse(string(templateData))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template with resume data
	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, resume); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Convert HTML to PDF
	if err := htmlToPDF(htmlBuffer.String(), outputPath); err != nil {
		return fmt.Errorf("failed to convert HTML to PDF: %w", err)
	}

	return nil
}

func htmlToPDF(html string, outputPath string) error {
	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create temporary HTML file
	tmpFile, err := ioutil.TempFile("", "resume-*.html")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(html)); err != nil {
		return fmt.Errorf("failed to write HTML to temp file: %w", err)
	}
	tmpFile.Close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Configure Chrome options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()

	taskCtx, taskCancel := chromedp.NewContext(allocCtx)
	defer taskCancel()

	// Generate PDF
	var pdfBuf []byte
	if err := chromedp.Run(taskCtx,
		chromedp.Navigate("file://"+tmpFile.Name()),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Print to PDF with A4 size
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).  // A4 width in inches
				WithPaperHeight(11.69). // A4 height in inches
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				WithPreferCSSPageSize(true).
				Do(ctx)
			if err != nil {
				return err
			}
			pdfBuf = buf
			return nil
		}),
	); err != nil {
		return fmt.Errorf("failed to generate PDF: %w", err)
	}

	// Write PDF to file
	if err := ioutil.WriteFile(outputPath, pdfBuf, 0644); err != nil {
		return fmt.Errorf("failed to write PDF file: %w", err)
	}

	return nil
}