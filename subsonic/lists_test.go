package subsonic

import (
	"encoding/xml"
	"testing"
	"time"
)

const openSubsonicAlbumList = `<?xml version="1.0" encoding="utf-8"?>
<subsonic-response openSubsonic="true" serverVersion="1" status="ok" type="lms" version="1.16.0"><albumList2><album artist="Au5, Danyka Nadeau" coverArt="al-5" created="2023-10-12T18:23:30.000" duration="36" genre="Unknown" id="al-5" isCompilation="false" musicBrainzId="7a4d48b2-66b7-4a93-afbf-2b176a0c92a6" name="Follow You (The Remixes)" originalReleaseDate="2014-07-11" songCount="7" year="2014"><artists id="ar-52" name="Au5"/><artists id="ar-53" name="Danyka Nadeau"/><genres name="Unknown"/><releaseTypes>ep</releaseTypes></album><album artist="Iron Maiden" artistId="ar-45" coverArt="al-4" created="2023-10-12T18:13:19.000" duration="120" genre="Heavy Metal" id="al-4" isCompilation="true" musicBrainzId="1e84b11c-6ea1-4685-87b9-ae35591809d7" name="A Real Live Dead One" originalReleaseDate="1993-03-23" songCount="23" year="1993"><artists id="ar-45" name="Iron Maiden"/><discTitles disc="1" title="Dead One"/><discTitles disc="2" title="Live One"/><genres name="Instrumental"/><genres name="Heavy Metal"/><genres name="Metal"/><releaseTypes>album</releaseTypes><releaseTypes>compilation</releaseTypes><releaseTypes>live</releaseTypes></album><album artist="DJ Banana" artistId="ar-44" coverArt="al-3" created="2023-05-28T18:35:21.000" duration="57" genre="Techno-Industrial" id="al-3" isCompilation="false" musicBrainzId="667851cb-0f84-3fdd-8882-33902fa16aef" name="Fruit Salad Mixtape" originalReleaseDate="" songCount="11" year="2001"><artists id="ar-44" name="DJ Banana"/><genres name="Techno-Industrial"/></album><album artist="Willbe" artistId="ar-2" coverArt="al-2" created="2020-11-19T13:17:03.000" duration="4797" genre="Unknown" id="al-2" isCompilation="true" musicBrainzId="31b3b92b-b212-43de-b57d-7f2961b7cc4a" name="Demovibes 3: Pixels in sequence" originalReleaseDate="2004-06-23" songCount="22" year="2004"><artists id="ar-2" name="Willbe"/><genres name="Electronic"/><genres name="Hip-Hop"/><genres name="Industrial"/><genres name="Rap"/><genres name="Dance"/><genres name="Happy Hardcore"/><genres name="Rave"/><genres name="Trance"/><genres name="Downtempo"/><genres name="Idm"/><genres name="Unknown"/><genres name="Lounge"/><genres name="Breaks"/><genres name="Minimal"/><genres name="Breakbeat"/><genres name="Gothic Metal"/><genres name="Hardcore"/><genres name="House"/><genres name="Instrumental"/><genres name="Ambient"/><genres name="Jazz"/><genres name="Indie"/><genres name="Techno"/><genres name="Stoner Rock"/><genres name="Heavy Metal"/><releaseTypes>album</releaseTypes><releaseTypes>compilation</releaseTypes></album><album artist="Willbe" artistId="ar-2" coverArt="al-1" created="2019-11-27T20:07:27.000" duration="4797" genre="Unknown" id="al-1" isCompilation="true" musicBrainzId="91eaab8e-b472-42f9-a6e4-9e06d0464f31" name="Demovibes 5: The mod inside" originalReleaseDate="2006-04-16" songCount="20" year="2006"><artists id="ar-2" name="Willbe"/><genres name="Disco"/><genres name="Electronic"/><genres name="New Wave"/><genres name="Christian"/><genres name="Hip-Hop"/><genres name="Industrial"/><genres name="Rap"/><genres name="Dance"/><genres name="Happy Hardcore"/><genres name="Rave"/><genres name="Trance"/><genres name="Downtempo"/><genres name="Idm"/><genres name="Unknown"/><genres name="Lounge"/><genres name="Progressive Rock"/><genres name="Psychedelic Rock"/><genres name="Avant-Garde"/><genres name="Drum and Bass"/><genres name="Breaks"/><genres name="Minimal"/><genres name="Breakbeat"/><releaseTypes>album</releaseTypes><releaseTypes>compilation</releaseTypes></album></albumList2></subsonic-response>`

