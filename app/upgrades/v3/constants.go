package v3

import (
	"github.com/Crypto-Dungeon/dungeonchain/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v3"
)

var Upgrade = upgrades.Fork{
	UpgradeName:    UpgradeName,
	UpgradeHeight:  10,
	BeginForkLogic: RunForkLogic,
}
