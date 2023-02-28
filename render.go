package smetana

import (
	"log"
	"os"
	"strings"
)

// Render a [Node] to an HTML string with the default settings. See
// [RenderHtmlOpts] for more fine-grained control.
func RenderHtml(node Node) string {
	return RenderHtmlOpts(node, false, nil)
}

// Render a [Node] to an HTML string specifying particular settings for the
// internal [Builder]. See the [Builder] struct for the available
// configuration values.
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

// Render a [StyleSheet] into a CSS string.
func RenderCss(styles StyleSheet) string {
	return RenderCssOpts(styles, nil)
}

// Render a [StyleSheet] into a CSS string specifying particular settings for
// the internal [Builder]. See the [Builder] struct for the available
// configuration values.
func RenderCssOpts(styles StyleSheet, logger *log.Logger) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, false, logger}
	styles.Compile(&builder)
	return builder.Buf.String()
}

// Render a [Sitemap] into an XML string.
func RenderSitemap(sitemap Sitemap) string {
	return RenderSitemapOpts(sitemap, nil)
}

// Render a [Sitemap] into an XML string specifying particular settings for the
// internal [Builder]. See the [Builder] struct for the available
// configuration values.
func RenderSitemapOpts(sitemap Sitemap, logger *log.Logger) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, false, logger}
	sitemap.ToXml(&builder)
	return builder.Buf.String()
}
