package app_test

import (
	"testing"
	"time"

	appConsumer "github.com/cosmos/interchain-security/app/consumer-democracy"
	ibctesting "github.com/cosmos/interchain-security/legacy_ibc_testing/testing"
	icstestingutils "github.com/cosmos/interchain-security/testutil/ibc_testing"
	"github.com/stretchr/testify/require"
)

func TestDemocracyGovernanceWhitelistingKeys(t *testing.T) {
	chain := ibctesting.NewTestChain(t, ibctesting.NewCoordinator(t, 0),
		icstestingutils.DemocracyConsumerAppIniter, "test")

	app := chain.App.(*appConsumer.App)
	paramKeeper := app.ParamsKeeper

	app.ConsumerKeeper.SetLatestBlockTimeValsetUpdate(chain.GetContext(), time.Now().Add(-(time.Hour * 2)))

	for paramKey := range appConsumer.WhitelistedParams {
		ss, ok := paramKeeper.GetSubspace(paramKey.Subspace)
		require.True(t, ok, "Unknown subspace %s", paramKey.Subspace)
		hasKey := ss.Has(chain.GetContext(), []byte(paramKey.Key))
		require.True(t, hasKey, "Invalid key %s for subspace %s", paramKey.Key, paramKey.Subspace)
	}
}
