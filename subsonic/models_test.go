package subsonic

import (
	"encoding/xml"
	"testing"
)

// XML example responses from official subsonic.org API documentation
const (
	albumXML = `<album id="11053" name="High Voltage" coverArt="al-11053" songCount="8" created="2004-11-27T20:23:32" duration="2414" artist="AC/DC" artistId="5432">
	<originalReleaseDate year="2020" month="02"/>
	<song id="71463" parent="71381" title="The Jack" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-08T23:36:11" duration="352" bitRate="128" size="5624132" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - The Jack.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71464" parent="71381" title="Tnt" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-08T23:36:11" duration="215" bitRate="128" size="3433798" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - TNT.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71458" parent="71381" title="It's A Long Way To The Top" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-27T20:23:32" duration="315" bitRate="128" year="1976" genre="Rock" size="5037357" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - It's a long way to the top if you wanna rock 'n 'roll.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71461" parent="71381" title="Rock 'n' Roll Singer." album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-27T20:23:33" duration="303" bitRate="128" track="2" year="1976" genre="Rock" size="4861680" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - Rock N Roll Singer.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71460" parent="71381" title="Live Wire" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-27T20:23:33" duration="349" bitRate="128" track="4" year="1976" genre="Rock" size="5600206" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - Live Wire.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71456" parent="71381" title="Can I sit next to you girl" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-27T20:23:32" duration="251" bitRate="128" track="6" year="1976" genre="Rock" size="4028276" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - Can I Sit Next To You Girl.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71459" parent="71381" title="Little Lover" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-27T20:23:33" duration="339" bitRate="128" track="7" year="1976" genre="Rock" size="5435119" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - Little Lover.mp3" albumId="11053" artistId="5432" type="music"/>
	<song id="71462" parent="71381" title="She's Got Balls" album="High Voltage" artist="AC/DC" isDir="false" coverArt="71381" created="2004-11-27T20:23:34" duration="290" bitRate="128" track="8" year="1976" genre="Rock" size="4651866" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="ACDC/High voltage/ACDC - Shes Got Balls.mp3" albumId="11053" artistId="5432" type="music"/>
	</album>`

	artistXML = `<artist id="5432" name="AC/DC" coverArt="ar-5432" albumCount="15">
	<album id="11047" name="Back In Black" coverArt="al-11047" songCount="10" created="2004-11-08T23:33:11" duration="2534" artist="AC/DC" artistId="5432"/>
	<album id="11048" name="Black Ice" coverArt="al-11048" songCount="15" created="2008-10-30T09:20:52" duration="3332" artist="AC/DC" artistId="5432"/>
	<album id="11049" name="Blow up your Video" coverArt="al-11049" songCount="10" created="2004-11-27T19:22:45" duration="2578" artist="AC/DC" artistId="5432"/>
	<album id="11050" name="Flick Of The Switch" coverArt="al-11050" songCount="10" created="2004-11-27T19:22:51" duration="2222" artist="AC/DC" artistId="5432"/>
	<album id="11051" name="Fly On The Wall" coverArt="al-11051" songCount="10" created="2004-11-27T19:22:57" duration="2405" artist="AC/DC" artistId="5432"/>
	<album id="11052" name="For Those About To Rock" coverArt="al-11052" songCount="10" created="2004-11-08T23:35:02" duration="2403" artist="AC/DC" artistId="5432"/>
	<album id="11053" name="High Voltage" coverArt="al-11053" songCount="8" created="2004-11-27T20:23:32" duration="2414" artist="AC/DC" artistId="5432"/>
	<album id="10489" name="Highway To Hell" coverArt="al-10489" songCount="12" created="2009-06-15T09:41:54" duration="2745" artist="AC/DC" artistId="5432"/>
	<album id="11054" name="If You Want Blood..." coverArt="al-11054" songCount="1" created="2004-11-27T20:23:32" duration="304" artist="AC/DC" artistId="5432"/>
	<album id="11056" name="Let There Be Rock" coverArt="al-11056" songCount="8" created="2004-11-27T20:33:40" duration="2449" artist="AC/DC" artistId="5432"/>
	<album id="11057" name="Live - Special Collector's Edition" coverArt="al-11057" songCount="22" created="2004-11-08T23:37:09" duration="6999" artist="AC/DC" artistId="5432"/>
	<album id="11058" name="Powerage" coverArt="al-11058" songCount="9" created="2004-11-27T20:33:41" duration="2380" artist="AC/DC" artistId="5432"/>
	<album id="11059" name="Stiff Upper Lip" coverArt="al-11059" songCount="11" created="2004-11-08T23:41:13" duration="2595" artist="AC/DC" artistId="5432"/>
	<album id="11060" name="The Razors Edge" coverArt="al-11060" songCount="12" created="2004-11-27T20:33:42" duration="2787" artist="AC/DC" artistId="5432"/>
	<album id="11061" name="Who Made Who" coverArt="al-11061" songCount="9" created="2004-11-08T23:43:18" duration="2291" artist="AC/DC" artistId="5432"/>
	</artist>`

	playlistXML = `<playlist id="15" name="kokos" comment="fan" owner="admin" public="true" songCount="6" duration="1391" created="2012-04-17T19:53:44" coverArt="pl-15">
	<allowedUser>sindre</allowedUser>
	<allowedUser>john</allowedUser>
	<entry id="657" parent="655" title="Making Me Nervous" album="I Don't Know What I'm Doing" artist="Brad Sucks" isDir="false" coverArt="655" created="2008-04-10T07:10:32" duration="159" bitRate="202" track="1" year="2003" size="4060113" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="Brad Sucks/I Don't Know What I'm Doing/01 - Making Me Nervous.mp3" albumId="58" artistId="45" type="music"/>
	<entry id="823" parent="784" title="Piano escena" album="BSO Sebastian" artist="PeerGynt Lobogris" isDir="false" coverArt="784" created="2009-01-14T22:26:29" duration="129" bitRate="170" track="8" year="2008" genre="Blues" size="2799954" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="PeerGynt Lobogris/BSO Sebastian/08 - Piano escena.mp3" albumId="75" artistId="54" type="music"/>
	<entry id="748" parent="746" title="Stories from Emona II" album="Between two worlds" artist="Maya Filipič" isDir="false" coverArt="746" created="2008-07-30T22:05:40" duration="335" bitRate="176" track="2" year="2008" genre="Classical" size="7458214" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="Maya Filipic/Between two worlds/02 - Stories from Emona II.mp3" albumId="68" artistId="51" type="music"/>
	<entry id="848" parent="827" title="Run enemy" album="Eve" artist="Shearer" isDir="false" coverArt="827" created="2009-01-15T22:54:38" duration="331" bitRate="195" track="14" year="2008" genre="Rock" size="8160185" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="Shearer/Eve/14 - Run enemy.mp3" albumId="77" artistId="55" type="music"/>
	<entry id="884" parent="874" title="Isolation" album="Kosmonaut" artist="Ugress" isDir="false" coverArt="874" created="2009-01-14T21:34:49" duration="320" bitRate="160" track="4" year="2006" genre="Electronic" size="6412176" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="Ugress/Kosmonaut/Ugress-KosmonautEP-04-Isolation.mp3" albumId="81" artistId="57" type="music"/>
	<entry id="805" parent="783" title="Bajo siete lunas (intro)" album="Broken Dreams" artist="PeerGynt Lobogris" isDir="false" coverArt="783" created="2008-12-19T14:13:58" duration="117" bitRate="225" track="1" year="2008" genre="Blues" size="3363271" suffix="mp3" contentType="audio/mpeg" isVideo="false" path="PeerGynt Lobogris/Broken Dreams/01 - Bajo siete lunas (intro).mp3" albumId="74" artistId="54" type="music"/>
	</playlist>`

	directoryXML = `<directory id="10" parent="9" name="ABBA" starred="2013-11-02T12:30:00">
	<child id="11" parent="10" title="Arrival" artist="ABBA" isDir="true" coverArt="22"/>
	<child id="12" parent="10" title="Super Trouper" artist="ABBA" isDir="true" coverArt="23"/>
	</directory>`

	playQueueXML = `<playQueue current="133" position="45000" username="admin" changed="2015-02-18T15:22:22.825Z" changedBy="android">
	<entry id="132" parent="131" isDir="false" title="These Are Days" album="MTV Unplugged" artist="10,000 Maniacs" track="1" year="1993" genre="Soft Rock" coverArt="131" size="5872262" contentType="audio/mpeg" suffix="mp3" duration="293" bitRate="160" path="10,000 Maniacs/MTV Unplugged/01 - These Are Days.mp3" isVideo="false" created="2004-10-25T20:36:03.000Z" albumId="0" artistId="0" type="music"/>
	<entry id="133" parent="131" isDir="false" title="Eat For Two" album="MTV Unplugged" artist="10,000 Maniacs" track="2" year="1993" genre="Soft Rock" coverArt="131" size="5253248" contentType="audio/mpeg" suffix="mp3" duration="262" bitRate="160" path="10,000 Maniacs/MTV Unplugged/02 - Eat For Two.mp3" isVideo="false" created="2004-10-25T20:36:06.000Z" albumId="0" artistId="0" type="music"/>
	<entry id="134" parent="131" isDir="false" title="Candy Everybody Wants" album="MTV Unplugged" artist="10,000 Maniacs" track="3" year="1993" genre="Soft Rock" coverArt="131" size="3993728" contentType="audio/mpeg" suffix="mp3" duration="199" bitRate="160" path="10,000 Maniacs/MTV Unplugged/03 - Candy Everybody Wants.mp3" isVideo="false" created="2004-10-25T20:36:09.000Z" albumId="0" artistId="0" type="music"/>
	</playQueue>`
)

