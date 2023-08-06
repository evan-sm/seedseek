package indexer

import (
	"context"
	"sort"

	"seedseek/internal/config"
	"seedseek/internal/entity"
	"seedseek/pkg/jackett"
	"seedseek/pkg/logger"
)

const (
	categoryVR  = uint(101823) // VR
	SortingSeed = "seeders"
	SortingGain = "gain"
)

type Indexer interface {
	Fetch(ctx context.Context, name string, sorting string, limit ...int) (items []entity.Item, err error)
}

type indexer struct {
	log     logger.Logger
	cfg     *config.Config
	jackett jackett.Jacketter
}

// New returns indexer service.
func New(log logger.Logger, cfg *config.Config, jackett jackett.Jacketter) Indexer {
	return &indexer{
		log:     log,
		cfg:     cfg,
		jackett: jackett,
	}
}

func (idx *indexer) Fetch(ctx context.Context, name string, sorting string, limit ...int) (items []entity.Item, err error) {
	res, err := idx.jackett.Fetch(ctx, name, categoryVR)
	if err != nil {
		idx.log.ErrorContext(ctx, "", err)
	}

	for i := range res {
		item := entity.Item{
			BannerUrl:            res[i].BannerUrl,
			BlackholeLink:        res[i].BlackholeLink,
			Category:             res[i].Category,
			CategoryDesc:         res[i].CategoryDesc,
			Comments:             res[i].Comments,
			Description:          res[i].Description,
			Details:              res[i].Details,
			DownloadVolumeFactor: res[i].DownloadVolumeFactor,
			Files:                res[i].Files,
			FirstSeen:            res[i].FirstSeen,
			Gain:                 res[i].Gain,
			Grabs:                res[i].Grabs,
			Guid:                 res[i].Guid,
			Imdb:                 res[i].Imdb,
			InfoHash:             res[i].InfoHash,
			Link:                 res[i].Link,
			MagnetUri:            res[i].MagnetUri,
			MinimumRatio:         res[i].MinimumRatio,
			MinimumSeedTime:      res[i].MinimumSeedTime,
			Peers:                res[i].Peers,
			PublishDate:          res[i].PublishDate,
			RageID:               res[i].RageID,
			Seeders:              res[i].Seeders,
			Size:                 res[i].Size,
			TMDb:                 res[i].TMDb,
			TVDBId:               res[i].TVDBId,
			TVMazeId:             res[i].TVMazeId,
			TraktId:              res[i].TraktId,
			DoubanId:             res[i].DoubanId,
			Title:                res[i].Title,
			Tracker:              res[i].Tracker,
			TrackerId:            res[i].TrackerId,
			TrackerType:          res[i].TrackerType,
			UploadVolumeFactor:   res[i].UploadVolumeFactor,
		}

		items = append(items, item)
	}

	switch sorting {
	case SortingSeed:
		sort.Slice(items, func(i, j int) bool {
			return items[i].Seeders > items[j].Seeders
		})
	case SortingGain:
		sort.Slice(items, func(i, j int) bool {
			return items[i].Gain > items[j].Gain
		})

	}

	if len(limit) > 0 {
		if len(items) >= limit[0] {
			items = items[:limit[0]]
		}
	}

	return items, nil
}
