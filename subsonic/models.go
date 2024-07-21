package subsonic

/* This file was automatically generated from the xsd schema provided by Subsonic, then manually modified.
 *   http://www.subsonic.org/pages/inc/api/schema/subsonic-rest-api-1.16.1.xsd
 *   xsdgen -o xml.go -pkg subsonic -ns "http://subsonic.org/restapi" subsonic-rest-api-1.16.1.xsd
 * Changes from the original include:
 * - Adding missing name (value of xml element) for each genre
 * - Capitalize "ID" in struct names and add missing ID fields.
 * - Merge *With* variants of structs.
 */

import (
	"bytes"
	"encoding/xml"
	"time"
)

// AlbumID3 is an album that's organized by music file tags.
type AlbumID3 struct {
	ID                  string    `xml:"id,attr"`        // Manually added
	Song                []*Child  `xml:"song,omitempty"` // Merged from AlbumWithSongsID3
	Name                string    `xml:"name,attr"`
	Artist              string    `xml:"artist,attr,omitempty"`
	ArtistID            string    `xml:"artistId,attr,omitempty"`
	Artists             []IDName  `xml:"artists,omitempty"` // OpenSubsonic extension
	CoverArt            string    `xml:"coverArt,attr,omitempty"`
	SongCount           int       `xml:"songCount,attr"`
	Duration            int       `xml:"duration,attr"`
	PlayCount           int64     `xml:"playCount,attr,omitempty"`
	Created             time.Time `xml:"created,attr"`
	Starred             time.Time `xml:"starred,attr,omitempty"`
	Year                int       `xml:"year,attr,omitempty"`
	ReleaseDate         *ItemDate `xml:"releaseDate,omitempty"`         // OpenSubsonic extension
	OriginalReleaseDate *ItemDate `xml:"originalReleaseDate,omitempty"` // OpenSubsonic extension
	Genre               string    `xml:"genre,attr,omitempty"`
	Genres              []IDName  `xml:"genres,omitempty"`       // OpenSubsonic extension
	ReleaseTypes        []string  `xml:"releaseTypes,omitempty"` // OpenSubsonic extension
	IsCompilation       bool      `xml:"isCompilation,attr"`     // OpenSubsonic extension
}

