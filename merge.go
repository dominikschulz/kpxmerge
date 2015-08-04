package main

import "log"

func mergeEntries(a, b []Entry) ([]Entry, int) {
	ents := make(map[string]Entry)
	merged := 0
	for _, e := range a {
		if orig, found := ents[e.Ident()]; found {
			log.Printf("WARNING: Self-Merging non-unique entry: %s: %v with %v", e.Ident(), ents[e.Ident()], e)
			changes, err := e.Merge(orig)
			if err != nil {
				log.Printf("Failed to merge %v and %v: %s", orig, e, err)
			} else {
				merged += changes
			}
		}
		ents[e.Ident()] = e
	}
	for _, e := range b {
		if orig, found := ents[e.Ident()]; found {
			changes, err := orig.Merge(e)
			if err != nil {
				log.Printf("Failed to merge %v and %v: %s\n", orig, e, err)
			} else {
				merged += changes
			}
			ents[e.Ident()] = orig
		} else {
			ents[e.Ident()] = e
		}
	}
	es := make([]Entry, 0, len(ents))
	for _, e := range ents {
		es = append(es, e)
	}
	return es, merged
}

func mergeGroups(a, b []Group) ([]Group, int) {
	groups := make(map[string]Group)
	// populate map
	for _, g := range a {
		if _, found := groups[g.Title]; found {
			log.Printf("WARNING: Overwriting group %s", g.Title)
		}
		groups[g.Title] = g
	}
	merged := 0
	// add uniques and merge duplicates
	for _, g := range b {
		if orig, found := groups[g.Title]; found {
			changes, err := orig.Merge(g)
			if err != nil {
				log.Printf("Failed to merge %v and %v: %s\n", orig, g, err)
			} else {
				merged += changes
			}
			groups[g.Title] = orig
		} else {
			groups[g.Title] = g
		}
	}
	gs := make([]Group, 0, len(groups))
	for _, g := range groups {
		gs = append(gs, g)
	}
	return gs, merged
}
