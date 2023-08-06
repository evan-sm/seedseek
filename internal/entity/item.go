package entity

import "time"

type Item struct {
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
