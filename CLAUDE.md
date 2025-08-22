# Gen Resume

## Project Overview

Gen Resume is a command-line application written in Go that takes user information and a resume template as input, then generates a professionally formatted PDF resume. The tool aims to streamline the resume creation process by separating content from presentation.

## Core Functionality

- **Input Processing**: Accept user information (personal details, experience, education, skills) via CLI flags, config files, or interactive prompts
- **Template System**: Support customizable resume templates for different formats and styles
- **PDF Generation**: Convert structured data and templates into high-quality PDF output
- **CLI Interface**: Provide an intuitive command-line interface with helpful flags and options

## Technical Requirements

### Language & Framework
- **Primary Language**: Go (latest stable version)
- **PDF Generation**: Evaluate libraries like `gofpdf`, `unidoc`, or `chromedp` for HTML-to-PDF conversion

### Architecture Goals
- Clean, modular code structure
- Separation of concerns (data parsing, template processing, PDF generation)
- Extensible template system
- Error handling and validation
- Cross-platform compatibility

### Input Formats
- JSON or YAML configuration files for resume data
- Command-line flags for quick modifications
- Interactive mode for guided input
- Support for importing from common formats

### Template System
- Template files (HTML/CSS, LaTeX, or custom format)
- Multiple built-in professional templates
- Custom template support
- Variable substitution and conditional sections

## Expected CLI Usage

```bash
# Basic usage with config file
gen-resume --config resume.json --template professional.html --output john_doe_resume.pdf
```

## Quality Standards

- Comprehensive error handling
- Input validation and sanitization
- Unit tests for core functionality
- Integration tests for CLI commands
- Clear documentation and examples
- Cross-platform builds and releases