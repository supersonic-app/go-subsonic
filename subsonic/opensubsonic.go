package subsonic

import "net/url"

const (
	SongLyricsExtension = "songLyrics"
	TranscodeOffset     = "transcodeOffset"
	IndexBasedQueue     = "indexBasedQueue"
)

// Get the list of supported OpenSubsonic extensions for this server.
func (c *Client) GetOpenSubsonicExtensions() ([]*OpenSubsonicExtension, error) {
	resp, err := c.Get("getOpenSubsonicExtensions", nil)
	if err != nil {
		return nil, err
	}
	if resp.OpenSubsonicExtensions == nil {
		return nil, nil
	}
	return resp.OpenSubsonicExtensions, nil
}

// Get structured lyrics for a track.
//
// Server must support OpenSubsonic songLyrics extension
func (c *Client) GetLyricsBySongId(songID string) (*LyricsList, error) {
	resp, err := c.Get("getLyricsBySongId", map[string]string{"id": songID})
	if err != nil {
		return nil, err
	}
	return resp.LyricsList, nil
}

// GetPlayQueueByIndex returns the state of the play queue for this user
// (as set by savePlayQueueByIndex). This includes the tracks in the
// play queue, the currently playing track by index, and the position
// within this track.
//
// Server must support OpenSubsonic playQueueByIndex extension.
func (c *Client) GetPlayQueueByIndex() (*PlayQueueByIndex, error) {
	resp, err := c.Get("getPlayQueueByIndex", nil)
	if err != nil {
		return nil, err
	}
	return resp.PlayQueueByIndex, nil
}

// SavePlayQueueByIndex saves the state of the play queue for this user.
// This includes the tracks in the play queue, the currently playing
// track by index, and the position within this track.
//
// Parameters:
//
//	songIDs: IDs of the songs in the play queue
//
// Optional parameters:
//
//	currentIndex: the index of the currently playing song (0-based)
//	position: The position in milliseconds within the currently playing song
//
// Server must support OpenSubsonic playQueueByIndex extension.
func (c *Client) SavePlayQueueByIndex(songIDs []string, params map[string]string) error {
	values := url.Values{}
	for _, trID := range songIDs {
		values.Add("id", trID)
	}
	for k, v := range params {
		values.Add(k, v)
	}
	_, err := c.getValues("savePlayQueueByIndex", values)
	return err
}
