package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/spanner"
	"github.com/sodefrin/spanner-emulator-ci/models"
)

func TestModel(t *testing.T) {
	a := models.Mytable{
		A: spanner.NullInt64{Valid: true, Int64: 10},
		B: spanner.NullInt64{Valid: true, Int64: 20},
	}

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, fmt.Sprintf("projects/%s/instances/%s/databases/%s", os.Getenv("SPANNER_PROJECT_ID"), os.Getenv("SPANNER_INSTANCE_ID"), os.Getenv("SPANNER_DATABASE_ID")))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.Apply(ctx, []*spanner.Mutation{a.Insert(ctx)}); err != nil {
		t.Fatal(err)
	}
}
