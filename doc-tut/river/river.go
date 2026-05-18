package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

type SortArgs struct {
	// Strings is a slice of strings to sort.
	Strings []string `json:"strings"`
	Numbers []int    `json:"numbers"`
}

func (SortArgs) Kind() string { return "sort" }

type SortWorker struct {
	// An embedded WorkerDefaults sets up default to fulfill the rest of
	// the Worker interface:
	river.WorkerDefaults[SortArgs]
}

func (w *SortWorker) Work(ctx context.Context, job *river.Job[SortArgs]) error {
	sort.Strings(job.Args.Strings)
	sort.Ints(job.Args.Numbers)
	fmt.Printf("Sorted strings: %+v\n", job.Args.Strings)
	fmt.Printf("Sorted numbers: %+v\n", job.Args.Numbers)
	return nil
}

func main() {
	workers := river.NewWorkers()

	// AddWorker panics if the worker is already registered or invalid
	river.AddWorker(workers, &SortWorker{})

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {
				MaxWorkers: 100,
			},
		},
		Workers: workers,
	})
	if err != nil {
		panic(err)
	}

	// run the client inline. All executed jobs will inherit from ctx
	if err := riverClient.Start(ctx); err != nil {
		panic(err)
	}

	// stop fetching new work and wait for active jobs to finish
	if err := riverClient.Stop(ctx); err != nil {
		panic(err)
	}

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback(ctx)

	// with transaction
	_, err = riverClient.InsertTx(ctx, tx, SortArgs{
		Strings: []string{
			"whale", "tiger", "bear",
		},
	}, nil)
	if err != nil {
		panic(err)
	}

	// without transaction
	_, err = riverClient.Insert(ctx, SortArgs{
		Strings: []string{
			"whale", "tiger", "bear",
		},
	}, nil)
	if err != nil {
		panic(err)
	}
}
