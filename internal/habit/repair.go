package habit

import "sort"

type RepairResult struct {
	Changed           bool
	DuplicateRemoved  int
	CompletionsSorted int
}

func Repair(store Store) (Store, RepairResult) {
	var result RepairResult
	for i, h := range store.Habits {
		seen := make(map[string]bool, len(h.Completions))
		completions := make([]string, 0, len(h.Completions))
		for _, completion := range h.Completions {
			if seen[completion] {
				result.DuplicateRemoved++
				continue
			}
			seen[completion] = true
			completions = append(completions, completion)
		}
		sorted := false
		if !sort.StringsAreSorted(completions) {
			sort.Strings(completions)
			result.CompletionsSorted++
			sorted = true
		}
		if len(completions) != len(h.Completions) || sorted {
			h.Completions = completions
			store.Habits[i] = h
			result.Changed = true
		}
	}
	return store, result
}
