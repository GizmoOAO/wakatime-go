package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestProjects(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")
	apiBase := os.Getenv("WAKATIME_API_BASE")

	client := NewClient(apiKey, nil, apiBase)
	ctx := context.Background()
	query := &ProjectsQuery{}
	_, err := client.Projects.Current(ctx, query)
	if err != nil {
		t.Error(err)
	}
	_, err = client.Projects.User(ctx, userID, query)
	if err != nil {
		t.Error(err)
	}

}
