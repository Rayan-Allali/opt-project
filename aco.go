package main

import (
	"fmt"
	"math/rand"
)

func smallestMotifIndex(motifs []string) int {
	smallestMotif := motifs[0]
	index := 0
	for i, motif := range motifs {
		if len(motif) < len(smallestMotif) {
			smallestMotif = motif
			index = i
		}
	}
	return index
}
func checkCommonSequenceExistenceInMotifs(sequence string, motifs []string, index int) bool {
	k := 0
	for i := 0; i < len(motifs); i++ {
		if i != index {
			k = 0
			for j := 0; j < len(motifs[i]); j++ {
				for l := 0; l < len(sequence); l++ {
					if motifs[i][j] == sequence[l] {
						k = k + 1 //  si au moins une lettre du petit motif existe dans ce motif on pass verifier les autres motifs
						break
					}
				}
			}
			if k == 0 {
				return false
			}
		}
	}
	return true
}
func checkSubSequenceInMotifs(sequence string, motifs []string) bool {
	k := 0
	for i := 0; i < len(motifs); i++ {
		k = 0
		for j := 0; j < len(motifs[i]); j++ {

			if motifs[i][j] == sequence[k] {
				k = k + 1
			}
			if k == len(sequence) {
				break
			}
		}
		if k < len(sequence) {
			return false
		}

	}
	return true
}

func evaporate(iteration int, pheromoneTable map[string]float64, evaporationRate float64) map[string]float64 {
	if iteration%10 == 0 {
		for i := range pheromoneTable {
			r := 1 - evaporationRate
			pheromoneTable[i] = r * pheromoneTable[i]
		}
	}
	return pheromoneTable
}
func RandomLetter(PheromoneTable map[string]float64) string {

	keys := make([]string, 0, len(PheromoneTable))

	cumulativePheromoneTable := make([]float64, 0, len(PheromoneTable))

	var cumulativeProbability float64
	for key, prob := range PheromoneTable {
		keys = append(keys, key)
		cumulativeProbability += prob
		cumulativePheromoneTable = append(cumulativePheromoneTable, cumulativeProbability)
	}

	randomNum := rand.Float64() * cumulativeProbability
	var selectedIndex int
	for i, cp := range cumulativePheromoneTable {
		if randomNum <= cp {
			selectedIndex = i
			break
		}
	}
	return keys[selectedIndex]
}

func initPheromoneTable(motif []string) map[string]float64 {
	pheromoneTable := make(map[string]float64)
	for i := 0; i < len(motif); i++ {
		if _, ok := pheromoneTable[string(motif[i])]; ok {
			pheromoneTable[string(motif[i])] += 0.1
		} else {
			pheromoneTable[string(motif[i])] = 0.1
		}
	}
	return pheromoneTable
}
func updatePheromoneTable(pheromoneTable map[string]float64, motif string, solutionQualite int, bestSolution int, pheromoneIncreaseRate float64, currentItr int, evaporationRate float64) map[string]float64 {
	for i := 0; i < len(motif); i++ {
		pheromoneTable[string(motif[i])] = pheromoneTable[string(motif[i])] + pheromoneTable[string(motif[i])]*float64(solutionQualite/bestSolution)*pheromoneIncreaseRate
	}
	pheromoneTable = evaporate(currentItr, pheromoneTable, evaporationRate)
	return pheromoneTable
}

func existanceLetters(smallestIndex string, motifs []string) []string {
	existeLetters := []string{}
	for i := 0; i < len(smallestIndex); i++ {
		if checkSubSequenceInMotifs(string(smallestIndex[i]), motifs) {
			existeLetters = append(existeLetters, string(smallestIndex[i]))
		}
	}
	return existeLetters
}
func getNbrAnts(existeLetters []string) int {
	if len(existeLetters) > 2 && len(existeLetters) < 5 {
		return 2
	} else if len(existeLetters) < 2 {
		return 1
	}
	return len(existeLetters) / 3
}
func aco(MaxIter int, motifs []string, Letters []string, nbrAnts int, pheromoneIncreaseRate float64, evaporationRate float64) string {
	bestSolution := ""
	longuest := 1
	phermoneTables := initPheromoneTable(Letters)
	for i := 0; i < nbrAnts; i++ {
		motif := ""
		j := 0
		for j < MaxIter {
			Locallonguest := longuest
			temp := motif
			lettre := RandomLetter(phermoneTables)
			fmt.Println("random letter is :", lettre)
			fmt.Println("phermoneTables  is :", phermoneTables)
			motif += lettre
			fmt.Println(motif)
			if checkSubSequenceInMotifs(motif, motifs) {
				Locallonguest = len(motif)
				if longuest <= Locallonguest {
					longuest = Locallonguest
					bestSolution = motif
					updatePheromoneTable(phermoneTables, lettre, len(bestSolution), len(bestSolution), pheromoneIncreaseRate, j, evaporationRate)
				} else {
					phermoneTables = updatePheromoneTable(phermoneTables, lettre, len(motif), len(bestSolution), pheromoneIncreaseRate, j, evaporationRate)
				}

			} else {
				motif = temp
			}
			j = j + 1
		}
	}
	return bestSolution
}

func main() {
	motifs := []string{"karmapns", "markpna", "mpnriknas", "mlkpina", "kabrpinsamasa"}
	index := smallestMotifIndex(motifs)
	check := checkCommonSequenceExistenceInMotifs(motifs[index], motifs, index)
	if !check {
		solutions := []string{}
		fmt.Println(solutions)
	} else {
		evaporationRate := 0.2
		pheromoneIncreaseRate := 0.8
		smallestIndex := motifs[index]
		maxItr := 4000
		Letters := existanceLetters(smallestIndex, motifs)
		nbrAnts := getNbrAnts(Letters)
		fmt.Println(nbrAnts)
		solution := aco(maxItr, motifs, Letters, nbrAnts, pheromoneIncreaseRate, evaporationRate)
		fmt.Println(solution, len(solution))

	}

}
