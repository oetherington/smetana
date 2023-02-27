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
	lastmod := time.Unix(1243744874, 0)
	loc := SitemapLocationMod(url, lastmod)
	assertEqual(t, url, loc.url)
	assertEqual(t, lastmod, *loc.lastmod)
	assertEqual(t, ChangeFreqNone, loc.changefreq)
	assertEqual(t, 0.5, loc.priority)
}

func TestCreateNewSitemapLocation(t *testing.T) {
	url := "https://duckduckgo.com"
	lastmod := time.Unix(1243944874, 0)
	changefreq := ChangeFreqMonthly
	priority := 0.8
	loc := NewSitemapLocation(url, lastmod, changefreq, priority)
	assertEqual(t, url, loc.url)
	assertEqual(t, lastmod, *loc.lastmod)
	assertEqual(t, changefreq, loc.changefreq)
	assertEqual(t, priority, loc.priority)
}

func TestCanRenderSitemap(t *testing.T) {
	sitemap := Sitemap{
		SitemapLocationUrl("https://duckduckgo.com"),
		SitemapLocationMod("https://lobste.rs", time.Unix(1243744874, 0)),
		NewSitemapLocation(
			"https://news.ycombinator.com",
			time.Unix(1243944874, 0),
			ChangeFreqAlways,
			0.9,
		),
	}
	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
		"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">" +
		"<url><loc>https://duckduckgo.com</loc></url>" +
		"<url><loc>https://lobste.rs</loc><lastmod>2009-05-31T05:41:14+01:00</lastmod></url>" +
		"<url><loc>https://news.ycombinator.com</loc><lastmod>2009-06-02T13:14:34+01:00</lastmod><changefreq>always</changefreq><priority>0.90</priority></url>" +
		"</urlset>"
	result := RenderSitemap(sitemap)
	assertEqual(t, expected, result)
}
