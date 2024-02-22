package govortex

import (
	"context"
	"testing"
)

func (ts *TestSuite) TestCreateTag(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := TagRequest{}
	resp, err := ts.VortexApiClient.CreateTag(ctx, request)
	if err != nil {
		t.Errorf("rror while creating tag. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while creating tag")
	}
}

func (ts *TestSuite) TestGetTags(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.Tags(ctx)
	if err != nil {
		t.Errorf("Error while fetching tags. %v", err)
		return
	}
	if len(resp.Data) == 0 {
		t.Errorf("Error while fetching tags. %s", "tags are empty")
	}
}

func (ts *TestSuite) TestUpdateTag(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.UpdateTag(ctx, 1, TagRequest{})
	if err != nil {
		t.Errorf("Error while updating tag. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while updating tag. %s", "status is not success")
	}
}

func (ts *TestSuite) TestDeleteTag(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	resp, err := ts.VortexApiClient.DeleteTag(ctx, 1)
	if err != nil {
		t.Errorf("Error while deleting tag. %v", err)
		return
	}
	if resp.Status != "success" {
		t.Errorf("Error while deleting tag. %s", "status is not success")
	}
}

func (ts *ErrorTestSuite) TestCreateTag(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	request := TagRequest{}
	_, err := ts.VortexApiClient.CreateTag(ctx, request)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestGetTags(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.Tags(ctx)
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestUpdateTag(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.UpdateTag(ctx, 1, TagRequest{})
	checkError429(t, err)
}

func (ts *ErrorTestSuite) TestDeleteTag(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, err := ts.VortexApiClient.DeleteTag(ctx, 1)
	checkError429(t, err)
}
