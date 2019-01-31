# Algorithms and Data Structures
  
A personal review and overview of important algorithms and the utilization of basic data structures

## KMP (Knuth Morris Pratt Searching Algorithm)

This algorithm is used to find patterns inside strings of text
text = "AARDBCDGH"  
pattern = "BCD"

### LPS (Longest Prefix that is also a Suffix)
* The LPS array must be calculated first
* The LPS array is created using the pattern  
`String text = "abcdabca"`  
`lps = [0, 0, 0, 1, 1, 2, 3, 1]`

### LPS Complexity
* Building the LPS array: O(n)

### Searching the text
* Time: O(n + m) where n is the length of the pattern and m is the length of the text  
* Space: O(n) where n is the length of the pattern