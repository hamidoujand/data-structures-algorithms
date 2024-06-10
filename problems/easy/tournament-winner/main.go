package main

import "fmt"

const (
	homeTeamWon = 1
)

func main() {
	tournament := [][]string{
		{"HTML", "C#"},
		{"C#", "Go"},
		{"Go", "HTML"},
	}
	results := []int{0, 0, 1}
	fmt.Println(tournamentWinner(tournament, results))
}

// O(n) Time | O(k) Space, "K" number of teams in the competition
func tournamentWinner(competitions [][]string, results []int) string {
	var currentBestTeam string
	scores := map[string]int{
		currentBestTeam: 0,
	}

	for i, comp := range competitions {
		result := results[i]
		homeTeam, awayTeam := comp[0], comp[1]
		var winner string
		if result == homeTeamWon {
			winner = homeTeam
		} else {
			winner = awayTeam
		}
		scores[winner] += 3 //3 points added to the winner
		if scores[winner] > scores[currentBestTeam] {
			currentBestTeam = winner
		}
	}
	return currentBestTeam
}
