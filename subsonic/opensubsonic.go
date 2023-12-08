package subsonic

const (
	SongLyricsExtension = "songLyrics"
)

func (c *Client) GetOpenSubsonicExtensions() ([]*OpenSubsonicExtension, error) {
	resp, err := c.Get("getPlaylists", nil)
	if err != nil {
		return nil, err
	}
	if resp.OpenSubsonicExtensions == nil {
		return nil, nil
	}
	return resp.OpenSubsonicExtensions.OpenSubsonicExtensions, nil
}

func (c *Client) GetLyricsBySongId(songID string) (*LyricsList, error) {
	resp, err := c.Get("getLyricsBySongId", map[string]string{"id": songID})
	if err != nil {
		return nil, err
	}
	return resp.LyricsList, nil
}
