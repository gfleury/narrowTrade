package models

import check "gopkg.in/check.v1"

func (s *ModeledApiSuite) TestGetWatchlist(c *check.C) {
	wl, err := s.ma.GetWatchlist(&WatchlistRequest{
		Arguments: Arguments{
			WatchlistID: "2491746",
			AssetTypes:  "Bond,Stock,StockIndex,CfdOnStock,CfdOnIndex,FxSpot,FxForwards,FxVanillaOption,FxKnockInOption,FxKnockOutOption,FxOneTouchOption,FxNoTouchOption",
			Index:       0,
			RowCount:    20,
		},
		RefreshRate: 500,
		Format:      "application/json",
		ContextID:   "7865091208",
		ReferenceID: "10",
	})
	c.Assert(err, check.IsNil)
	c.Assert(wl, check.NotNil)

}
