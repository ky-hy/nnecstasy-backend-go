package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ky-hy/nnecstasy-backend-go/graph/generated"
	"github.com/ky-hy/nnecstasy-backend-go/graph/model"
)

func (r *queryResolver) AdultVideos(ctx context.Context, filter *model.VideoFilter) ([]*model.Video, error) {
	raw, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &r.videos)
	return r.videos, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
