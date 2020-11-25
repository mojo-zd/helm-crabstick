package get

import "sort"

type Category []string

func (g *getter) Category() Category {
	out := Category{}
	filter := make(map[string]bool)
	file, err := g.cache.LoadIndex()
	if err != nil {
		return out
	}
	for _, chartVersion := range file.Entries {
		for _, chart := range chartVersion {
			if chart.Annotations == nil {
				continue
			}

			category, ok := chart.Annotations[CategoryKey]
			if ok && !filter[category] {
				filter[category] = true
				out = append(out, category)
			}
		}
	}
	sort.Sort(out)
	return out
}

func (category Category) Len() int {
	return len(category)
}

func (category Category) Less(i, j int) bool {
	return category[i] < category[j]
}

func (category Category) Swap(i, j int) {
	category[i], category[j] = category[j], category[i]
}
