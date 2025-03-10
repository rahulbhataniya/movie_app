package rating
import(
	"context"
	"errors"
	"movieexample.com/rating/internal/repository"
	"movieexample.com/rating/pkg/model"
)

//ErrNotFound is returned when no rating are found for a record

var ErrNotFound=errors.New("rating not found for a record")
//record

type ratingRepository interface{
	Get(ctx context.Context,recordID model.RecordID,recordTye model.RecordType)([]model.Rating,error)
	Put(ctx context.Context,recordID model.RecordID,recordType model.RecordType,rating *model.Rating) error
}

//controller define a string rating service controller.

type Controller struct{
	repo ratingRepository
}

//New create a rating service controller

func New(repo ratingRepository) *Controller{
	return &Controller{repo}
}

//GetAggregatedRating returns the aggregated rating a record or ErrNotFound if there are no rating for it

func (c *Controller) GetAggregatedRating(ctx context.Context,recordID model.RecordID,recordType model.RecordType) (float64,error){
	ratings, err := c.repo.Get(ctx, recordID, recordType)
    if err != nil && err == repository.ErrNotFound {
        return 0, ErrNotFound
    } else if err != nil {
        return 0, err
    }
    sum := float64(0)
    for _, r := range ratings {
        sum += float64(r.Value)
    }
    return sum / float64(len(ratings)), nil
}

// PutRating writes a rating for a given record.
func (c *Controller) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
    return c.repo.Put(ctx, recordID, recordType, rating)
}