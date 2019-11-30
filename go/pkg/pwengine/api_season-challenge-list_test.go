package pwengine

import (
	"context"
	"testing"

	"pathwar.land/go/internal/testutil"
	"pathwar.land/go/pkg/errcode"
)

func TestEngine_SeasonChallengeList(t *testing.T) {
	engine, cleanup := TestingEngine(t, Opts{Logger: testutil.Logger(t)})
	defer cleanup()
	ctx := testingSetContextToken(context.Background(), t)

	// fetch user session to ensure account is created
	_, err := engine.UserGetSession(ctx, nil)
	checkErr(t, "", err)

	seasons := map[string]int64{}
	for _, season := range testingSeasons(t, engine).Items {
		seasons[season.Name] = season.ID
	}

	var tests = []struct {
		name          string
		input         *SeasonChallengeList_Input
		expectedErr   error
		expectedItems int
	}{
		{"empty", &SeasonChallengeList_Input{}, errcode.ErrMissingInput, 0},
		{"unknown-season-id", &SeasonChallengeList_Input{SeasonID: -42}, errcode.ErrInvalidSeasonID, 0},
		{"solo-mode", &SeasonChallengeList_Input{SeasonID: seasons["Solo Mode"]}, nil, 5},
		{"test-season", &SeasonChallengeList_Input{SeasonID: seasons["Test Season"]}, errcode.ErrUserHasNoTeamForSeason, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := engine.SeasonChallengeList(ctx, test.input)
			testSameErrcodes(t, "", test.expectedErr, err)
			if err != nil {
				return
			}

			//fmt.Println(godev.PrettyJSON(ret.Items))
			for _, item := range ret.Items {
				testSameInt64s(t, "", test.input.SeasonID, item.SeasonID)
			}
			if len(ret.Items) != test.expectedItems {
				t.Errorf("Expected %d, got %d.", test.expectedItems, len(ret.Items))
			}
		})
	}
}
