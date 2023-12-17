package subsonic

import (
	"errors"
	"fmt"
)

var validJukeboxControlArgs = []string{"status", "set", "start", "stop", "skip", "add", "clear", "remove", "shuffle", "setGain"}

// JukeboxControl controls the jukebox, i.e., playback directly on the server's audio hardware.
// Note: The user must be authorized to control the jukebox (see Settings > Users > User is
// allowed to play files in jukebox mode).
//
// Parameters:
//
//	action: The action to perform. Must be one of "status", "set", "start", "stop",
//	        "skip", "add", "clear", "remove", "shuffle", "setGain"
//	   Note: the "get" action is implemented as the GetJukeboxPlaylist() function.
//
// Optional Parameters:
//
//	index:  Used by skip and remove. Zero-based index of the song to skip to or remove.
//	offset: Used by skip. Start playing this many seconds into the track.
//	id:     Used by add and set. ID of song to add to the jukebox playlist. Use multiple id parameters to add many songs in the same request.
//	gain:   Used by setGain to control the playback volume. A float value between 0.0 and 1.0.
func (c *Client) JukeboxControl(action string, parameters map[string]string) (*JukeboxStatus, error) {
	// validate action arg
	if action == "get" {
		return nil, errors.New("use GetJukeboxPlaylist instead to retrieve the jukebox playlist")
	}
	valid := false
	for _, val := range validJukeboxControlArgs {
		if action == val {
			valid = true
			break
		}
	}
	if !valid {
		return nil, fmt.Errorf("unsupported jukebox action: %q", action)
	}

	params := make(map[string]string)
	params["action"] = action
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := c.Get("jukeboxControl", params)
	if err != nil {
		return nil, err
	}
	if resp.JukeboxStatus == nil {
		return nil, errors.New("server did not reply with jukebox status")
	}
	return resp.JukeboxStatus, nil
}

// GetJukeboxPlaylist retrieves the current jukebox playlist.
func (c *Client) GetJukeboxPlaylist() (*JukeboxPlaylist, error) {
	resp, err := c.Get("jukeboxControl", map[string]string{"action": "get"})
	if err != nil {
		return nil, err
	}
	if resp.JukeboxPlaylist == nil {
		return nil, errors.New("server did not return jukebox playlist")
	}
	return resp.JukeboxPlaylist, nil
}
