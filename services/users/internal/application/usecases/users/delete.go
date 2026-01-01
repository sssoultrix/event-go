package users

import (
	"context"
)

func (uc *usersUseCase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
