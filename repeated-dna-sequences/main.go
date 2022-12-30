package dna

// O(n^2)
func checkDNAString(k int, s string) (result []string) {
	for i := 0; i < len(s); i++ {
		possibleDna := string(s[i : i+k])
		if i+k >= len(s) {
			break
		}

		for j := i + 1; j < len(s); j++ {
			if j+k > len(s) {
				break
			}

			compareDna := string(s[j : j+k])
			if possibleDna == compareDna {
				result = append(result, possibleDna)
			}
		}
	}

	return result
}

// O(n)
func checkDNAString2(k int, s string) (result []string) {
	seen := map[string]string{}

	for i := 0; i < len(s); i++ {
		if i+k > len(s) {
			break
		}
		possibleDna := string(s[i : i+k])

		if seen[possibleDna] != "" {
			result = append(result, possibleDna)
			continue
		}
		seen[possibleDna] = possibleDna
	}

	return
}
