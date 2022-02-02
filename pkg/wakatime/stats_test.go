package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestStats(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")
	apiBase := os.Getenv("WAKATIME_API_BASE")

	client := NewClient(apiKey, nil, apiBase)
	ctx := context.Background()
	query := &StatsQuery{}
	_, err := client.Stats.Current(ctx, RangeLast7Days, query)
	if err != nil {
		t.Error(err)
	}

	_, err = client.Stats.User(ctx, userID, RangeLast7Days, query)
	if err != nil {
		t.Error(err)
	}

}
