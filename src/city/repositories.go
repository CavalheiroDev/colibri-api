package city

import (
	"context"

	"github.com/colibri-project-io/colibri-sdk-go/pkg/base/types"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
)

const (
	listQuery = `
		SELECT id, name, state, created_at, updated_at
		FROM cities
		WHERE 1=1
		AND ($1 = '' OR LOWER(name) LIKE LOWER(CONCAT('%', $1, '%')))
		AND ($2 = '' OR LOWER(state) LIKE LOWER(CONCAT('%', $2, '%')))
	`
	createQuery = "INSERT INTO cities (name, state) VALUES ($1, $2) RETURNING id, name, state, created_at, updated_at"
	existsQuery = "SELECT COUNT(*) FROM cities WHERE name = $1 AND state = $2"
)

type CityDBRepository struct{}

func (r *CityDBRepository) List(ctx context.Context, params *CityListQueryParams) (CityPageResponse, error) {
	return sqlDB.NewPageQuery[CityResponse](
		ctx,
		types.NewPageRequest(params.Page, params.PageSize, []types.Sort{types.NewSort(types.ASC, "name")}),
		listQuery,
		params.Name,
		params.State,
	).Execute()
}

func (r *CityDBRepository) Create(ctx context.Context, dto *CityCreateRequest) (*CityResponse, error) {
	return sqlDB.NewQuery[CityResponse](ctx, createQuery, dto.Name, dto.State).One()
}

func (r *CityDBRepository) ExistsByUniqueKey(ctx context.Context, name string, uf string) (*bool, error) {
	return sqlDB.NewQuery[bool](ctx, existsQuery, name, uf).One()
}
