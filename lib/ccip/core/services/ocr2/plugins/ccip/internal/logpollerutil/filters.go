package logpollerutil

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

func RegisterLpFilters(lp logpoller.LogPoller, filters []logpoller.Filter, qopts ...pg.QOpt) error {
	for _, lpFilter := range filters {
		if filterContainsZeroAddress(lpFilter.Addresses) {
			continue
		}
		if err := lp.RegisterFilter(lpFilter, qopts...); err != nil {
			return err
		}
	}
	return nil
}

func UnregisterLpFilters(lp logpoller.LogPoller, filters []logpoller.Filter, qopts ...pg.QOpt) error {
	for _, lpFilter := range filters {
		if filterContainsZeroAddress(lpFilter.Addresses) {
			continue
		}
		if err := lp.UnregisterFilter(lpFilter.Name, qopts...); err != nil {
			return err
		}
	}
	return nil
}

func FiltersDiff(filtersBefore, filtersNow []logpoller.Filter) (created, deleted []logpoller.Filter) {
	created = make([]logpoller.Filter, 0, len(filtersNow))
	deleted = make([]logpoller.Filter, 0, len(filtersBefore))

	for _, f := range filtersNow {
		if !containsFilter(filtersBefore, f) {
			created = append(created, f)
		}
	}

	for _, f := range filtersBefore {
		if !containsFilter(filtersNow, f) {
			deleted = append(deleted, f)
		}
	}

	return created, deleted
}

func containsFilter(filters []logpoller.Filter, f logpoller.Filter) bool {
	for _, existing := range filters {
		if existing.Name == f.Name {
			return true
		}
	}
	return false
}

func filterContainsZeroAddress(addrs []common.Address) bool {
	for _, addr := range addrs {
		if addr == utils.ZeroAddress {
			return true
		}
	}
	return false
}
