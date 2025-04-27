package protein

import "errors"

/*

| Codon              | Protein       |
| :----------------- | :------------ |
| AUG                | Methionine    |
| UUU, UUC           | Phenylalanine |
| UUA, UUG           | Leucine       |
| UCU, UCC, UCA, UCG | Serine        |
| UAU, UAC           | Tyrosine      |
| UGU, UGC           | Cysteine      |
| UGG                | Tryptophan    |
| UAA, UAG, UGA      | STOP          |

*/

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("Invalid Base")
var dict = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUG": "Leucine",
	"UUA": "Leucine",
	"UCG": "Serine",
	"UCA": "Serine",
	"UCC": "Serine",
	"UCU": "Serine",
	"UAC": "Tyrosine",
	"UAU": "Tyrosine",
	"UGC": "Cysteine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
	"UGA": "STOP",
	"UAG": "STOP",
	"UAA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	result := []string{}
	for i := 0; i < len(rna); i += 3 {
		r := rna[i : i+3]
		val, err := FromCodon(r)
		if err != nil {
			if err.Error() == "STOP" {
				break
			}
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	val, ok := dict[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if val == "STOP" {
		return "", ErrStop
	}
	return val, nil
}
