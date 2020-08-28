// +build integration
// +build !unit

package models

import check "gopkg.in/check.v1"

func (s *ModeledApiSuite) TestGetWatchlist(c *check.C) {
	wl, err := s.ma.GetWatchlist(&WatchlistRequest{
		Arguments: Arguments{
			WatchlistID: "2491746",
			// WatchlistID: "2452785",
			AssetTypes: "Bond,Stock,StockIndex,CfdOnStock,CfdOnIndex,FxSpot,FxForwards,FxVanillaOption,FxKnockInOption,FxKnockOutOption,FxOneTouchOption,FxNoTouchOption",
			Index:      0,
			RowCount:   20,
		},
		RefreshRate: 500,
		Format:      "application/json",
		ContextID:   "3232",
		ReferenceID: "13",
	})
	c.Assert(err, check.IsNil)
	c.Assert(wl, check.NotNil)

}
