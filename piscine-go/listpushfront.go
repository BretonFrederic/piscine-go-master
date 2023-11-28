package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Tail == nil { // Si la queue de liste = null
		l.Head = n // Tête de liste = noeud(1) de type node
		l.Tail = n // Queue de liste = noeud(1) de type node
		return     // Evite la boucle infini
	}

	// l.Tail.Next = n // noeud suivant pointé par la queue de liste = prochain noeud de type node en mémoire
	// l.Tail = n      // noeud de queue de liste = noeud (courant)

	n.Next = l.Head // Le noeud suivant = pointeur de tête de liste
	l.Head = n      // noeud de queue de liste = noeud (courant)
}
