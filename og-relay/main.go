package main

import (
	"context"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var pageTemplate = template.Must(template.New("page").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>

    <!-- Open Graph / Matrix / Signal -->
    <meta property="og:type" content="website">
    <meta property="og:url" content="{{.FullURL}}">
    <meta property="og:title" content="{{.Title}}">
    <meta property="og:description" content="Access {{.Title}} on {{.Domain}}">
    <meta property="og:image" content="{{.ImageURL}}">
    <meta property="og:site_name" content="{{.Domain}}">

    <!-- Twitter Card -->
    <meta name="twitter:card" content="summary_large_image">
    <meta name="twitter:url" content="{{.FullURL}}">
    <meta name="twitter:title" content="{{.Title}}">
    <meta name="twitter:description" content="Access {{.Title}} on {{.Domain}}">
    <meta name="twitter:image" content="{{.ImageURL}}">
</head>
<body>
    <p>{{.Title}} Link Preview</p>
</body>
</html>`))

var svgTemplate = template.Must(template.New("svg").Parse(`<svg xmlns="http://www.w3.org/2000/svg" width="1200" height="630" viewBox="0 0 1200 630">
  <defs>
    <linearGradient id="bg" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" stop-color="#0f172a"/>
      <stop offset="50%" stop-color="#1e293b"/>
      <stop offset="100%" stop-color="#0f172a"/>
    </linearGradient>
    <linearGradient id="accent" x1="0%" y1="0%" x2="100%" y2="0%">
      <stop offset="0%" stop-color="#38bdf8"/>
      <stop offset="100%" stop-color="#818cf8"/>
    </linearGradient>
    <filter id="shadow" x="-10%" y="-10%" width="120%" height="120%">
      <feDropShadow dx="0" dy="10" stdDeviation="20" flood-color="#000" flood-opacity="0.5"/>
    </filter>
  </defs>

  <rect width="1200" height="630" fill="url(#bg)"/>
  
  <g opacity="0.05" stroke="#ffffff" stroke-width="1">
    <path d="M 0 100 L 1200 100 M 0 200 L 1200 200 M 0 300 L 1200 300 M 0 400 L 1200 400 M 0 500 L 1200 500" />
    <path d="M 200 0 L 200 630 M 400 0 L 400 630 M 600 0 L 600 630 M 800 0 L 800 630 M 1000 0 L 1000 630" />
  </g>

  <rect x="0" y="0" width="1200" height="8" fill="url(#accent)"/>

  <g transform="translate(100, 140)" filter="url(#shadow)">
    <rect x="0" y="0" width="80" height="80" rx="20" fill="url(#accent)"/>
    <path d="M 25 40 L 35 50 L 55 30" fill="none" stroke="#ffffff" stroke-width="6" stroke-linecap="round" stroke-linejoin="round"/>

    <text x="110" y="55" font-family="system-ui, -apple-system, sans-serif" font-size="54" font-weight="700" fill="#f8fafc">
      {{.Title}}
    </text>

    <text x="0" y="180" font-family="system-ui, -apple-system, sans-serif" font-size="32" font-weight="400" fill="#94a3b8">
      {{.Domain}}
    </text>

    <text x="0" y="240" font-family="system-ui, -apple-system, sans-serif" font-size="24" font-weight="400" fill="#64748b">
      Protected Service • {{.Domain}}
    </text>
  </g>

  <text x="100" y="550" font-family="system-ui, -apple-system, sans-serif" font-size="20" font-weight="500" fill="#475569">
    Secure Access Link
  </text>
</svg>`))

type PageData struct {
	Title    string
	Domain   string
	FullURL  string
	ImageURL string
}

func formatTitle(host string) string {
	parts := strings.Split(host, ".")
	mainPart := parts[0]
	if mainPart == "" {
		mainPart = "Service"
	}

	var sb strings.Builder
	for i, r := range mainPart {
		if i > 0 && r >= 'A' && r <= 'Z' && mainPart[i-1] >= 'a' && mainPart[i-1] <= 'z' {
			sb.WriteRune(' ')
		}
		sb.WriteRune(r)
	}
	cleaned := strings.ReplaceAll(sb.String(), "-", " ")
	cleaned = strings.ReplaceAll(cleaned, "_", " ")

	words := strings.Fields(cleaned)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'unsafe-inline'; img-src 'self'")

		host := r.Header.Get("X-Forwarded-Host")
		if host == "" {
			host = r.Host
		}
		if host == "" {
			host = "Service"
		}

		ua := r.Header.Get("User-Agent")
		log.Printf("[OG-Relay] Host: %s | Path: %s | User-Agent: %s", host, r.URL.Path, ua)

		uri := r.Header.Get("X-Forwarded-Uri")
		if uri == "" {
			uri = r.URL.RequestURI()
		}

		proto := r.Header.Get("X-Forwarded-Proto")
		if proto == "" {
			proto = "https"
		}

		title := formatTitle(host)
		domain := host
		fullURL := fmt.Sprintf("%s://%s%s", proto, host, uri)
		imageURL := fmt.Sprintf("%s://%s/__og_relay__/image.svg", proto, host)

		if r.URL.Path == "/__og_relay__/image.svg" || r.URL.Path == "/__og_relay__/image.png" {
			w.Header().Set("Content-Type", "image/svg+xml")
			w.Header().Set("Cache-Control", "public, max-age=86400")
			data := PageData{Title: html.EscapeString(title), Domain: html.EscapeString(domain)}
			_ = svgTemplate.Execute(w, data)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		data := PageData{
			Title:    title,
			Domain:   domain,
			FullURL:  fullURL,
			ImageURL: imageURL,
		}
		_ = pageTemplate.Execute(w, data)
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("OG-Relay (Go) listening on :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down OG-Relay...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}

	log.Println("OG-Relay stopped cleanly.")
}
