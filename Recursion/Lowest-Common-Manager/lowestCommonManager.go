// similar to [Graph]/Youngest Common Ancestor question

package lcm

type Hierarchy struct {
	Name    string
	Reports []*Hierarchy
}

func (h *Hierarchy) AddReports(reports ...string) {
	for _, report := range reports {
		h.Reports = append(h.Reports, &Hierarchy{
			Name: report,
		})
	}
}

type orgInfo struct {
	lowestCommonManager *Hierarchy
	numImportantReports int
}

// Big O: O(n) time | O(d) space, where d is the depth of tree-like structure Hierarchy
// at most d recursive calls waiting on the callstack
func LowestCommonManager(topManager Hierarchy, reportOne, reportTwo string) *Hierarchy {
	return getOrgInfo(topManager, reportOne, reportTwo).lowestCommonManager
}

func getOrgInfo(manager Hierarchy, reportOne, reportTwo string) orgInfo {
	numImportantReports := 0
	for _, directReport := range manager.Reports {
		orgInfo := getOrgInfo(*directReport, reportOne, reportTwo)
		if orgInfo.lowestCommonManager != nil {
			return orgInfo
		}
		numImportantReports += orgInfo.numImportantReports
	}
	if manager.Name == reportOne || manager.Name == reportTwo {
		numImportantReports += 1
	}
	var lowestCommonManager *Hierarchy
	if numImportantReports == 2 {
		lowestCommonManager = &manager
	}
	return orgInfo{
		lowestCommonManager: lowestCommonManager,
		numImportantReports: numImportantReports,
	}
}
