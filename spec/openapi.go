package spec

import (
	"fmt"
	"net/http"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
)

func ExposeSwaggerUI(e *echo.Echo, openapi *docs.OpenApi) {
	path := "/documentation"
	openAPIPath := fmt.Sprintf("%s.json", path)

	// Json
	e.GET(openAPIPath, func(c echo.Context) error { return c.JSON(http.StatusOK, openapi) })

	// Swagger UI
	e.GET(path, func(c echo.Context) error { return c.Redirect(http.StatusMovedPermanently, "/documentation/") })
	e.GET(fmt.Sprintf("%s/*", path), func(c echo.Context) error { return c.HTML(http.StatusOK, swaggeruiHTML(openAPIPath, "OpenAPI")) })
}

func swaggeruiHTML(openAPIPath string, title string) string {
	return fmt.Sprintf(`<!doctype html>
		<html lang="en">
		<head>
			<meta charset="utf-8" />
			<meta name="referrer" content="same-origin" />
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
			<title>%s</title>
			<!-- Embed elements Elements via Web Component -->
			<link href="https://unpkg.com/@stoplight/elements@9.0.0/styles.min.css" rel="stylesheet" />
			<script src="https://unpkg.com/@stoplight/elements@9.0.0/web-components.min.js" integrity="sha256-Tqvw1qE2abI+G6dPQBc5zbeHqfVwGoamETU3/TSpUw4="
					crossorigin="anonymous"></script>
		</head>
		<body style="height: 100vh;">

			<elements-api
			apiDescriptionUrl="%s"
			router="hash"
			layout="responsive"
			tryItCredentialsPolicy="same-origin"
			hideSchemas="true"
			/>

		</body>
		</html>`, title, openAPIPath,
	)
}
