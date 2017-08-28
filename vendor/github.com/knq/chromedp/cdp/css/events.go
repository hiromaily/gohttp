package css

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventMediaQueryResultChanged fires whenever a MediaQuery result changes
// (for example, after a browser window has been resized.) The current
// implementation considers only viewport-dependent media features.
type EventMediaQueryResultChanged struct{}

// EventFontsUpdated fires whenever a web font gets loaded.
type EventFontsUpdated struct{}

// EventStyleSheetChanged fired whenever a stylesheet is changed as a result
// of the client operation.
type EventStyleSheetChanged struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

// EventStyleSheetAdded fired whenever an active document stylesheet is
// added.
type EventStyleSheetAdded struct {
	Header *StyleSheetHeader `json:"header"` // Added stylesheet metainfo.
}

// EventStyleSheetRemoved fired whenever an active document stylesheet is
// removed.
type EventStyleSheetRemoved struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // Identifier of the removed stylesheet.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventCSSMediaQueryResultChanged,
	cdp.EventCSSFontsUpdated,
	cdp.EventCSSStyleSheetChanged,
	cdp.EventCSSStyleSheetAdded,
	cdp.EventCSSStyleSheetRemoved,
}