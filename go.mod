module tutorials/introdemo

go 1.15

// This dependency of stellar/go no longer exists; use a forked version of the repo instead.
replace bitbucket.org/ww/goautoneg => github.com/adjust/goautoneg v0.0.0-20150426214442-d788f35a0315

require (
	github.com/kinecosystem/agora-common v0.68.0
	github.com/kinecosystem/kin-go v0.6.0
)
