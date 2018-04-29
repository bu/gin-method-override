# Gin MethodOverride Middleware

![Go Report Card](https://goreportcard.com/badge/github.com/bu/gin-method-override)

A Gin web framework middleware for method override by POST form param _method, inspired by Ruby's same name rack

## Usage

### Server-side

```go

package main

import (
	gin "github.com/gin-gonic/gin"
	method "github.com/bu/gin-method-override"
)

func main() {
	// create a Gin engine
	r := gin.Default()

	// our middle-ware
	r.Use(ProcessMethodOverride(r))

	// routes
	r.PUT("/test", func (c *gin.Context) {
		c.String(200, "1")
	})

	// listen to request
	r.Run(":8080")
}

```

## Client side

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
</head>
<body>
    <form action="/test">
        <input type="hidden" name="_method" value="PUT">
        <input type="text" name="testing" value="1">
        <button type="submit">Send</button>
    </form>
</body>
</html>
```