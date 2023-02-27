package smetana

import (
	"testing"
	"time"
)

type ChangeFreqTestCase struct {
	changefreq ChangeFreq
	expected   string
}

func TestConvertChangeFreqToString(t *testing.T) {
	tests := []ChangeFreqTestCase{
		{ChangeFreqAlways, "always"},
		{ChangeFreqHourly, "hourly"},
		{ChangeFreqDaily, "daily"},
		{ChangeFreqWeekly, "weekly"},
		{ChangeFreqMonthly, "monthly"},
		{ChangeFreqYearly, "yearly"},
		{ChangeFreqNever, "never"},
		{ChangeFreqNone, ""},
	}
	for _, test := range tests {
		result := test.changefreq.String()
		assertEqual(t, test.expected, result)
	}
}

func TestCreateSitemapLocationWithUrl(t *testing.T) {
	url := "https://duckduckgo.com"
	loc := SitemapLocationUrl(url)
	assertEqual(t, url, loc.url)
	assertEqual(t, nil, loc.lastmod)
	assertEqual(t, ChangeFreqNone, loc.changefreq)
	assertEqual(t, 0.5, loc.priority)
}

func TestCreateSitemapLocationWithLastMod(t *testing.T) {
	url := "https://duckduckgo.com"
	lastmod, err := time.Parse(sitemapDateFormat, "2020-02-03T12:00:00+01:00")
	assertEqual(t, nil, err)
	loc := SitemapLocationMod(url, lastmod)
	assertEqual(t, url, loc.url)
	assertEqual(t, lastmod, *loc.lastmod)
	assertEqual(t, ChangeFreqNone, loc.changefreq)
	assertEqual(t, 0.5, loc.priority)
}

func TestCreateNewSitemapLocation(t *testing.T) {
	url := "https://duckduckgo.com"
	lastmod, err := time.Parse(sitemapDateFormat, "2022-02-03T12:00:00+01:00")
	assertEqual(t, nil, err)
	changefreq := ChangeFreqMonthly
	priority := 0.8
	loc := NewSitemapLocation(url, lastmod, changefreq, priority)
	assertEqual(t, url, loc.url)
	assertEqual(t, lastmod, *loc.lastmod)
	assertEqual(t, changefreq, loc.changefreq)
	assertEqual(t, priority, loc.priority)
}

func TestCanRenderSitemap(t *testing.T) {
	lastmod1, err := time.Parse(sitemapDateFormat, "2020-02-03T12:00:00+01:00")
	assertEqual(t, nil, err)
	lastmod2, err := time.Parse(sitemapDateFormat, "2022-02-03T12:00:00+01:00")
	assertEqual(t, nil, err)
	sitemap := Sitemap{
		SitemapLocationUrl("https://duckduckgo.com"),
		SitemapLocationMod("https://lobste.rs", lastmod1),
		NewSitemapLocation(
			"https://news.ycombinator.com",
			lastmod2,
			ChangeFreqAlways,
			0.9,
		),
	}
	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
		"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">" +
		"<url><loc>https://duckduckgo.com</loc></url>" +
		"<url><loc>https://lobste.rs</loc><lastmod>2020-02-03T12:00:00+01:00</lastmod></url>" +
		"<url><loc>https://news.ycombinator.com</loc><lastmod>2022-02-03T12:00:00+01:00</lastmod><changefreq>always</changefreq><priority>0.90</priority></url>" +
		"</urlset>"
	result := RenderSitemap(sitemap)
	assertEqual(t, expected, result)
}
