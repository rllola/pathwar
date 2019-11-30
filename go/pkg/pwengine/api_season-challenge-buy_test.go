package pwengine

import (
	"context"
	"testing"

	"pathwar.land/go/internal/testutil"
	"pathwar.land/go/pkg/errcode"
)

func TestEngine_ChallengeBuy(t *testing.T) {
	engine, cleanup := TestingEngine(t, Opts{Logger: testutil.Logger(t)})
	defer cleanup()
	ctx := testingSetContextToken(context.Background(), t)

	solo := testingSoloSeason(t, engine)

	// fetch user session
	session, err := engine.UserGetSession(ctx, nil)
	checkErr(t, "", err)
	activeTeam := session.User.ActiveTeamMember.Team

	// fetch challenges
	challenges, err := engine.SeasonChallengeList(ctx, &SeasonChallengeList_Input{solo.ID})
	checkErr(t, "", err)

	var tests = []struct {
		name        string
		input       *SeasonChallengeBuy_Input
		expectedErr error
	}{
		{"nil", nil, errcode.ErrMissingInput},
		{"empty", &SeasonChallengeBuy_Input{}, errcode.ErrMissingInput},
		{"invalid season challenge ID", &SeasonChallengeBuy_Input{SeasonChallengeID: 42, TeamID: activeTeam.ID}, errcode.ErrInvalidSeason},
		{"invalid team ID", &SeasonChallengeBuy_Input{SeasonChallengeID: challenges.Items[0].ID, TeamID: 42}, errcode.ErrInvalidTeam},
		{"valid 1", &SeasonChallengeBuy_Input{SeasonChallengeID: challenges.Items[0].ID, TeamID: activeTeam.ID}, nil},
		{"valid 2 (duplicate)", &SeasonChallengeBuy_Input{SeasonChallengeID: challenges.Items[0].ID, TeamID: activeTeam.ID}, errcode.ErrChallengeAlreadySubscribed},
		// FIXME: check for a team and a challenge in different seasons
		// FIXME: check for a team from another user
		// FIXME: check for a challenge in draft mode
	}

	for _, test := range tests {
		subscription, err := engine.SeasonChallengeBuy(ctx, test.input)
		testSameErrcodes(t, test.name, test.expectedErr, err)
		if err != nil {
			continue
		}

		testSameInt64s(t, test.name, test.input.TeamID, subscription.ChallengeSubscription.TeamID)
		testSameInt64s(t, test.name, test.input.SeasonChallengeID, subscription.ChallengeSubscription.SeasonChallengeID)
		testSameInt64s(t, test.name, session.User.ID, subscription.ChallengeSubscription.BuyerID)

		// check if challenge subscription is now visible in season challenge list
		challenges, err := engine.SeasonChallengeList(ctx, &SeasonChallengeList_Input{solo.ID})
		checkErr(t, test.name, err)

		found := 0
		for _, challenge := range challenges.Items {
			if challenge.ID == subscription.ChallengeSubscription.SeasonChallengeID {
				found++
				if len(challenge.Subscriptions) != 1 {
					t.Errorf("%s: Expected only one subscription, got %d.", test.name, len(challenge.Subscriptions))
				}

				testSameInt64s(t, test.name, subscription.ChallengeSubscription.ID, challenge.Subscriptions[0].ID)
				testSameInt64s(t, test.name, test.input.TeamID, challenge.Subscriptions[0].TeamID)
			}
		}
		if found != 1 {
			t.Errorf("%s: Expected 1 found, got %d.", test.name, found)
		}
	}
}
