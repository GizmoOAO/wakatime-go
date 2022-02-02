package wakatime

import (
	"context"
	"os"
	"testing"
)

func TestUsers(t *testing.T) {
	apiKey := os.Getenv("WAKATIME_API_KEY")
	userID := os.Getenv("WAKATIME_USER_ID")
	apiBase := os.Getenv("WAKATIME_API_BASE")

	client := NewClient(apiKey, nil, apiBase)
	ctx := context.Background()
	query1 := &UsersQuery{}
	_, err := client.Users.Current(ctx, query1)
	if err != nil {
		t.Error(err)
	}
	_, err = client.Users.User(ctx, userID, query1)
	if err != nil {
		t.Error(err)
	}

}
