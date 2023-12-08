package subsonic

import (
	"encoding/xml"
	"testing"
)

const (
	songLyricsResp string = `<subsonic-response status="ok" version="1.16.1" type="AwesomeServerName" serverVersion="0.1.3 (tag)" openSubsonic="true">
    <lyricsList>
      <structuredLyrics displayArtist="Muse" displayTitle="Hysteria" lang="en" offset="-100" synced="true">
        <line start="0">It's bugging me</line>
        <line start="2000">Grating me</line>
        <line start="3001">And twisting me around...</line>
      </structuredLyrics>
      <structuredLyrics displayArtist="Muse" displayTitle="Hysteria" lang="en" offset="100" synced="false">
        <line>It's bugging me</line>
        <line>Grating me</line>
        <line>And twisting me around...</line>
      </structuredLyrics>
    </lyricsList>
  </subsonic-response>`
)

func TestLyricsBySongId(t *testing.T) {
	var response Response
	err := xml.Unmarshal([]byte(songLyricsResp), &response)
	if err != nil {
		t.Errorf("Error unmarshaling lyrics XML: %v", err)
	}
	if response.OpenSubsonic == false {
		t.Error("wrong OpenSubsonic")
	}
	if response.LyricsList == nil {
		t.Error("nil response.LyricsList")
	}
	if l := len(response.LyricsList.StructuredLyrics); l != 2 {
		t.Errorf("wrong length of StructuredLyrics (want 2, got %d)", l)
	}
	if o := response.LyricsList.StructuredLyrics[0].Offset; o != -100 {
		t.Errorf("wrong lyric offset (want -100, got %d)", o)
	}
	if l := len(response.LyricsList.StructuredLyrics[0].Lines); l != 3 {
		t.Errorf("wrong line count (want 3, got %d)", l)
	}
	line := response.LyricsList.StructuredLyrics[0].Lines[1]
	if line.Start != 2000 {
		t.Errorf("wrong line start (want 2000, got %d)", line.Start)
	}
	if line.Text != "Grating me" {
		t.Errorf("wrong line text, want %q, got %q)", "Grating me", line.Text)
	}
}
