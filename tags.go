package govortex

import (
	"context"
	"fmt"
)

func (v *VortexApi) Tags(ctx context.Context) (*TagsResponse, error) {
	var resp TagsResponse
	_, err := v.doJson(ctx, "GET", URITags, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) CreateTag(ctx context.Context, request TagRequest) (*TagResponse, error) {
	var resp TagResponse
	_, err := v.doJson(ctx, "POST", URITags, request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) UpdateTag(ctx context.Context, tag_id int, request TagRequest) (*TagResponse, error) {
	var resp TagResponse
	_, err := v.doJson(ctx, "PUT", fmt.Sprintf(URITag, tag_id), request, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) DeleteTag(ctx context.Context, tag_id int) (*TagResponse, error) {
	var resp TagResponse
	_, err := v.doJson(ctx, "DELETE", fmt.Sprintf(URITag, tag_id), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
