package subsonic

import (
	"net/url"
)

// GetPlayQueue Returns the state of the play queue for this user
// (as set by savePlayQueue). This includes the tracks in the
// play queue, the currently playing track, and the position
// within this track. Typically used to allow a user to move
// between different clients/apps while retaining the same
// play queue (for instance when listening to an audio book).
func (c *Client) GetPlayQueue() (*PlayQueue, error) {
	resp, err := c.Get("getPlayQueue", nil)
	if err != nil {
		return nil, err
	}
	return resp.PlayQueue, nil
}

// SavePlayQueue saves the state of the play queue for this user.
// This includes the tracks in the play queue, the currently playing
// track, and the position within this track. Typically used to allow
// a user to move between different clients/apps while retaining the
// same play queue (for instance when listening to an audio book).
//
// Parameters:
//
//	songIDs: IDs of the songs in the play queue
//
// Optional parameters:
//
//	current: the ID of the currently playing song
//	position: The position in milliseconds within the currently playing song
func (c *Client) SavePlayQueue(songIDs []string, params map[string]string) error {
	values := url.Values{}
	for _, trID := range songIDs {
		values.Add("id", trID)
	}
	for k, v := range params {
		values.Add(k, v)
	}
	_, err := c.getValues("savePlayQueue", values)
	return err
}
