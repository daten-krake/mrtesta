package main

import (
	"log"
	"os"
	"time"

	"mrtesta/internal/config"
	"mrtesta/internal/github"
)

const (
	githubOwner  = "daten-krake"
	githubRepo   = "mrtesta"
	githubBranch = "main"
	configPath   = "config/config.json"
	pollInterval = 1 * time.Hour
)

func tokenFromEnv() string {
	return os.Getenv("GITHUB_TOKEN")
}

func fetchAndApply(token string) *config.Config {
	data, err := github.FetchRaw(githubOwner, githubRepo, githubBranch, configPath, token)
	if err != nil {
		log.Printf("config fetch failed: %v", err)
		return nil
	}

	cfg, err := config.Parse(data)
	if err != nil {
		log.Printf("config parse failed: %v", err)
		return nil
	}

	log.Printf("config loaded: %d enabled test(s)", len(cfg.EnabledTests))
	for i, name := range cfg.EnabledTests {
		log.Printf("  [%d] %s", i+1, name)
	}
	return cfg
}

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	token := tokenFromEnv()
	log.Printf("mrtesta agent starting (token present: %v)", token != "")

	fetchAndApply(token)

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for range ticker.C {
		fetchAndApply(token)
	}
}