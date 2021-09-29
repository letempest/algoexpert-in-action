package array

const (
	HOME_TEAM_WON = 1
)

// Big O: O(n) time | O(k) space, where n is the number of competitions,
// k is the number of teams
func TournamentWinner(competitions [][2]string, results []int) string {
	var currentBestTeam string
	scores := map[string]int{
		currentBestTeam: 0,
	}

	for i, competition := range competitions {
		var winner string
		homeTeam, awayTeam := competition[0], competition[1]
		if results[i] != HOME_TEAM_WON {
			winner = awayTeam
		} else {
			winner = homeTeam
		}
		updateScores(scores, winner, 3)
		if scores[winner] > scores[currentBestTeam] {
			currentBestTeam = winner
		}
	}
	return currentBestTeam
}

func updateScores(scores map[string]int, team string, points int) {
	if _, exists := scores[team]; !exists {
		scores[team] = 0
	}
	scores[team] += points
}
