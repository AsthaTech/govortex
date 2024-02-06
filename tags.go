package govortex

import (
	"context"
	"fmt"
)

// Tags retrieves the list of tags from the Vortex API.
// It takes a context as input.
// It returns a TagsResponse and an error.
func (v *VortexApi) Tags(ctx context.Context) (*TagsResponse, error) {
	var resp TagsResponse
	_, err := v.doJson(ctx, "GET", URITags, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTag creates a new tag with the Vortex API.
// It takes a context and a TagRequest as input.
// It returns a TagResponse and an error.
func (v *VortexApi) CreateTag(ctx context.Context, request TagRequest) (*TagResponse, error) {
	var resp TagResponse
	_, err := v.doJson(ctx, "POST", URITags, request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTag updates an existing tag with the Vortex API.
// It takes a context, a tag ID, and a TagRequest as input.
// It returns a TagResponse and an error.
func (v *VortexApi) UpdateTag(ctx context.Context, tag_id int, request TagRequest) (*TagResponse, error) {
	var resp TagResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URITag, tag_id), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTag deletes an existing tag with the Vortex API.
// It takes a context and a tag ID as input.
// It returns a TagResponse and an error.
func (v *VortexApi) DeleteTag(ctx context.Context, tag_id int) (*TagResponse, error) {
	var resp TagResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URITag, tag_id), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
