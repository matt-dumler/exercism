package protein

import (
    "errors"
    "regexp"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

var proteinFromCodon = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine", "UUC": "Phenylalanine",
    "UUA": "Leucine", "UUG": "Leucine",
    "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
    "UAU": "Tyrosine", "UAC": "Tyrosine",
    "UGU": "Cysteine", "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
    proteins := make([]string, 0)

    re := regexp.MustCompile(`[ACGU]{3}`)
    for _, codon := range re.FindAllString(rna, -1) {
        protein, err := FromCodon(codon)
        if err != nil {
            if err == ErrInvalidBase {
                return nil, ErrInvalidBase
            } else if err == ErrStop {
                break
            } else {
                return nil, err
            }
        }

        proteins = append(proteins, protein)
    }

    return proteins, nil
}

func FromCodon(codon string) (string, error) {
    protein, ok := proteinFromCodon[codon]
    if !ok {
        return "", ErrInvalidBase
    }

    if protein == "STOP" {
        return "", ErrStop
    }

    return protein, nil
}
