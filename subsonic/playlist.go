package subsonic

import (
	"errors"
	"net/url"
	"strconv"
)

// GetPlaylists returns all playlists a user is allowed to play.
//
// Optional Parameters:
//
//	user: get playlists visible to this username rather than the current user. Must have admin permission.
func (s *Client) GetPlaylists(parameters map[string]string) ([]*Playlist, error) {
	resp, err := s.Get("getPlaylists", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Playlists.Playlist, nil
}

// GetPlaylist returns a listing of files in a saved playlist.
func (s *Client) GetPlaylist(id string) (*Playlist, error) {
	resp, err := s.Get("getPlaylist", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Playlist, nil
}

// CreatePlaylist creates (or updates) a playlist.
//
// Optional Parameters:
//
//	songId:     ID of a song in the playlist. Use one songId parameter for each song in the playlist.
//
// Mutually Exclusive Parameters:
//
//	playlistId: The playlist ID.
//	name:       The human-readable name of the playlist.
//
// This returns a Playlist object in Subsonic > 1.14.0, so it cannot consistently return a *Playlist
func (s *Client) CreatePlaylist(parameters map[string]string) error {
	_, idPresent := parameters["playlistId"]
	_, namePresent := parameters["name"]
	if !(idPresent || namePresent) {
		return errors.New("One of name or playlistId is mandatory, to create or update a playlist respectively")
	}
	_, err := s.Get("createPlaylist", parameters)
	if err != nil {
		return err
	}
	return nil
}

// CreatePlaylistWithTracks creates (or updates) a playlist, with the given tracks being added.
//
// Mutually Exclusive Parameters:
//
//	playlistId: The playlist ID.
//	name:       The human-readable name of the playlist.
//
// This returns a Playlist object in Subsonic > 1.14.0, so it cannot consistently return a *Playlist
func (s *Client) CreatePlaylistWithTracks(trackIDs []string, parameters map[string]string) error {
	_, idPresent := parameters["playlistId"]
	_, namePresent := parameters["name"]
	if !(idPresent || namePresent) {
		return errors.New("one of name or playlistId is mandatory, to create or update a playlist respectively")
	}
	values := url.Values{}
	for _, trID := range trackIDs {
		values.Add("songId", trID)
	}
	for k, v := range parameters {
		values.Add(k, v)
	}
	_, err := s.getValues("createPlaylist", values)
	return err
}

// UpdatePlaylist updates a playlist. Only the owner of a playlist is allowed to update it.
//
// Optional Parameters:
//
//	name:              The human-readable name of the playlist.
//	comment:           The playlist comment.
//	public:            true if the playlist should be visible to all users, false otherwise.
//	songIdToAdd:       Add this song with this ID to the playlist. Multiple parameters allowed.
//	songIndexToRemove: Remove the song at this position in the playlist. Multiple parameters allowed.
func (s *Client) UpdatePlaylist(playlistId string, parameters map[string]string) error {
	params := make(map[string]string)
	for k, v := range parameters {
		params[k] = v
	}
	params["playlistId"] = playlistId
	_, err := s.Get("updatePlaylist", params)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePlaylistTracks updates the tracks in a playlist.
// Only the owner of a playlist is allowed to update it.
func (s *Client) UpdatePlaylistTracks(
	playlistId string,
	trackIDsToAdd []string,
	trackIndexesToRemove []int,
) error {
	values := url.Values{}
	values.Add("playlistId", playlistId)
	for _, trId := range trackIDsToAdd {
		values.Add("songIdToAdd", trId)
	}
	for _, trIdx := range trackIndexesToRemove {
		values.Add("songIndexToRemove", strconv.Itoa(trIdx))
	}
	_, err := s.getValues("updatePlaylist", values)
	return err
}

// DeletePlaylist deletes a saved playlist.
func (s *Client) DeletePlaylist(playlistId string) error {
	_, err := s.Get("deletePlaylist", map[string]string{"id": playlistId})
	if err != nil {
		return err
	}
	return nil
}
