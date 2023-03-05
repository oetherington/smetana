package smetana

import (
	"log"
	"os"
	"strings"
)

// Render a [Node] to an HTML string with the default settings.
// See [RenderHtmlOpts] for more fine-grained control.
func RenderHtml(node Node) string {
	return RenderHtmlOpts(node, false, nil)
}

// Render a [Node] to an HTML string specifying particular settings for the
// internal [Builder].
// See the [Builder] struct for the available configuration values.
// See [RenderHtml] for a simpler interface with default values.
func RenderHtmlOpts(
	node Node,
	deterministicAttrs bool,
	logger *log.Logger,
) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, deterministicAttrs, logger}
	node.ToHtml(&builder)
	return builder.Buf.String()
}

// Render a [StyleSheet] into a CSS string with the default settings.
// See [RenderCssOpts] for more fine-grained control.
func RenderCss(styles StyleSheet, palette Palette) string {
	return RenderCssOpts(styles, palette, nil)
}

// Render a [StyleSheet] into a CSS string specifying particular settings for
// the internal [Builder].
// See the [Builder] struct for the available configuration values.
// See [RenderCss] for a simpler interface with default values.
func RenderCssOpts(
	styles StyleSheet,
	palette Palette,
	logger *log.Logger,
) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, false, logger}
	styles.ToCss(&builder, palette)
	return builder.Buf.String()
}

// Render a [Sitemap] into an XML string with the default settings.
// See [RenderSitemapOpts] for more fine-grained control.
func RenderSitemap(sitemap Sitemap) string {
	return RenderSitemapOpts(sitemap, nil)
}

// Render a [Sitemap] into an XML string specifying particular settings for the
// internal [Builder].
// See the [Builder] struct for the available configuration values.
// See [RenderSitemap] for a simpler interface with default values.
func RenderSitemapOpts(sitemap Sitemap, logger *log.Logger) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, false, logger}
	sitemap.ToXml(&builder)
	return builder.Buf.String()
}
