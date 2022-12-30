Given a string, s, that represents a DNA sequence, and a number, k—return all the contiguous sequences (substrings) of length k that occur more than once in the string. The order of the returned subsequences does not matter. If no repeated subsequence is found, the function should return an empty array.

Note: The DNA sequence is composed of a series of nucleotides abbreviated as AA, CC, GG, and TT. For example, ACGAATTCCGACGAATTCCG is a DNA sequence. When studying DNA, it is useful to identify repeated sequences in it.

Constraints:
- 11 ≤≤ s.length ≤≤ 104104
- s[i] is either A, C, G, or T.

example:
1.  input:
    k: 4
    s: CCATATGTATGGATAT
    output:
    - ATAT
    - TATG
2.  input:
    k: 6
    s: AGAGCTCCAGAGCTTG
    output:
    - AGAGCT



test case:
"AAAAACCCCCAAAAACCCCCC" , 8
"GGGGGGGGGGGGGGGGGGGGGGGGG" , 12
"TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT" , 10
"AAAAAACCCCCCCAAAAAAAACCCCCCCTG" , 10
"ATATATATATATATAT" , 6