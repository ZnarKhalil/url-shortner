# URL Shortener in Go

This is a simple URL shortener application built using Go (Golang). It allows users to shorten long URLs and then redirects users to the original URL using the short code.

## Features

- **URL Shortening:** Takes a long URL as input and generates a short, unique code.
- **Redirection:** Redirects users from a short URL to the original long URL.
- **In-Memory Storage:** Uses an in-memory map to store the short URL to long URL mappings. Note that the data will not persist after the server restarts.

## Getting Started

### Installation and Running the Application

1.  **Clone the repository:** (If you are using git, otherwise simply download the zip file)

    ```bash
    git clone git@github.com:ZnarKhalil/url-shortner.git
    cd url-shortner
    ```

2.  **Run the application:**

    ```bash
    go run main.go
    ```

    This command starts the server on `http://localhost:8080`.

## Usage

### Shortening a URL

1.  Send a POST request to `/shorten` with the long URL as a form parameter called `url`.
2.  You can use `curl`, Postman, or other HTTP clients.

    **Example using `curl`:**

    ```bash
    curl -X POST -d "url=https://www.example.com" http://localhost:8080/shorten
    ```

3.  The server will respond with the shortened URL, such as `Shortened URL: http://localhost:8080/<shortCode>`.

### Redirecting to a Long URL

1.  Access the shortened URL in your web browser or use `curl`.
2.  The server will automatically redirect you to the corresponding long URL.

    **Example:** Access the URL `http://localhost:8080/<shortCode>` in your browser and will redirect you to the long URL.

## Code Overview

### `main.go`

- **`urlMap`:** A map (dictionary) to store the short URL and their corresponding long URL.
- **`main()` Function:**
  - Starts the HTTP server and listens on port 8080.
  - Registers the `redirectURLHandler` for the root path (`/`) to handle redirections.
  - Registers the `shortenURLHandler` for the `/shorten` path to handle URL shortening requests.
- **`generateShortCode()` Function:**
  - Generates a random, cryptographically secure short code.
  - Uses `crypto/rand` and `base64.RawURLEncoding` to generate a URL-safe random string of 12 characters.
- **`shortenURLHandler()` Function:**
  - Handles POST requests to the `/shorten` endpoint.
  - Extracts the long URL from the request.
  - Generates a unique short code using `generateShortCode()`.
  - Saves the short-to-long URL mapping in the `urlMap`.
  - Returns the generated short URL to the user.
- **`redirectURLHandler()` Function:**
  - Handles GET requests to any path of the server (`/` path will be catched here).
  - Extracts the short code from the URL path.
  - If the short code exists in the `urlMap`, it redirects the user to the corresponding long URL.
  - If the short code is not found, it returns a `404 Not Found` error.

## Considerations

- **In-Memory Storage:** The current implementation stores URL mappings in memory. This means the mappings will be lost if the server restarts.
- **No collision detection:** This implementation does not have collision detection.
- **Error Handling:** The error handling is basic. You can improve the error responses and add logging.

## Future Enhancements

- **Persistent Storage:** Implement database or file-based persistence.
- **Custom Short Codes:** Allow users to specify custom short codes.
- **More Robust Random Code Generation:** Implement some collision detection in the short code generation
- **Improved Error Handling:** Improve the error messages and logging.
- **Input Validation:** Validate the input URLs.
