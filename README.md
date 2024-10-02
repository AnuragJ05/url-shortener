# URL Shortener Service

## Overview

A simple URL shortening service built using Golang and the Gin web framework. This service allows users to shorten URLs, redirect shortened URLs to their original counterparts, and retrieve the top domains that have been shortened the most.


## How It Works

1. **Shorten URL**: Accepts a long URL as input and returns a shortened URL.
2. **Redirect**: When a user accesses the shortened URL, they are redirected to the original long URL.
3. **Top Domains**: Returns the top domains that have been shortened the most.

## APIs

### Shorten URL  
**Endpoint**: `http://localhost:8080/shorten`  
**Method**: `POST`  
**Description**: Shortens the given URL.  
**Request Body**:
```json
{
    "url": "https://example.com"
}
```
**Response Body**:
```json
{
    "url": "http://localhost:8080/pTfS5uXy"
}
```

### Redirect to Original URL
**Endpoint**: `http://localhost:8080/pTfS5uXy`
**Method**: `GET`
**Description**: Redirects to the original URL based on the shortened URL.
**Response**: Redirects to the original URL.

### Top Shortened Domains

**Endpoint**: `http://localhost:8080/topdomains?count=3`  
**Method**: `POST`  
**Description**: Returns the top most frequently shortened domains.  

**Response Body**:
```json
[
    {
        "Key": "www.google.com",
        "Value": 3
    },
    {
        "Key": "www.yahoo.com",
        "Value": 2
    },
    {
        "Key": "www.gmail.com",
        "Value": 1
    }
]
```
