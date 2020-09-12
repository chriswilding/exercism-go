package protein

import "errors"

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	output := make([]string, 0, len(rna)/3)
	for i := 0; i < len(rna); i += 3 {
		aminoAcid, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			return output, nil
		}
		if err == ErrInvalidBase {
			return output, ErrInvalidBase
		}
		output = append(output, aminoAcid)
	}
	return output, nil
}
