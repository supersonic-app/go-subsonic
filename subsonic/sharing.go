package subsonic

// GetShares returns information about shared media this user is allowed to manage.
func (c *Client) GetShares() ([]*Share, error) {
	resp, err := c.Get("getShares", nil)
	if err != nil {
		return nil, err
	}
	return resp.Shares.Share, nil
}

// CreateShare creates a public URL that can be used by anyone to stream music or video from the
// Subsonic server. The URL is short and suitable for posting on Facebook, Twitter etc. Note: The
// user must be authorized to share.
//
// Optional Parameters:
//
//	description: A user-defined description that will be displayed to people visiting the shared media.
//	expires:     The time at which the share expires. Given as milliseconds since 1970.
func (c *Client) CreateShare(id string, parameters map[string]string) (*Share, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := c.Get("createShare", params)
	if err != nil {
		return nil, err
	}
	return resp.Shares.Share[0], nil
}

// UpdateShare updates the description and/or expiration date for an existing share.
//
// Optional Parameters:
//
//	description: A user-defined description that will be displayed to people visiting the shared media.
//	expires:     The time at which the share expires. Given as milliseconds since 1970, or zero to remove the expiration.
func (c *Client) UpdateShare(id string, parameters map[string]string) error {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	_, err := c.Get("updateShare", params)
	if err != nil {
		return err
	}
	return nil
}

// DeleteShare deletes an existing share.
func (c *Client) DeleteShare(id string) error {
	_, err := c.Get("deleteShare", map[string]string{"id": id})
	if err != nil {
		return err
	}
	return nil
}
