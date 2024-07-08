#include "rna_transcription.h"

// to_rna: converts a DNA string of characters to an RNA string.
char *to_rna(const char *dna)
{
    const int length = strlen(dna) + 1; // length including string terminator

    char *rna = malloc(length);
    for (int i = 0; i < length; i++) {
        switch (dna[i]) {
            case 'G':
                rna[i] = 'C';
                break;
            case 'C':
                rna[i] = 'G';
                break;
            case 'T':
                rna[i] = 'A';
                break;
            case 'A':
                rna[i] = 'U';
                break;
            case '\0':
                rna[i] = '\0';
                break;
            default:
                free(rna);
                return NULL;
        }
    }

    return rna;
}
