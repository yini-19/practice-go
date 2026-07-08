# ASCII Art Web

## Description

ASCII Art Web is a web application written in Go that converts user-provided text into ASCII art using different banner styles. The application provides a simple web interface where users can enter text, select a banner style, and generate ASCII art output.

Supported banners:

* standard
* shadow
* thinkertoy

The application follows a modular architecture that separates the ASCII rendering engine, HTTP handlers, templates, and testing.

---

## Authors

* Lantana Yusuf
* Yiyakazah Nicodemus
* Michael Bulus


---

## Usage

### Clone the repository

```bash
git clone https://acad.learn2earn.ng/git/mbulus/ascii-art-web
cd ascii-art-web
```

### Run the application

```bash
go run ./cmd/web
```

The server will start on:

```text
http://localhost:8080
```

Open your browser and navigate to the URL above.

---

## Implementation Details

### Algorithm

1. Receive user input through the web form.
2. Validate the submitted text and banner selection.
3. Load the selected banner file.
4. Parse the banner file into a character map.
5. Convert each character of the input text into its ASCII representation.
6. Assemble the ASCII output line by line.
7. Return the generated ASCII art to the user.
8. Display the result in the browser.

---

# Team Responsibilities

## 1. ASCII Engine & Validation

### Team Member

**Lantana Yusuf**

### Responsibilities

* Load banner files
* Build ASCII rendering logic
* Input validation
* Error handling
* Unit tests for rendering

### Files

```text
internal/ascii/
├── loader.go
├── render.go
└── validate.go
```

---

## 2. HTTP Server & Handlers

### Team Member

**Yiyakazah Nicodemus**

### Responsibilities

* Server setup
* Route registration
* GET and POST handlers
* Status code handling
* Connect handlers to ASCII engine
* Handler tests

### Files

```text
cmd/web/
└── main.go

internal/handlers/
├── home.go
├── ascii.go
└── errors.go
```

---

## 3. Templates, Documentation & Testing

### Team Member

**Michael Bulus**

### Responsibilities

* HTML templates
* Form creation
* Result page
* Error pages
* README
* End-to-end testing

### Files

```text
templates/
├── index.html
├── result.html
└── error.html

README.md
```

---

# Project Structure

```text
ascii-art-web/
│
├── cmd/
│   └── web/
│       └── main.go
│
├── internal/
│   ├── ascii/
│   │   ├── loader.go
│   │   ├── render.go
│   │   └── validate.go
│   │
│   └── handlers/
│       ├── home.go
│       ├── ascii.go
│       └── errors.go
│
├── templates/
│   ├── index.html
│   ├── result.html
│   └── error.html
│
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── tests/
│   ├── render_test.go
│   └── handler_test.go
│
├── go.mod
├── go.sum
└── README.md
```

---

# Feature Breakdown

## ASCII Engine

### Building

```go
func LoadBanner(name string) error

func RenderLines(input, banner string) (string, error)

func ValidateInput(input, banner string) error
```

---

## HTTP Handlers

### Build

```go
func HomeHandler(w http.ResponseWriter, r *http.Request)

func ASCIIHandler(w http.ResponseWriter, r *http.Request)
```

### Register Routes

```go
http.HandleFunc("/", HomeHandler)
http.HandleFunc("/ascii-art", ASCIIHandler)
```

---

## Templates

### Create Templates

```text
index.html
result.html
error.html
```

---

# Git Branches

```text
main

feature/ascii-engine

feature/http-handlers

feature/templates
```

