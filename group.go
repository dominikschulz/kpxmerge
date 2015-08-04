package main

type Group struct {
	Title   string  `xml:"title"`
	Icon    int     `xml:"icon"`
	Entries []Entry `xml:"entry"`
	Groups  []Group `xml:"group"`
}

func (g *Group) Merge(o Group) (int, error) {
	var mergedEntries int
	var mergedGroups int
	g.Entries, mergedEntries = mergeEntries(g.Entries, o.Entries)
	g.Groups, mergedGroups = mergeGroups(g.Groups, o.Groups)
	return mergedEntries + mergedGroups, nil
}

func (g *Group) Count() int {
	sum := 0
	for _, sg := range g.Groups {
		sum += sg.Count()
	}
	sum += len(g.Entries)
	return sum
}
