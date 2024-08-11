package sublist

// Relation type is defined in relations.go file.

func equal(l1, l2 []int) bool {
    if len(l1) != len(l2) {
        return false
    }

    for i, n := range l1 {
        if n != l2[i] {
            return false
        }
    }

    return true
}

func sublist(l1, l2 []int) bool {
    for i := range l2 {
        sublen := i + len(l1)

        if len(l2) >= sublen {
            if equal(l1, l2[i:sublen]) {
                return true
            }
        } else {
            break
        }
    }

    return false
}

func Sublist(l1, l2 []int) Relation {
    if equal(l1, l2) {
        return RelationEqual
    } else if sublist(l1, l2) {
        return RelationSublist
    } else if sublist(l2, l1) {
        return RelationSuperlist
    } else {
        return RelationUnequal
    }
}
