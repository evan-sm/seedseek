package jackett

import (
	"context"
	"time"

	"github.com/wmw64/go-jackett"
)

type Jacketter interface {
	Fetch(context.Context, string, ...uint) ([]Result, error)
}

type Jackett struct {
	jackett *jackett.Jackett
}

type Result struct {
	BannerUrl            string
	BlackholeLink        string
	Category             []uint
	CategoryDesc         string
	Comments             string
	Details              string
	Description          string
	DownloadVolumeFactor float32
	Files                uint
	FirstSeen            time.Time
	Gain                 float32
	Grabs                uint
	Guid                 string
	Imdb                 uint
	InfoHash             string
	Link                 string
	MagnetUri            string
	MinimumRatio         float32
	MinimumSeedTime      uint
	Peers                uint
	PublishDate          time.Time
	RageID               uint
	Seeders              uint
	Size                 uint
	TMDb                 uint
	TVDBId               uint
	TVMazeId             uint
	TraktId              uint
	DoubanId             uint
	Title                string
	Tracker              string
	TrackerId            string
	TrackerType          string
	UploadVolumeFactor   float32
}

func New(url string, apiKey string) Jacketter {
	j := jackett.NewJackett(&jackett.Settings{
		ApiURL: url,
		ApiKey: apiKey,
	})

	return &Jackett{
		jackett: j,
	}
}

func (j *Jackett) Fetch(ctx context.Context, query string, categories ...uint) (res []Result, err error) {
	resp, err := j.jackett.Fetch(ctx, &jackett.FetchRequest{
		Categories: categories, //[]uint{101823},
		Query:      query,
	})
	if err != nil {
		return res, err
	}

	for i := range resp.Results {
		r := Result{
			BannerUrl:            resp.Results[i].BannerUrl,
			BlackholeLink:        resp.Results[i].BlackholeLink,
			Category:             resp.Results[i].Category,
			CategoryDesc:         resp.Results[i].CategoryDesc,
			Comments:             resp.Results[i].Comments,
			Description:          resp.Results[i].Description,
			Details:              resp.Results[i].Details,
			DownloadVolumeFactor: resp.Results[i].DownloadVolumeFactor,
			Files:                resp.Results[i].Files,
			FirstSeen:            resp.Results[i].FirstSeen.Time,
			Gain:                 resp.Results[i].Gain,
			Grabs:                resp.Results[i].Grabs,
			Guid:                 resp.Results[i].Guid,
			Imdb:                 resp.Results[i].Imdb,
			InfoHash:             resp.Results[i].InfoHash,
			Link:                 resp.Results[i].Link,
			MagnetUri:            resp.Results[i].MagnetUri,
			MinimumRatio:         resp.Results[i].MinimumRatio,
			MinimumSeedTime:      resp.Results[i].MinimumSeedTime,
			Peers:                resp.Results[i].Peers,
			PublishDate:          resp.Results[i].PublishDate.Time,
			RageID:               resp.Results[i].RageID,
			Seeders:              resp.Results[i].Seeders,
			Size:                 resp.Results[i].Size,
			TMDb:                 resp.Results[i].TMDb,
			TVDBId:               resp.Results[i].TVDBId,
			TVMazeId:             resp.Results[i].TVMazeId,
			TraktId:              resp.Results[i].TraktId,
			DoubanId:             resp.Results[i].DoubanId,
			Title:                resp.Results[i].Title,
			Tracker:              resp.Results[i].Tracker,
			TrackerId:            resp.Results[i].TrackerId,
			TrackerType:          resp.Results[i].TrackerType,
			UploadVolumeFactor:   resp.Results[i].UploadVolumeFactor,
		}

		res = append(res, r)
	}

	return res, nil
}
