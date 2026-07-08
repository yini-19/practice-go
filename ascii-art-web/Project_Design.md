
Team Responsibilities

# Team Member: lantana Yusuf 
ASCII Engine & Validation Team Member: 
Responsibilities
- Load banner files
- Build ASCII rendering logic
- Input validation
- Error handling
- Unit tests for rendering
Files
internal/ascii/
├── loader.go
├── render.go
|---split.go
|---generate.go
└── validate.go


# Team Member: Yiyakazah Nicodemus 
HTTP Server & Handlers

Responsibilities
- Server setup
- Route registration
- GET and POST handlers
- Status code handling
- Connect handlers to ASCII engine
- Handler tests
Files
cmd/web/
└── main.go

internal/handlers/
├── home.go
├── ascii.go
└── errors.go


# Team Member : Michael Bulus
Templates, Documentation & Testing
Responsibilities
- HTML templates
- Form creation
- Result page
- Error pages
- README
- End-to-end testing
Files
templates/
├── index.html
├── result.html
├── error.html

README.md


# Project Structure
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
|   |   |── split.go
|   |   |── generate.go
│   │
│   └── handlers/
│       ├── home.go
│       ├── ascii.go
│       └── errors.go
│
├── templates/
│   ├── index.html
│   ├── result.html
│   ├── error.html
│   
│   
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


# Feature Breakdown

## Build:
func LoadBanner(name string) (map[rune]string,error)
func RenderLines(input, banner string) [] string
func Generate(input string, banner map[rune]string) []string
func ValidateInput(input, banner string) error
func SplitInput(input string) []string
func HomeHandler(w http.ResponseWriter, r *http.Request)
func ASCIIHandler(w http.ResponseWriter, r *http.Request)
func RenderError(w http.ResponseWriter, r *http.Request)


## Register routes:
http.HandleFunc("/", HomeHandler)
http.HandleFunc("/ascii-art", ASCIIHandler)



## Create templates:
index.html
result.html
Error.html



## Git Branches
main

feature/ascii-engine
feature/http-handlers
feature/templates