func TestUnmarshalAlbum(t *testing.T) {
	var a AlbumID3
	err := xml.Unmarshal([]byte(albumXML), &a)
	if err != nil {
		t.Errorf("Error unmarshaling album XML: %v", err)
	}
	if a.Created.IsZero() {
		t.Error("Failed to unmarshal album Created timestamp")
	}
	if len(a.Song) < 1 {
		t.Error("Failed to unmarshal album songs")
	}
	if a.Song[0].Created.IsZero() {
		t.Error("Failed to unmarshal album song timestamp")
	}
	if *a.OriginalReleaseDate.Year != 2020 {
		t.Error("Failed to unmarshal originalReleaseDate properly")
	}
	if a.OriginalReleaseDate.Date != nil {
		t.Error("Failed to unmarshal originalReleaseDate properly")
	}
}

func TestUnmarshalArtist(t *testing.T) {
	var a ArtistID3
	err := xml.Unmarshal([]byte(artistXML), &a)
	if err != nil {
		t.Errorf("Error unmarshaling artist XML: %v", err)
	}
	if len(a.Album) < 1 {
		t.Error("Failed to unmarshal artist albums")
	}
	if a.Album[0].Created.IsZero() {
		t.Error("Failed to unmarshal artist album timestamp")
	}
}

