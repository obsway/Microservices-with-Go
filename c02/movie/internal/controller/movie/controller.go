package movie

import (
	"context"
	"errors"

	metadatamodel "movieexample.com/metadata/pkg"
	"movieexample.com/movie/internal/gateway"
	model "movieexample.com/movie/pkg"
	ratingmodel "movieexample.com/rating/pkg"
)

// ErrNotFound is returned when the movie metadata
// is not found
var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregrateRating(ctx context.Context, recordID ratingmodel.RecordID,
		recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID,
		recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Controller defines a movie service controller.
type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new movie service controller.
func New(ratingGateway ratingGateway, memetadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, memetadataGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}

	rating, err := c.ratingGateway.GetAggregrateRating(ctx,
		ratingmodel.RecordID(id),
		ratingmodel.RecordTypeMovie)

	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just process in this case.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
