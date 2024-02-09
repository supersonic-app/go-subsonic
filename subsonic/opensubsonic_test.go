package subsonic

import (
	"bytes"
	"testing"
)

const getOpenSubsonicExtensionsResponse = `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.16.1" type="navidrome" serverVersion="0.51.0 (fd61b29a)" openSubsonic="true">
<openSubsonicExtensions name="transcodeOffset">
<versions>1</versions>
</openSubsonicExtensions>
<openSubsonicExtensions name="formPost">
<versions>1</versions>
</openSubsonicExtensions>
<openSubsonicExtensions name="songLyrics">
<versions>1</versions>
</openSubsonicExtensions>
</subsonic-response>`

func Test_GetOpenSubsonicExtensions(t *testing.T) {
	resp := bytes.NewReader([]byte(getOpenSubsonicExtensionsResponse))
	unmarshaled, err := unmarshalResponse(resp)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	if l := len(unmarshaled.OpenSubsonicExtensions); l != 3 {
		t.Errorf("Wrong number of OpenSubsonic extensions: %d", l)
	}
	if n := unmarshaled.OpenSubsonicExtensions[0].Name; n != "transcodeOffset" {
		t.Errorf("Wrong name for OpenSubsonic extension: %s", n)
	}
	if l := len(unmarshaled.OpenSubsonicExtensions[0].Versions); l != 1 {
		t.Errorf("Wrong number of OpenSubsonic extension versions: %d", l)
	}
	if v := unmarshaled.OpenSubsonicExtensions[0].Versions[0]; v != 1 {
		t.Errorf("Wrong OpenSubsonic extension version: %d", v)
	}
}
