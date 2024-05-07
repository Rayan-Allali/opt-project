package main

import (
	"fmt"
)

// declaration d'un noead dune arbre
type Node struct {
	sequence string
	Children []*Node
}

// insertion d'un noead dune arbre
func (tree *Node) insert(motif string) *Node {
	node := &Node{sequence: motif}
	tree.Children = append(tree.Children, node)
	return node
}

// Fonction pour trouver l'index du motif le plus petit

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

// Fonction pour vérifier si le motif a déjà été traité (motif existe dans "closed")

func checkMotifInList(motif string, closed []string) bool {
	for _, s := range closed {
		if s == motif {
			return true
		}
	}
	return false
}

// Fonction pour vérifier si le motif  est une sous-séquence de tous les autres motifs sauf le motif index
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

// Fonction pour vérifier si au moins une lettre du petit motif existe dans les autres motifs, car s'il n'existe pas, il n'est pas nécessaire de vérifier
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

// Fonction bfs ici on va deroule le bfs algorithme
func bfs(tree *Node, motifs []string, index int) []string {
	longuest := 0
	open := []*Node{}
	open = append(open, tree)
	closed := []string{}
	solutions := []string{}
	find := false
	find = checkSubSequenceInMotifs(motifs[index], motifs, index) //verifer c'est le motif le plus petit est une sous sequence de tous les autres motifs
	if find {
		solutions = append(solutions, motifs[index])
		return solutions
	}
	for len(open) > 0 {
		longuestLocal := longuest
		motif := open[0]
		open = open[1:] // on retire le premier element de la liste open  // like file do
		fmt.Println("motif")
		fmt.Println(motif)
		fmt.Println("motif/open")
		for i := 0; i < len(motif.sequence); i++ {
			if len(motif.sequence) > 1 {
				subMotif := motif.sequence[:i] + motif.sequence[i+1:] // Supprime une lettre de la séquence actuelle et crée une nouvelle sous-chaîne
				fmt.Println(subMotif)
				if !checkMotifInList(subMotif, closed) { // verifier c'est la sous sequence existe pas dans closed
					if checkSubSequenceInMotifs(subMotif, motifs, index) {
						longuestLocal = len(subMotif)
						if longuestLocal < longuest { // verifier c'est la taille de sous sequence inferieur a la taille de la longuest sous sequence on s'arrete
							break
						}
						longuest = len(subMotif)                // update la valeur de longuest sous sequence
						solutions = append(solutions, subMotif) // ajouter la sous sequence dans la liste des solutions
					}
					open = append(open, tree.insert(subMotif)) // ajouter la sequence dans la file pour apres la manipuler
					closed = append(closed, subMotif)          // ajouter la sequence a closed
				} else {
					continue
				}
			}
		}
		if longuestLocal < longuest { // verifier c'est la taille de sous sequence inferieur a la taille de la longuest sous sequence on s'arrete
			break
		}
	}
	return solutions
}
func main() {
	motifs := []string{"CCCTGAGACA", "CTCCCATAACCTGTATTTCG"}
	index := smallestMotifIndex(motifs)

	check := checkCommonSequenceExistenceInMotifs(motifs[index], motifs, index)
	fmt.Println(check)
	if !check {
		solutions := []string{}
		fmt.Println(solutions)
	} else {
		tree := &Node{sequence: motifs[index]}
		solutions := bfs(tree, motifs, index)
		fmt.Println(solutions)
	}
}