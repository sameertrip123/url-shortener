# URL Shortener

This is a custom URL shortener project built using Go and the Fiber web framework. It provides a simple and efficient way to generate shortened URLs for convenient sharing and tracking purposes.

## Features

- Generate unique shortened URLs for long URLs
- Custom alias support for personalized shortened URLs
- Efficient redirection from shortened URLs to the original long URLs
- Rate limiting functionality to restrict users to a maximum of 10 API calls in 30 minutes
- URL validation and checking to ensure only valid and accessible URLs are accepted

## Technologies Used

- Go programming language
- Fiber web framework
- Redis database

## Getting Started

### Prerequisites

- Go (version 1.20)
- Redis (version v9)



## Usage

1. To shorten a URL, send a `POST` request to `/api/v1` with the following parameters:
   - `url`: The long URL you want to shorten.
   - `alias` (optional): Custom alias in the body for the shortened URL.

2. The API response will contain the shortened URL.

3. Use the shortened URL to redirect users to the original long URL by accessing it in a web browser.

Example using cURL:

### Request

`POST /api/v1/shortenthislongurl`

    curl.exe -X POST -d http://localhost:3000/api/v1/shortenthislongurl

### Response

    {
        "shortened_url": "http://localhost:3000/abc123"
    }