func (t *AlbumID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AlbumID3
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *AlbumID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AlbumID3
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

// AlbumInfo is a collection of notes and links describing an album.
type AlbumInfo struct {
	Notes          string `xml:"notes,omitempty"`
	MusicBrainzID  string `xml:"musicBrainzId,omitempty"`
	LastFmUrl      string `xml:"lastFmUrl,omitempty"`
	SmallImageUrl  string `xml:"smallImageUrl,omitempty"`
	MediumImageUrl string `xml:"mediumImageUrl,omitempty"`
	LargeImageUrl  string `xml:"largeImageUrl,omitempty"`
}

type albumList struct {
	Album []*Child `xml:"album,omitempty"`
}

type albumList2 struct {
	Album []*AlbumID3 `xml:"album,omitempty"`
}

// Artist is an artist from the server, organized in the folders pattern.
type Artist struct {
	ID             string    `xml:"id,attr"`
	Name           string    `xml:"name,attr"`
	ArtistImageUrl string    `xml:"artistImageUrl,attr,omitempty"`
	Starred        time.Time `xml:"starred,attr,omitempty"`
	UserRating     int       `xml:"userRating,attr,omitempty"`
	AverageRating  float64   `xml:"averageRating,attr,omitempty"`
}

func (t *Artist) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Artist
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Artist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Artist
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

// ArtistID3 is an artist from the server, organized by ID3 tag.
type ArtistID3 struct {
	ID             string      `xml:"id,attr"`         // Manually added
	Album          []*AlbumID3 `xml:"album,omitempty"` // Merged with ArtistWithAlbumsID3
	Name           string      `xml:"name,attr"`
	CoverArt       string      `xml:"coverArt,attr,omitempty"`
	ArtistImageUrl string      `xml:"artistImageUrl,attr,omitempty"`
	AlbumCount     int         `xml:"albumCount,attr"`
	Starred        time.Time   `xml:"starred,attr,omitempty"`
}

func (t *ArtistID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ArtistID3
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *ArtistID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtistID3
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo.
type ArtistInfo struct {
	SimilarArtist  []*Artist `xml:"similarArtist,omitempty"`
	Biography      string    `xml:"biography,omitempty"`
	MusicBrainzID  string    `xml:"musicBrainzId,omitempty"`
	LastFmUrl      string    `xml:"lastFmUrl,omitempty"`
	SmallImageUrl  string    `xml:"smallImageUrl,omitempty"`
	MediumImageUrl string    `xml:"mediumImageUrl,omitempty"`
	LargeImageUrl  string    `xml:"largeImageUrl,omitempty"`
}

// ArtistInfo2 is all auxillary information about an artist from GetArtistInfo2, with similar artists organized by ID3 tags.
type ArtistInfo2 struct {
	SimilarArtist  []*ArtistID3 `xml:"similarArtist,omitempty"`
	Biography      string       `xml:"biography,omitempty"`
	MusicBrainzID  string       `xml:"musicBrainzId,omitempty"`
	LastFmUrl      string       `xml:"lastFmUrl,omitempty"`
	SmallImageUrl  string       `xml:"smallImageUrl,omitempty"`
	MediumImageUrl string       `xml:"mediumImageUrl,omitempty"`
	LargeImageUrl  string       `xml:"largeImageUrl,omitempty"`
}

// ArtistsID3 is an index of every artist on the server organized by ID3 tag, from getArtists.
type ArtistsID3 struct {
	Index           []*IndexID3 `xml:"index,omitempty"`
	IgnoredArticles string      `xml:"ignoredArticles,attr"`
}

type Bookmark struct {
	Entry    *Child    `xml:"entry"`
	Position int64     `xml:"position,attr"`
	Username string    `xml:"username,attr"`
	Comment  string    `xml:"comment,attr,omitempty"`
	Created  time.Time `xml:"created,attr"`
	Changed  time.Time `xml:"changed,attr"`
}

func (t *Bookmark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Bookmark
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *Bookmark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Bookmark
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}

type bookmarks struct {
	Bookmark []*Bookmark `xml:"bookmark,omitempty"`
}

type ChatMessage struct {
	Username string `xml:"username,attr"`
	Time     int64  `xml:"time,attr"`
	Message  string `xml:"message,attr"`
}

type chatMessages struct {
	ChatMessage []*ChatMessage `xml:"chatMessage,omitempty"`
}

// Child is a song, or a generic entry in the hierarchical directory structure of the database.
// You can tell if Child is used as a song contextually based on what it was returned by, or if the IsDir boolean was set to true.
type Child struct {
	ID                    string        `xml:"id,attr"` // Manually added
	Parent                string        `xml:"parent,attr,omitempty"`
	IsDir                 bool          `xml:"isDir,attr"`
	Title                 string        `xml:"title,attr"`
	Album                 string        `xml:"album,attr,omitempty"`
	Artist                string        `xml:"artist,attr,omitempty"`
	Artists               []IDName      `xml:"artists,omitempty"`                 // OpenSubsonic extension
	AlbumArtists          []IDName      `xml:"albumArtists,omitempty"`            // OpenSubsonic extension
	DisplayArtist         string        `xml:"displayArtist,attr,omitempty"`      // OpenSubsonic extension
	DisplayAlbumArtist    string        `xml:"displayAlbumArtist,attr,omitempty"` // OpenSubsonic extension
	Contributors          []Contributor `xml:"contributors,omitempty"`            // OpenSubsonic extension
	DisplayComposer       string        `xml:"displayComposer,attr,omitempty"`    // OpenSubsonic extension
	Track                 int           `xml:"track,attr,omitempty"`
	Year                  int           `xml:"year,attr,omitempty"`
	Genre                 string        `xml:"genre,attr,omitempty"`
	Genres                []IDName      `xml:"genres,omitempty"`             // OpenSubsonic extension
	Comment               string        `xml:"comment,attr,omitempty"`       // OpenSubsonic extension
	BPM                   int           `xml:"bpm,attr"`                     // OpenSubsonic extension
	MusicBrainzID         string        `xml:"musicBrainzId,attr,omitempty"` // OpenSubsonic extension
	CoverArt              string        `xml:"coverArt,attr,omitempty"`
	Size                  int64         `xml:"size,attr,omitempty"`
	ContentType           string        `xml:"contentType,attr,omitempty"`
	Suffix                string        `xml:"suffix,attr,omitempty"`
	TranscodedContentType string        `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string        `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int           `xml:"duration,attr,omitempty"`
	BitRate               int           `xml:"bitRate,attr,omitempty"`
	BitDepth              int           `xml:"bitDepth,attr,omitempty"`     // OpenSubsonic extension
	SamplingRate          int           `xml:"samplingRate,attr,omitempty"` // OpenSubsonic extension
	ChannelCount          int           `xml:"channelCount,attr,omitempty"` // OpenSubsonic extension
	Path                  string        `xml:"path,attr,omitempty"`
	IsVideo               bool          `xml:"isVideo,attr,omitempty"`
	UserRating            int           `xml:"userRating,attr,omitempty"`
	AverageRating         float64       `xml:"averageRating,attr,omitempty"`
	PlayCount             int64         `xml:"playCount,attr,omitempty"`
	DiscNumber            int           `xml:"discNumber,attr,omitempty"`
	Created               time.Time     `xml:"created,attr,omitempty"`
	Starred               time.Time     `xml:"starred,attr,omitempty"`
	Played                time.Time     `xml:"played,attr,omitempty"` // OpenSubsonic extension
	AlbumID               string        `xml:"albumId,attr,omitempty"`
	ArtistID              string        `xml:"artistId,attr,omitempty"`
	Type                  string        `xml:"type,attr,omitempty"` // May be one of music, podcast, audiobook, video
	BookmarkPosition      int64         `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int           `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int           `xml:"originalHeight,attr,omitempty"`
	ReplayGain            *ReplayGain   `xml:"replayGain,omitempty"` // OpenSubsonic extension
}

func (t *Child) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Child
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
		Played  *xsdDateTime `xml:"played,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	layout.Played = (*xsdDateTime)(&layout.T.Played)
	return e.EncodeElement(layout, start)
}
func (t *Child) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Child
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
		Played  *xsdDateTime `xml:"played,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	overlay.Played = (*xsdDateTime)(&overlay.T.Played)
	return d.DecodeElement(&overlay, &start)
}

// Directory is an entry in the hierarchical folder structure organization of the server database.
type Directory struct {
	ID            string    `xml:"id,attr"` // Manually added
	Child         []*Child  `xml:"child,omitempty"`
	Parent        string    `xml:"parent,attr,omitempty"`
	Name          string    `xml:"name,attr"`
	Starred       time.Time `xml:"starred,attr,omitempty"`
	UserRating    int       `xml:"userRating,attr,omitempty"`
	AverageRating float64   `xml:"averageRating,attr,omitempty"`
	PlayCount     int64     `xml:"playCount,attr,omitempty"`
}

func (t *Directory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Directory
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Directory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Directory
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

type Error struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:"message,attr,omitempty"`
}

// Genre is a style tag for a collection of songs and albums.
type Genre struct {
	Name       string `xml:",chardata"` // Added manually
	SongCount  int    `xml:"songCount,attr"`
	AlbumCount int    `xml:"albumCount,attr"`
}

type genres struct {
	Genre []*Genre `xml:"genre,omitempty"`
}

// Index is a collection of artists that begin with the same first letter, along with that letter or category.
type Index struct {
	Artist []*Artist `xml:"artist,omitempty"`
	Name   string    `xml:"name,attr"`
}

// Index is a collection of artists by ID3 tag that begin with the same first letter, along with that letter or category.
type IndexID3 struct {
	Artist []*ArtistID3 `xml:"artist,omitempty"`
	Name   string       `xml:"name,attr"`
}

// Indexes is the full index of the database, returned by getIndex.
// It contains some Index structs for each letter of the DB, plus Child entries for individual tracks.
type Indexes struct {
	Shortcut        []*Artist `xml:"shortcut,omitempty"`
	Index           []*Index  `xml:"index,omitempty"`
	Child           []*Child  `xml:"child,omitempty"`
	LastModified    int64     `xml:"lastModified,attr"`
	IgnoredArticles string    `xml:"ignoredArticles,attr"`
}

type InternetRadioStation struct {
	Name        string `xml:"name,attr"`
	StreamUrl   string `xml:"streamUrl,attr"`
	HomePageUrl string `xml:"homePageUrl,attr,omitempty"`
}

type internetRadioStations struct {
	InternetRadioStation []*InternetRadioStation `xml:"internetRadioStation,omitempty"`
}

type JukeboxPlaylist struct {
	Entry        []*Child `xml:"entry,omitempty"`
	CurrentIndex int      `xml:"currentIndex,attr"`
	Playing      bool     `xml:"playing,attr"`
	Gain         float32  `xml:"gain,attr"`
	Position     int      `xml:"position,attr,omitempty"`
}

type JukeboxStatus struct {
	CurrentIndex int     `xml:"currentIndex,attr"`
	Playing      bool    `xml:"playing,attr"`
	Gain         float32 `xml:"gain,attr"`
	Position     int     `xml:"position,attr,omitempty"`
}

// License contains information about the Subsonic server's license validity and contact information in the case of a trial subscription.
type License struct {
	Valid          bool      `xml:"valid,attr"`
	Email          string    `xml:"email,attr,omitempty"`
	LicenseExpires time.Time `xml:"licenseExpires,attr,omitempty"`
	TrialExpires   time.Time `xml:"trialExpires,attr,omitempty"`
}

func (t *License) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T License
	var layout struct {
		*T
		LicenseExpires *xsdDateTime `xml:"licenseExpires,attr,omitempty"`
		TrialExpires   *xsdDateTime `xml:"trialExpires,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LicenseExpires = (*xsdDateTime)(&layout.T.LicenseExpires)
	layout.TrialExpires = (*xsdDateTime)(&layout.T.TrialExpires)
	return e.EncodeElement(layout, start)
}
func (t *License) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T License
	var overlay struct {
		*T
		LicenseExpires *xsdDateTime `xml:"licenseExpires,attr,omitempty"`
		TrialExpires   *xsdDateTime `xml:"trialExpires,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LicenseExpires = (*xsdDateTime)(&overlay.T.LicenseExpires)
	overlay.TrialExpires = (*xsdDateTime)(&overlay.T.TrialExpires)
	return d.DecodeElement(&overlay, &start)
}

type Lyrics struct {
	Artist string `xml:"artist,attr,omitempty"`
	Title  string `xml:"title,attr,omitempty"`
	Text   string `xml:",chardata"`
}

// MusicFolder is a representation of a source of music files added to the server. It is identified primarily by the numeric ID.
type MusicFolder struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr,omitempty"`
}

type musicFolders struct {
	MusicFolder []*MusicFolder `xml:"musicFolder,omitempty"`
}

type newestPodcasts struct {
	Episode []*PodcastEpisode `xml:"episode,omitempty"`
}

type nowPlaying struct {
	Entry []*NowPlayingEntry `xml:"entry,omitempty"`
}

// NowPlayingEntry is one individual stream coming from the server along with information about who was streaming it.
type NowPlayingEntry struct {
	Username              string    `xml:"username,attr"`
	MinutesAgo            int       `xml:"minutesAgo,attr"`
	PlayerID              int       `xml:"playerId,attr"`
	PlayerName            string    `xml:"playerName,attr,omitempty"`
	Parent                string    `xml:"parent,attr,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"`
	Title                 string    `xml:"title,attr"`
	Album                 string    `xml:"album,attr,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"` // May be one of music, podcast, audiobook, video
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"`
}

func (t *NowPlayingEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T NowPlayingEntry
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *NowPlayingEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NowPlayingEntry
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

type PlayQueue struct {
	Entries   []*Child  `xml:"entry,omitempty"`
	Current   string    `xml:"current,attr,omitempty"`
	Position  int64     `xml:"position,attr,omitempty"`
	Username  string    `xml:"username,attr"`
	Changed   time.Time `xml:"changed,attr"`
	ChangedBy string    `xml:"changedBy,attr"`
}

func (t *PlayQueue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PlayQueue
	var layout struct {
		*T
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	layout.T = (*T)(t)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *PlayQueue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PlayQueue
	var overlay struct {
		*T
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}

// Playlist is a collection of songs with metadata like a name, comment, and information about the total duration of the playlist.
type Playlist struct {
	ID          string    `xml:"id,attr"`         // Added manually
	Entry       []*Child  `xml:"entry,omitempty"` // Merged from PlaylistWithSongs
	AllowedUser []string  `xml:"allowedUser,omitempty"`
	Name        string    `xml:"name,attr"`
	Comment     string    `xml:"comment,attr,omitempty"`
	Owner       string    `xml:"owner,attr,omitempty"`
	Public      bool      `xml:"public,attr,omitempty"`
	SongCount   int       `xml:"songCount,attr"`
	Duration    int       `xml:"duration,attr"`
	Created     time.Time `xml:"created,attr"`
	Changed     time.Time `xml:"changed,attr,omitempty"`
	CoverArt    string    `xml:"coverArt,attr,omitempty"`
}

func (t *Playlist) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Playlist
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *Playlist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Playlist
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}

type playlists struct {
	Playlist []*Playlist `xml:"playlist,omitempty"`
}

type PodcastChannel struct {
	Episode          []*PodcastEpisode `xml:"episode,omitempty"`
	Url              string            `xml:"url,attr"`
	Title            string            `xml:"title,attr,omitempty"`
	Description      string            `xml:"description,attr,omitempty"`
	CoverArt         string            `xml:"coverArt,attr,omitempty"`
	OriginalImageUrl string            `xml:"originalImageUrl,attr,omitempty"`
	Status           string            `xml:"status,attr"` // May be one of new, downloading, completed, error, deleted, skipped
	ErrorMessage     string            `xml:"errorMessage,attr,omitempty"`
}

type PodcastEpisode struct {
	StreamID              string    `xml:"streamId,attr,omitempty"`
	ChannelID             string    `xml:"channelId,attr"`
	Description           string    `xml:"description,attr,omitempty"`
	Status                string    `xml:"status,attr"` // May be one of new, downloading, completed, error, deleted, skipped
	PublishDate           time.Time `xml:"publishDate,attr,omitempty"`
	Parent                string    `xml:"parent,attr,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"`
	Title                 string    `xml:"title,attr"`
	Album                 string    `xml:"album,attr,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"` // May be one of music, podcast, audiobook, video
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"`
}

func (t *PodcastEpisode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PodcastEpisode
	var layout struct {
		*T
		PublishDate *xsdDateTime `xml:"publishDate,attr,omitempty"`
		Created     *xsdDateTime `xml:"created,attr,omitempty"`
		Starred     *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PublishDate = (*xsdDateTime)(&layout.T.PublishDate)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *PodcastEpisode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PodcastEpisode
	var overlay struct {
		*T
		PublishDate *xsdDateTime `xml:"publishDate,attr,omitempty"`
		Created     *xsdDateTime `xml:"created,attr,omitempty"`
		Starred     *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PublishDate = (*xsdDateTime)(&overlay.T.PublishDate)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

type podcasts struct {
	Channel []*PodcastChannel `xml:"channel,omitempty"`
}

// Response is the main target for unmarshalling data from the API - everything within the "subsonic-response" key
type Response struct {
	License                *License                 `xml:"license"`
	MusicFolders           *musicFolders            `xml:"musicFolders"`
	Indexes                *Indexes                 `xml:"indexes"`
	Directory              *Directory               `xml:"directory"`
	Genres                 *genres                  `xml:"genres"`
	Artists                *ArtistsID3              `xml:"artists"`
	Artist                 *ArtistID3               `xml:"artist"`
	Album                  *AlbumID3                `xml:"album"`
	Song                   *Child                   `xml:"song"`
	NowPlaying             *nowPlaying              `xml:"nowPlaying"`
	SearchResult2          *SearchResult2           `xml:"searchResult2"`
	SearchResult3          *SearchResult3           `xml:"searchResult3"`
	Playlists              *playlists               `xml:"playlists"`
	Playlist               *Playlist                `xml:"playlist"`
	JukeboxStatus          *JukeboxStatus           `xml:"jukeboxStatus"`
	JukeboxPlaylist        *JukeboxPlaylist         `xml:"jukeboxPlaylist"`
	Users                  *users                   `xml:"users"`
	User                   *User                    `xml:"user"`
	ChatMessages           *chatMessages            `xml:"chatMessages"`
	AlbumList              *albumList               `xml:"albumList"`
	AlbumList2             *albumList2              `xml:"albumList2"`
	RandomSongs            *songs                   `xml:"randomSongs"`
	SongsByGenre           *songs                   `xml:"songsByGenre"`
	Lyrics                 *Lyrics                  `xml:"lyrics"`
	Podcasts               *podcasts                `xml:"podcasts"`
	NewestPodcasts         *newestPodcasts          `xml:"newestPodcasts"`
	InternetRadioStations  *internetRadioStations   `xml:"internetRadioStations"`
	Bookmarks              *bookmarks               `xml:"bookmarks"`
	PlayQueue              *PlayQueue               `xml:"playQueue"`
	Shares                 *shares                  `xml:"shares"`
	Starred                *Starred                 `xml:"starred"`
	Starred2               *Starred2                `xml:"starred2"`
	AlbumInfo              *AlbumInfo               `xml:"albumInfo"`
	ArtistInfo             *ArtistInfo              `xml:"artistInfo"`
	ArtistInfo2            *ArtistInfo2             `xml:"artistInfo2"`
	SimilarSongs           *similarSongs            `xml:"similarSongs"`
	SimilarSongs2          *similarSongs2           `xml:"similarSongs2"`
	TopSongs               *topSongs                `xml:"topSongs"`
	ScanStatus             *ScanStatus              `xml:"scanStatus"`
	Error                  *Error                   `xml:"error"`
	Status                 string                   `xml:"status,attr"`  // May be one of ok, failed
	Version                string                   `xml:"version,attr"` // Must match the pattern \d+\.\d+\.\d+
	OpenSubsonic           bool                     `xml:"openSubsonic,attr"`
	OpenSubsonicExtensions []*OpenSubsonicExtension `xml:"openSubsonicExtensions"`
	LyricsList             *LyricsList              `xml:"lyricsList"`
}

type ScanStatus struct {
	Scanning bool  `xml:"scanning,attr"`
	Count    int64 `xml:"count,attr,omitempty"`
}

// SearchResult2 is a collection of songs, albums, and artists related to a query.
type SearchResult2 struct {
	Artist []*Artist `xml:"artist,omitempty"`
	Album  []*Child  `xml:"album,omitempty"`
	Song   []*Child  `xml:"song,omitempty"`
}

// SearchResult3 is a collection of songs, albums, and artists related to a query.
type SearchResult3 struct {
	Artist []*ArtistID3 `xml:"artist,omitempty"`
	Album  []*AlbumID3  `xml:"album,omitempty"`
	Song   []*Child     `xml:"song,omitempty"`
}

type Share struct {
	ID          string    `xml:"id,attr"`
	Entry       []*Child  `xml:"entry,omitempty"`
	Url         string    `xml:"url,attr"`
	Description string    `xml:"description,attr,omitempty"`
	Username    string    `xml:"username,attr"`
	Created     time.Time `xml:"created,attr"`
	Expires     time.Time `xml:"expires,attr,omitempty"`
	LastVisited time.Time `xml:"lastVisited,attr,omitempty"`
	VisitCount  int       `xml:"visitCount,attr"`
}

func (t *Share) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Share
	var layout struct {
		*T
		Created     *xsdDateTime `xml:"created,attr"`
		Expires     *xsdDateTime `xml:"expires,attr,omitempty"`
		LastVisited *xsdDateTime `xml:"lastVisited,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Expires = (*xsdDateTime)(&layout.T.Expires)
	layout.LastVisited = (*xsdDateTime)(&layout.T.LastVisited)
	return e.EncodeElement(layout, start)
}
func (t *Share) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Share
	var overlay struct {
		*T
		Created     *xsdDateTime `xml:"created,attr"`
		Expires     *xsdDateTime `xml:"expires,attr,omitempty"`
		LastVisited *xsdDateTime `xml:"lastVisited,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Expires = (*xsdDateTime)(&overlay.T.Expires)
	overlay.LastVisited = (*xsdDateTime)(&overlay.T.LastVisited)
	return d.DecodeElement(&overlay, &start)
}

type shares struct {
	Share []*Share `xml:"share,omitempty"`
}

type similarSongs struct {
	Song []*Child `xml:"song,omitempty"`
}

type similarSongs2 struct {
	Song []*Child `xml:"song,omitempty"`
}

type songs struct {
	Song []*Child `xml:"song,omitempty"`
}

// Starred is a collection of songs, albums, and artists annotated by a user as starred.
type Starred struct {
	Artist []*Artist `xml:"artist,omitempty"`
	Album  []*Child  `xml:"album,omitempty"`
	Song   []*Child  `xml:"song,omitempty"`
}

// Starred2 is a collection of songs, albums, and artists organized by ID3 tags annotated by a user as starred.
type Starred2 struct {
	Artist []*ArtistID3 `xml:"artist,omitempty"`
	Album  []*AlbumID3  `xml:"album,omitempty"`
	Song   []*Child     `xml:"song,omitempty"`
}

type topSongs struct {
	Song []*Child `xml:"song,omitempty"`
}

type User struct {
	Folder              []int     `xml:"folder,omitempty"`
	Username            string    `xml:"username,attr"`
	Email               string    `xml:"email,attr,omitempty"`
	ScrobblingEnabled   bool      `xml:"scrobblingEnabled,attr"`
	MaxBitRate          int       `xml:"maxBitRate,attr,omitempty"`
	AdminRole           bool      `xml:"adminRole,attr"`
	SettingsRole        bool      `xml:"settingsRole,attr"`
	DownloadRole        bool      `xml:"downloadRole,attr"`
	UploadRole          bool      `xml:"uploadRole,attr"`
	PlaylistRole        bool      `xml:"playlistRole,attr"`
	CoverArtRole        bool      `xml:"coverArtRole,attr"`
	CommentRole         bool      `xml:"commentRole,attr"`
	PodcastRole         bool      `xml:"podcastRole,attr"`
	StreamRole          bool      `xml:"streamRole,attr"`
	JukeboxRole         bool      `xml:"jukeboxRole,attr"`
	ShareRole           bool      `xml:"shareRole,attr"`
	VideoConversionRole bool      `xml:"videoConversionRole,attr"`
	AvatarLastChanged   time.Time `xml:"avatarLastChanged,attr,omitempty"`
}

func (t *User) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T User
	var layout struct {
		*T
		AvatarLastChanged *xsdDateTime `xml:"avatarLastChanged,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvatarLastChanged = (*xsdDateTime)(&layout.T.AvatarLastChanged)
	return e.EncodeElement(layout, start)
}
func (t *User) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T User
	var overlay struct {
		*T
		AvatarLastChanged *xsdDateTime `xml:"avatarLastChanged,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvatarLastChanged = (*xsdDateTime)(&overlay.T.AvatarLastChanged)
	return d.DecodeElement(&overlay, &start)
}

type users struct {
	User []*User `xml:"user,omitempty"`
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	if s == "" {
		*t = time.Time{}
		return nil
	}
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