func TestUnmarshalPlaylist(t *testing.T) {
	var p Playlist
	err := xml.Unmarshal([]byte(playlistXML), &p)
	if err != nil {
		t.Errorf("Error unmarshaling album XML: %v", err)
	}
	if p.Created.IsZero() {
		t.Error("Failed to unmarshal album Created timestamp")
	}
	if len(p.Entry) < 1 {
		t.Error("Failed to unmarshal playlist songs")
	}
	if p.Entry[0].Created.IsZero() {
		t.Error("Failed to unmarshal playlist song timestamp")
	}
}

func TestUnmarshalDirectory(t *testing.T) {
	var d Directory
	err := xml.Unmarshal([]byte(directoryXML), &d)
	if err != nil {
		t.Errorf("Error unmarshaling directory XML: %v", err)
	}
	if d.Starred.IsZero() {
		t.Error("Failed to unmarshal directory Starred timestamp")
	}
	if len(d.Child) < 1 {
		t.Error("Failed to unmarshal directory contents")
	}
}

func TestUnmarshalPlayQueue(t *testing.T) {
	var p PlayQueue
	err := xml.Unmarshal([]byte(playQueueXML), &p)
	if err != nil {
		t.Errorf("Error unmarshaling play queue XML: %v", err)
	}
	if p.Changed.IsZero() {
		t.Error("Failed to unmarshal play queue Changed timestamp")
	}
	if p.Current != "133" {
		t.Error("Failed to unmarshal play queue current ID properly")
	}
	if len(p.Entries) < 1 {
		t.Error("Failed to unmarshal play queue contents")
	}
	if p.Entries[0].Title != "These Are Days" {
		t.Error("Failed to unmarshal play queue entry properly")
	}
}
