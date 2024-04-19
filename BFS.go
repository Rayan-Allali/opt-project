package main

import (
	"fmt"
)

type Node struct {
	sequence string
	Children []*Node
}

func (tree *Node) insert(motif string) *Node {
	node := &Node{sequence: motif}
	tree.Children = append(tree.Children, node)
	return node
}

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

func checkMotifInList(motif string, closed []string) bool {
	for _, s := range closed {
		if s == motif {
			return true
		}
	}
	return false
}

func checkSubSequenceInMotifs(sequence string, motifs []string, index int) bool {
	k := 0
	for i := 0; i < len(motifs); i++ {
		if i != index {
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
	}
	return true
}

func checkCommonSequenceExistenceInMotifs(sequence string, motifs []string, index int) bool {
	k := 0
	for i := 0; i < len(motifs); i++ {
		if i != index {
			k = 0
			for j := 0; j < len(motifs[i]); j++ {
				for l := 0; l < len(sequence); l++ {
					if motifs[i][j] == sequence[l] {
						k = k + 1
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
func bfs(tree *Node, motifs []string, index int) []string {
	open := []*Node{}
	open = append(open, tree)
	closed := []string{}
	solutions := []string{}
	find := false
	find = checkSubSequenceInMotifs(motifs[index], motifs, index)
	if find {
		solutions = append(solutions, motifs[index])
		return solutions
	}
	for len(open) > 0 {
		motif := open[0]
		open = open[1:]
		fmt.Println("motif")
		fmt.Println(motif)
		fmt.Println("motif")
		for i := 0; i < len(motif.sequence); i++ {
			subMotif := motif.sequence[:i] + motif.sequence[i+1:]
			fmt.Println(subMotif)
			if !checkMotifInList(subMotif, closed) {
				if checkSubSequenceInMotifs(subMotif, motifs, index) {
					find = true
					solutions = append(solutions, subMotif)
				}
				open = append(open, tree.insert(subMotif))
				closed = append(closed, subMotif)
			} else {
				continue
			}
		}
		if find {
			break
		}
	}
	return solutions
}
func main() {
	motifs := []string{"le test motif est un motif test qui test le test motif qui lui meme test le motif test", "le motif est un motif test qui test le test motif qui lui meme test", "le lion testera motif qui lui meme test le motif test"}
	index := smallestMotifIndex(motifs)

	check := checkCommonSequenceExistenceInMotifs(motifs[index], motifs, index)
	fmt.Println(check)
	if !check {
		fmt.Println("there is no common sub string")
	} else {
		tree := &Node{sequence: motifs[index]}
		solutions := bfs(tree, motifs, index)
		fmt.Println(solutions)
	}

}
