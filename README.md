# Gen Resume

A command-line tool written in Go that generates professional PDF resumes from JSON data and HTML templates.

## Features

- **JSON-based data input** - Separate your resume content from presentation
- **Customizable HTML templates** - Use Go template syntax for flexible layouts
- **PDF generation** - Creates high-quality PDFs optimized for A4 printing
- **Professional design** - Includes a default template with modern styling
- **Two-column skills layout** - Separates technical and soft skills
- **Print-optimized** - Ensures resume fits perfectly on a single A4 page

## Installation

### Prerequisites

- Go 1.16 or higher
- Chrome or Chromium browser (for PDF generation)

### Build from source

```bash
git clone https://github.com/yourusername/gen-resume.git
cd gen-resume
go mod download
go build -o gen-resume main.go
```

## Usage

```bash
gen-resume -json <path> [-template <path>] [-output <path>]
```

### Options

- `-json` (required) - Path to JSON file containing resume data
- `-template` - Path to HTML template file (default: `template.html`)
- `-output` - Path for output PDF file (default: `resume.pdf`)
- `-help` - Show help message

### Example

```bash
gen-resume -json example.json -template template.html -output john_doe_resume.pdf
```

## JSON Data Structure

The JSON file should follow this structure:

```json
{
  "name": "John Doe",
  "title": "Senior Software Engineer",
  "email": "john.doe@email.com",
  "phone": "(555) 123-4567",
  "location": "San Francisco, CA",
  "linkedin": "linkedin.com/in/johndoe",
  "github": "github.com/johndoe",
  "website": "johndoe.com",
  "summary": "Professional summary text...",
  "experience": [
    {
      "title": "Job Title",
      "company": "Company Name",
      "startDate": "June 2020",
      "endDate": "Present",
      "responsibilities": [
        "Achievement or responsibility 1",
        "Achievement or responsibility 2"
      ]
    }
  ],
  "technicalSkills": [
    {
      "category": "Languages",
      "items": ["JavaScript", "Python", "Go"]
    }
  ],
  "softSkills": [
    {
      "category": "Leadership",
      "items": ["Team Management", "Mentoring"]
    }
  ],
  "education": [
    {
      "degree": "Bachelor of Science in Computer Science",
      "school": "University Name",
      "startDate": "2012",
      "endDate": "2016"
    }
  ]
}
```

All sections are optional. See `example.json` for a complete example.

## Templates

The tool uses Go's `html/template` package. You can create custom templates by modifying `template.html` or creating your own.

### Template Variables

- Basic fields: `{{.Name}}`, `{{.Title}}`, `{{.Email}}`, etc.
- Collections: `{{range .Experience}}...{{end}}`
- Conditionals: `{{if .Summary}}...{{end}}`

## Project Structure

```
gen-resume/
├── main.go           # Main application code
├── template.html     # Default resume template
├── example.html      # Static HTML example
├── example.json      # Sample resume data
├── go.mod           # Go module file
└── README.md        # This file
```

## Dependencies

- [chromedp](https://github.com/chromedp/chromedp) - Chrome DevTools Protocol for PDF generation
- No external binaries required (uses system Chrome/Chromium)

## License

MIT
