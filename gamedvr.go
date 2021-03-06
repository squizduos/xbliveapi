package xbliveapi

import (
	"fmt"
	"net/url"
	"time"
)

// GameDVR represents an individual user's game video record.
type GameDVR struct {
	XUID             string       `json:"xuid"`
	ClipName         string       `json:"clipName"`
	TitleID          int          `json:"titleId"`
	TitleName        string       `json:"titleName"`
	TitleData        string       `json:"titleData"`
	GameClipLocale   string       `json:"gameClipLocale"`
	GameClipID       string       `json:"gameClipId"`
	State            string       `json:"state"`
	DateRecorded     time.Time    `json:"dateRecorded"`
	LastModified     time.Time    `json:"lastModified"`
	Caption          string       `json:"userCaption"`
	Type             string       `json:"type"`
	Source           string       `json:"source"`
	Visibility       string       `json:"visibility"`
	Duration         int          `json:"durationInSeconds"`
	SCID             string       `json:"scid"`
	Rating           float64      `json:"rating"`
	RatingCount      int          `json:"ratingCount"`
	Views            int          `json:"views"`
	Properties       string       `json:"systemProperties"`
	SavedByUser      bool         `json:"savedByUser"`
	AchievementID    string       `json:"achievementId"`
	GreatestMomentID string       `json:"greatestMomentId"`
	Thumbnails       []GameDVRURI `json:"thumbnails"`
	URIs             []GameDVRURI `json:"gameClipUris"`
}

// GameDVRThumbnail describes the thumbnail of a particular game clip.
type GameDVRThumbnail struct {
	URI      string `json:"uri"`
	FileSize int64  `json:"fileSize"`
	Type     string `json:"thumbnailType"`
}

type GameDVRURI struct {
	URI        string    `json:"uri"`
	URIType    string    `json:"uriType"`
	FileSize   int64     `json:"fileSize"`
	Expiration time.Time `json:"expiration"`
}

type gameDVRResponse struct {
	GameDVRs   []GameDVR  `json:"gameClips"`
	PagingInfo pagingInfo `json:"pagingInfo"`
}

// GameDVRs retrieves all game clips for the provided user XID.
func (c *Client) GameDVRs(xid string) ([]GameDVR, error) {
	queryParams := url.Values{"maxItems": {"1000"}, "orderBy": {"EndingSoon"}}
	u := url.URL{
		Scheme:   "https",
		Host:     "gameclipsmetadata.xboxlive.com",
		Path:     fmt.Sprintf("/users/xuid(%s)/clips", xid),
		RawQuery: queryParams.Encode(),
	}

	var resp gameDVRResponse
	err := c.get(u.String(), vBoth, &resp)
	if err != nil {
		return nil, err
	}

	gamedvrs := make([]GameDVR, 0, resp.PagingInfo.TotalRecords)
	gamedvrs = append(gamedvrs, resp.GameDVRs...)
	for resp.PagingInfo.ContinuationToken != nil {
		queryParams.Set("continuationToken", *resp.PagingInfo.ContinuationToken)
		u.RawQuery = queryParams.Encode()

		err := c.get(u.String(), vBoth, &resp)
		if err != nil {
			return nil, err
		}
		gamedvrs = append(gamedvrs, resp.GameDVRs...)
	}
	return gamedvrs, nil
}
