package client_test

import (
	"regexp"
	"sync"
	"testing"

	"github.com/isucon10-qualify/isucon10-qualify/bench/client"
)

var botUserAgentRegExpList []*regexp.Regexp = []*regexp.Regexp{
	regexp.MustCompile(`Isuuglebot(-Mobile)?`),
	regexp.MustCompile(`Isuuglebot-Image\/`),
	regexp.MustCompile(`Mediapartners-Isuugle`),
	regexp.MustCompile(`IsuhooSlurp`),
	regexp.MustCompile(`IsuhooFeedSeeker(Beta)?Jp`),
	regexp.MustCompile(`crawler \(http:\/\/(listing\.isuhoo\.co\.jp\/support\/faq\/|help\.isuhoo\.co\.jp\/help\/jp\/)`),
	regexp.MustCompile(`isuingbot`),
	regexp.MustCompile(`Baisuspider`),
	regexp.MustCompile(`Baisuspider(-image)?\+`),
	regexp.MustCompile(`(?i)(bot|crawler|spider)(?:[-_ .\/;@()]|$)`),
}

func isBotUserAgent(ua string) bool {
	for _, botUserAgentRegExp := range botUserAgentRegExpList {
		if botUserAgentRegExp.MatchString(ua) {
			return true
		}
	}
	return false
}

func Test_UserAgent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ua := client.GenerateUserAgent()
			if isBotUserAgent(ua) {
				t.Errorf("Bot User Agent was generated by GenerateUserAgent func: %v", ua)
			}
		}()
	}

	wg.Wait()
}

func Test_BotUserAgent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ua := client.GenerateBotUserAgent()
			if !isBotUserAgent(ua) {
				t.Errorf("User Agent was generated by GenerateBotUserAgent func: %v", ua)
			}
		}()
	}

	wg.Wait()
}
