package indexer

import (
	"context"
	"sort"

	"seedseek/internal/config"
	"seedseek/internal/entity"
	"seedseek/pkg/logger"

	"github.com/go-redis/redis"
	"github.com/wmw64/go-jackett"
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
	jackett *jackett.Jackett
	cache   *redis.Client
}

// New returns indexer service.
func New(log logger.Logger, cfg *config.Config, r *redis.Client) Indexer {
	j := jackett.NewJackett(&jackett.Settings{
		ApiURL: cfg.JackettURL,
		ApiKey: cfg.JackettApiKey,
	})

	return &indexer{
		log:     log,
		cfg:     cfg,
		jackett: j,
		cache:   r,
	}
}

func (idx *indexer) Fetch(ctx context.Context, name string, sorting string, limit ...int) (items []entity.Item, err error) {
	res, err := idx.jackett.Fetch(ctx, &jackett.FetchRequest{
		Categories: []uint{categoryVR},
		Query:      name,
	})
	if err != nil {
		idx.log.ErrorContext(ctx, "", err)
	}

	for _, r := range res.Results {
		item := entity.Item{
			BannerUrl:            r.BannerUrl,
			BlackholeLink:        r.BlackholeLink,
			Category:             r.Category,
			CategoryDesc:         r.CategoryDesc,
			Comments:             r.Comments,
			Description:          r.Description,
			Details:              r.Details,
			DownloadVolumeFactor: r.DownloadVolumeFactor,
			Files:                r.Files,
			FirstSeen:            r.FirstSeen.Time,
			Gain:                 r.Gain,
			Grabs:                r.Grabs,
			Guid:                 r.Guid,
			Imdb:                 r.Imdb,
			InfoHash:             r.InfoHash,
			Link:                 r.Link,
			MagnetUri:            r.MagnetUri,
			MinimumRatio:         r.MinimumRatio,
			MinimumSeedTime:      r.MinimumSeedTime,
			Peers:                r.Peers,
			PublishDate:          r.PublishDate.Time,
			RageID:               r.RageID,
			Seeders:              r.Seeders,
			Size:                 r.Size,
			TMDb:                 r.TMDb,
			TVDBId:               r.TVDBId,
			TVMazeId:             r.TVMazeId,
			TraktId:              r.TraktId,
			DoubanId:             r.DoubanId,
			Title:                r.Title,
			Tracker:              r.Tracker,
			TrackerId:            r.TrackerId,
			TrackerType:          r.TrackerType,
			UploadVolumeFactor:   r.UploadVolumeFactor,
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

	for i := range items {
		idx.log.InfoContext(ctx, items[i].Details)
		// new := idx.cache.SAdd(items[i].Details).Val()
		new, err := idx.cache.SAdd("postedResults", items[i].Details).Result()
		if err != nil {
			idx.log.ErrorContext(ctx, "", err)
		}

		if new == 0 {
			items[i].IsNew = true
		}
	}

	return items, nil
}
