package smetana

import (
	"fmt"
	"time"
)

// A valid [Sitemap] "change frequency"
type ChangeFreq int

const (
	// A [Sitemap] entry with no specified change frequency
	ChangeFreqNone ChangeFreq = iota
	// A [Sitemap] entry with change frequency "always"
	ChangeFreqAlways
	// A [Sitemap] entry with change frequency "hourly"
	ChangeFreqHourly
	// A [Sitemap] entry with change frequency "daily"
	ChangeFreqDaily
	// A [Sitemap] entry with change frequency "weekly"
	ChangeFreqWeekly
	// A [Sitemap] entry with change frequency "monthly"
	ChangeFreqMonthly
	// A [Sitemap] entry with change frequency "yearly"
	ChangeFreqYearly
	// A [Sitemap] entry with change frequency "never"
	ChangeFreqNever
)

// Convert a [Sitemap] change frequency to a string
func (freq ChangeFreq) String() string {
	switch freq {
	case ChangeFreqAlways:
		return "always"
	case ChangeFreqHourly:
		return "hourly"
	case ChangeFreqDaily:
		return "daily"
	case ChangeFreqWeekly:
		return "weekly"
	case ChangeFreqMonthly:
		return "monthly"
	case ChangeFreqYearly:
		return "yearly"
	case ChangeFreqNever:
		return "never"
	}
	return ""
}

// Represents a single entry in a [Sitemap] with a URL string and an optional
// last modified date. The date can be any type implementing [fmt.Stringer],
// but is most commonly a string or a [time.Time].
type SitemapLocation struct {
	url        string
	lastmod    *time.Time
	changefreq ChangeFreq
	priority   float64
}

// [Sitemap] represents an XML sitemap according to the schema at
// https://www.sitemaps.org/protocol.html
// Convert to an XML string with the [ToXml] method.
type Sitemap []SitemapLocation

const sitemapDateFormat = "2006-01-02T15:04:05Z07:00"

// Convert a [Sitemap] to an XML string.
func (sitemap Sitemap) ToXml(builder *Builder) {
	builder.Buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	builder.Buf.WriteString("<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">")
	for _, loc := range sitemap {
		builder.Buf.WriteString("<url><loc>")
		builder.Buf.WriteString(loc.url)
		builder.Buf.WriteString("</loc>")
		if loc.lastmod != nil {
			builder.Buf.WriteString("<lastmod>")
			date := loc.lastmod.Format(sitemapDateFormat)
			builder.Buf.WriteString(date)
			builder.Buf.WriteString("</lastmod>")
		}
		if loc.changefreq != ChangeFreqNone {
			builder.Buf.WriteString("<changefreq>")
			builder.Buf.WriteString(loc.changefreq.String())
			builder.Buf.WriteString("</changefreq>")
		}
		if loc.priority != 0.5 {
			builder.Buf.WriteString("<priority>")
			builder.Buf.WriteString(fmt.Sprintf("%.2f", loc.priority))
			builder.Buf.WriteString("</priority>")
		}
		builder.Buf.WriteString("</url>")
	}
	builder.Buf.WriteString("</urlset>")
}

// Create a [SitemapLocation] with just a URL.
func SitemapLocationUrl(url string) SitemapLocation {
	return SitemapLocation{url, nil, ChangeFreqNone, 0.5}
}

// Create a [SitemapLocation] with just a URL and last modified date.
func SitemapLocationMod(url string, lastmod time.Time) SitemapLocation {
	return SitemapLocation{url, &lastmod, ChangeFreqNone, 0.5}
}

// Create a [SitemapLocation] with all available parameters: url, lastmod,
// changefreq and priority.
func NewSitemapLocation(
	url string,
	lastmod time.Time,
	changefreq ChangeFreq,
	priority float64,
) SitemapLocation {
	return SitemapLocation{url, &lastmod, changefreq, priority}
}
