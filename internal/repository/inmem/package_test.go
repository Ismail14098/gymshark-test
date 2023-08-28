package inmem

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackageRepository(t *testing.T) {
	ctx := context.Background()
	repository := NewPackageRepository()

	pkg, err := repository.Store(ctx, 10000)
	require.NoError(t, err)
	fmt.Println(pkg.ID)

	_, err = repository.Store(ctx, 10000)
	require.Error(t, err)
	fmt.Println(err)

	pkg.Size = 5000
	err = repository.Update(ctx, pkg)
	require.Error(t, err)
	fmt.Println(err)

	pkg.Size = 4000
	err = repository.Update(ctx, pkg)
	require.NoError(t, err)

	packages := repository.GetAllOrderedBySize(ctx)
	start := packages[0].Size
	for i := 1; i < len(packages); i++ {
		require.LessOrEqual(t, start, packages[i].Size)
		start = packages[i].Size
	}
}