func TestUnmarshalOpenSubsonicAlbumList(t *testing.T) {
	parsed := Response{}
	err := xml.Unmarshal([]byte(openSubsonicAlbumList), &parsed)
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}
	if parsed.AlbumList2 == nil {
		t.Fatal("No AlbumList2")
	}
	if len(parsed.AlbumList2.Album) < 1 {
		t.Error("No albums in AlbumList2")
	}
	if parsed.AlbumList2.Album[1].Artist == "" {
		t.Error("Did not parse album correctly")
	}
	if len(parsed.AlbumList2.Album[0].Artists) < 2 {
		t.Error("Did not parse OpenSubsonic Artists attribute correctly")
	}
}

func runListsTests(client Client, t *testing.T) {
	sampleGenre := getSampleGenre(client)

	t.Run("GetAlbumList", func(t *testing.T) {
		_, err := client.GetAlbumList("foobar", nil)
		if err == nil {
			t.Error("No error was returned with an invalid listType argument")
		}
		_, err = client.GetAlbumList("byYear", nil)
		if err == nil {
			t.Error("Failed to validate byYear parameters")
		}
		_, err = client.GetAlbumList("byYear", map[string]string{"fromYear": "1990"})
		if err == nil {
			t.Error("Failed to validate partial byYear parameters")
		}
		_, err = client.GetAlbumList("byGenre", nil)
		if err == nil {
			t.Error("Failed to validate byGenre parameters")
		}
		albums, err := client.GetAlbumList("random", nil)
		if err != nil {
			t.Error(err)
		}
		if albums == nil {
			t.Error("No albums were returned in a call to random getAlbumList")
		}
		for _, album := range albums {
			if album.Title == "" {
				t.Errorf("Album %#v has an empty name :(", album)
			}
		}
		// Work out genre matching
		albums, err = client.GetAlbumList("byGenre", map[string]string{"genre": sampleGenre.Name})
		if err != nil {
			t.Error(err)
		}
		if albums == nil || len(albums) < 1 {
			t.Error("No albums were returned in a call to a byGenre getAlbumList")
		}
		var empty time.Time
		for _, album := range albums {
			if album.Created == empty {
				t.Errorf("Album %#v has empty created time", album)
			}
		}
	})

	t.Run("GetAlbumList2", func(t *testing.T) {
		// Test incorrect parameters
		_, err := client.GetAlbumList2("foobar", nil)
		if err == nil {
			t.Error("No error was returned with an invalid listType argument")
		}
		_, err = client.GetAlbumList2("byYear", nil)
		if err == nil {
			t.Error("Failed to validate byYear parameters")
		}
		_, err = client.GetAlbumList2("byYear", map[string]string{"fromYear": "1990"})
		if err == nil {
			t.Error("Failed to validate partial byYear parameters")
		}
		_, err = client.GetAlbumList2("byGenre", nil)
		if err == nil {
			t.Error("Failed to validate byGenre parameters")
		}
		// Test with proper parameters
		albums, err := client.GetAlbumList2("newest", nil)
		if err != nil {
			t.Error(err)
		}
		if albums == nil {
			t.Error("No albums were returned in a call to newest getAlbumList2")
		}
		var empty time.Time
		for _, album := range albums {
			if album.Name == "" {
				t.Errorf("Album %#v has an empty name", album)
			}
			if album.Created == empty {
				t.Errorf("Album %#v has empty created time", album)
			}
		}
	})

	t.Run("GetRandomSongs", func(t *testing.T) {
		songs, err := client.GetRandomSongs(nil)
		if err != nil || songs == nil {
			t.Error("Basic call to getRandomSongs failed")
		}
		var empty time.Time
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
		songs, err = client.GetRandomSongs(map[string]string{"size": "1"})
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by getRandomSongs failed: expected 1, length actual %d", len(songs))
		}
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
	})

	t.Run("GetSongsByGenre", func(t *testing.T) {
		songs, err := client.GetSongsByGenre(sampleGenre.Name, nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Errorf("No songs returned for genre %v", sampleGenre)
		}
		songs, err = client.GetSongsByGenre(sampleGenre.Name, map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by GetSongsByGenre failed: expected 1, length actual %d", len(songs))
		}
		var empty time.Time
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
	})

	t.Run("GetNowPlaying", func(t *testing.T) {
		// This test is essentially a no-op because we can't depend on the state of playing something in a test environment
		entries, err := client.GetNowPlaying()
		if err != nil {
			t.Error(err)
		}
		var empty time.Time
		for _, nowPlaying := range entries {
			//t.Logf("NowPlaying %d minutes ago, created %v", nowPlaying.MinutesAgo, nowPlaying.Created.Format("2006-01-02T15:04:05.999999-07:00"))
			if nowPlaying.Created == empty {
				t.Errorf("NowPlayingEntry %#v had an empty created", nowPlaying)
			}
		}
	})

	t.Run("GetStarred", func(t *testing.T) {
		// State dependent test
		_, err := client.GetStarred(nil)
		if err != nil {
			t.Error(err)
		}
		_, err = client.GetStarred2(nil)
		if err != nil {
			t.Error(err)
		}
	})
}
