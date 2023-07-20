package main

import "fmt"

const HOME_TEAM_WON = 1

var matchStats = [][]string{
	{"HTML", "C#"},
	{"C#", "Python"},
	{"Python", "HTML"},
}

var results = []int{0, 1, 0}

func tournamenWinner(matchStats [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{"currentBestTeam": 0}

	for idx, competition := range matchStats {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)
		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
			fmt.Println(currentBestTeam)
		}
	}
	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
	fmt.Println(scores)
}

func main() {
	res := tournamenWinner(matchStats, results)
	fmt.Println(res)
}
