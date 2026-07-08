# ASCII Banner Web Server

## Description

ASCII Banner Web Server is a Go web application that converts plain text into
ASCII art and serves it over HTTP. Users send a text string and choose a font
style; the server looks up each character in a pre-loaded banner file, assembles
the corresponding 8-row art block, and returns the rendered result as plain text.

The project is useful for generating stylised terminal output, decorative page
headers, or any context where human-readable ASCII art is preferred over images.
Because the renderer is driven by plain text banner files, new font styles can be
added without changing any Go code.

---

## Authors

- **Yiyakazah Nicodemus** — back-end logic and HTTP layer
- **lantana yusuf** — banner file loading and ascii functions

---

## Usage

Follow every step below in order on a machine that has never run this project.
No prior setup is assumed beyond having a working internet connection.

### 1. Install Go

Check whether Go is already installed:

```bash
go version
```

This project requires **Go 1.21 or later**. If Go is missing or older, install it
from the official download page:

```
https://go.dev/dl/
```

After installing, confirm the version:

```bash
go version
# Expected output: go version go1.21.x <OS/arch>
```

### 2. Clone the repository

```bash
git clone https://acad.learn2earn.ng/git/ynicodem/ascii-art-web.git
cd ascii-art-web
```

### 3. Confirm the project layout

```bash
ls
```

You should see at minimum:

```
banners/
  standard.txt
  shadow.txt
  thinkertoy.txt

cd/
  ascii.go

templates/
  index.html
main.go
go.mod
README.md
```

### 4. Download dependencies

```bash
go mod tidy
```

If the project has no external dependencies this command still validates
`go.mod` and `go.sum` and is safe to run.

### 5. Run the server

```bash
go run .
```

Expected terminal output:

```
Server started on http://localhost:8080
```

The server listens on **port 8080**. Leave this terminal open.

### 6. Open the application

Open a browser and visit:

```
http://localhost:8080
```

You will see a form where you can type text and select a banner style
(`standard`, `shadow`, or `thinkertoy`). Submit the form to receive the ASCII
art output.

### 7. Test via curl (optional)

```bash
curl -X POST http://localhost:8080/ascii-art \
  -d "text=Hello&banner=standard"
```

The response is plain text containing the rendered ASCII art.

### 8. Run the test suite

Open a second terminal in the project directory and run:

```bash
go test -v ./...
```

All tests should pass with output ending in `PASS`.

### 9. Stop the server

Return to the first terminal and press `Ctrl + C`.

---

## Implementation Details

### Banner file format

Each banner file (e.g. `banners/standard.txt`) encodes every printable ASCII
character from space (decimal 32) to tilde (decimal 126) — 95 characters in
total.

A single character is represented as **exactly 8 lines of text** followed by
**one blank line** as a separator. The final character in the file has no
trailing blank line. Example — the letter `A` in `standard.txt`:

```
 /\
/  \
/ /\ \
/ ____ \
/_/    \_\




```

The 8-row height is fixed across all banner files so that the renderer can
always align multiple characters side by side on the same output rows.

### Character-to-ASCII mapping (`LoadBanner`)

`LoadBanner` in `banner.go` reads a banner file line by line using
`bufio.Scanner` and builds a `map[rune][]string` where each key is a Unicode
code point and each value is the 8-element slice of art rows for that character.

The mapping algorithm:

1. Start with `charcode = 32` (the space character).
2. Accumulate non-blank lines into a temporary `lines` slice.
3. When a blank line is encountered and `lines` is non-empty, store
   `banner[rune(charcode)] = lines`, reset `lines`, and increment `charcode`.
4. After the scanner finishes, store any remaining `lines` for the last
   character (the tilde, code 126).
5. Validate: if both `banner` and `lines` are empty the file is blank and
   `os.ErrInvalid` is returned; if the final `lines` slice does not have
   exactly 8 entries the file is malformed and `os.ErrInvalid` is returned.

### ASCII art rendering

The HTTP handler receives a text string and a banner name. It calls
`LoadBanner` for the chosen font, then renders the output row by row:

1. Split the input on `\n` to handle multi-line input.
2. For each input line, iterate over rows 0–7.
3. For each row, concatenate `banner[char][row]` for every character in the
   line.
4. Append the assembled row to the result buffer.
5. Join all rows with newlines and write the result to the HTTP response with
   `Content-Type: text/plain`.

This row-first traversal means all characters on one input line are rendered
in parallel across their 8 rows, producing correctly aligned ASCII art
regardless of character width.

### HTTP layer

`main.go` starts a standard `net/http` server on port 8080 and registers two
routes:

| Route | Method | Purpose |
|---|---|---|
| `/` | GET | Serves the HTML form (`templates/index.html`) |
| `/ascii-art` | POST | Accepts text and banner form fields, returns ASCII art |

The handler validates both parameters before calling the renderer. If `text` is
empty it returns a 400 Bad Request; if the requested banner file cannot be
loaded it returns a 500 Internal Server Error with a plain-text error message.