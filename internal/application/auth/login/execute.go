package login

import (
	"authsvc/internal/core/domainerr"
	"authsvc/internal/domain/session"
	"authsvc/internal/domain/token"
	"authsvc/internal/domain/user"
	"context"
	"time"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	u, g, err := i.loginAndGrantTokens(ctx, p)
	if err != nil {
		return nil, err
	}

	err = i.txManager.Do(ctx, func(ctx context.Context) error {
		err := i.refreshRepo.Save(ctx, session.New(g.Refresh.ID, u.ID, g.Refresh.ExpiresAt))
		if err != nil {
			return errors.Wrap(err, "failed to persist refresh token")
		}

		err = i.userRepo.UpdateLastLoginAttempt(ctx, u.ID, time.Now())
		if err != nil {
			return errors.Wrap(err, "faied to update last login attempt")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to run transaction")
	}

	return g, nil
}

func (i *implementation) loginAndGrantTokens(ctx context.Context, p *Payload) (*user.User, *token.GrantResult, error) {
	u, err := i.userRepo.FindByUsername(ctx, p.Username)
	if err != nil {
		return nil, nil, domainerr.Join(ErrAccountDoesNotExist, err)
	}

	if correctPassword := i.ph.Compare(p.Password, u.Password); !correctPassword {
		return nil, nil, errors.Wrap(ErrIncorrectPassword, "incorrect password received")
	}

	g, err := i.tokenProvider.Grant(ctx, u.ID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to grant tokens")
	}

	return u, g, nil
}
