package strand

func ToRNA(dna string) string {
    rna := []rune{}
    for _, nucleotide := range dna {
        rna = append(rna, map[rune]rune{
            'G': 'C',
            'C': 'G',
            'T': 'A',
            'A': 'U',
        }[nucleotide])
    }

    return string(rna)
}
