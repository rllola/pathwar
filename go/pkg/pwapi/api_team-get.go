package pwapi

import (
	"context"
	"math/rand"

	"pathwar.land/go/pkg/errcode"
	"pathwar.land/go/pkg/pwdb"
)

func (svc *service) TeamGet(ctx context.Context, in *TeamGet_Input) (*TeamGet_Output, error) {
	if in == nil || in.TeamID == 0 {
		return nil, errcode.ErrMissingInput
	}

	var item pwdb.Team
	err := svc.db.
		Preload("Season").
		Preload("Organization").
		Preload("Members").                // FIXME: only if member of the team or if admin
		Preload("ChallengeSubscriptions"). // FIXME: only if member of the team or if admin
		Preload("Achievements").
		Where(pwdb.Team{
			ID:             in.TeamID,
			DeletionStatus: pwdb.DeletionStatus_Active,
		}).
		First(&item).
		Error
	if err != nil {
		return nil, errcode.ErrGetTeam.Wrap(err)
	}

	ret := TeamGet_Output{Item: &item}

	// tmp: fake data
	// FIXME: use real data
	item.GoldMedals = int64(rand.Intn(3))
	item.SilverMedals = int64(rand.Intn(3))
	item.BronzeMedals = int64(rand.Intn(4))
	item.Score = int64(rand.Intn(100))
	item.Cash = int64(rand.Intn(100))
	item.NbAchievements = int64(rand.Intn(10))

	return &ret, nil
}
