package subsonic

import "time"

// Used for any entity where only a name and an ID appear.
// Eg. OpenSubsonic artists list for an album.
type IDName struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type OpenSubsonicExtension struct {
	Name     string `xml:"name,attr"`
	Versions []int  `xml:"versions"`
}

type openSubsonicExtensions struct {
	OpenSubsonicExtensions []*OpenSubsonicExtension `xml:"openSubsonicExtension,omitempty"`
}

type LyricsList struct {
	StructuredLyrics []*StructuredLyrics `xml:"structuredLyrics,omitempty"`
}

type StructuredLyrics struct {
	DisplayArtist string      `xml:"displayArtist,attr"`
	DisplayTitle  string      `xml:"displayTitle,attr"`
	Lang          string      `xml:"lang,attr"`
	Offset        int         `xml:"offset,attr"`
	Synced        bool        `xml:"synced,attr"`
	Lines         []LyricLine `xml:"line,omitempty"`
}

type LyricLine struct {
	Start int    `xml:"start,attr"`
	Text  string `xml:",chardata"`
	// Navidrome 0.51.0 - 0.52.5 incorrecty returns the lyric line text here
	// This will be removed in the future
	Value string `xml:"value"`
}

type ItemDate struct {
	Year  *int `xml:"year,attr,omitempty"`
	Month *int `xml:"month,attr,omitempty"`
	Day   *int `xml:"day,attr,omitempty"`
}

type Contributor struct {
	Role   string `xml:"role,attr"`
	Artist IDName `xml:"artist"`
}

type ReplayGain struct {
	TrackGain float64 `xml:"trackGain,omitempty,attr"    json:"trackGain,omitempty"`
	AlbumGain float64 `xml:"albumGain,omitempty,attr"    json:"albumGain,omitempty"`
	TrackPeak float64 `xml:"trackPeak,omitempty,attr"    json:"trackPeak,omitempty"`
	AlbumPeak float64 `xml:"albumPeak,omitempty,attr"    json:"albumPeak,omitempty"`
}

type PlayQueueByIndex struct {
	Entries      []*Child  `xml:"entry,omitempty"`
	CurrentIndex int64     `xml:"currentIndex,attr,omitempty"`
	Position     int64     `xml:"position,attr,omitempty"`
	Username     string    `xml:"username,attr"`
	Changed      time.Time `xml:"changed,attr"`
	ChangedBy    string    `xml:"changedBy,attr"`
}
