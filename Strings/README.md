# Go String Problems — Interview Prep (113 Problems)

All solutions are in `package strings_problems`. Each file is named `NN_description.go`.

---

## Problem List

| # | File | Function Signature | Description |
|---|------|-------------------|-------------|
| 1 | `01_string_length.go` | `StringLength(s string) (byteLen, charLen int)` | Calculate byte length vs rune (character) length |
| 2 | `02_char_frequency.go` | `CharFrequency(s string) map[rune]int` | Count frequency of each character |
| 3 | `03_first_last_two_chars.go` | `FirstLastTwo(s string) string` | Get string of first 2 and last 2 chars |
| 4 | `04_replace_first_char.go` | `ReplaceFirstChar(s string) string` | Replace all occurrences of first char (except itself) with `$` |
| 5 | `05_swap_first_two_chars.go` | `SwapFirstTwo(a, b string) string` | Swap first 2 chars of two strings |
| 6 | `06_add_ing_or_ly.go` | `AddIngOrLy(s string) string` | Add `ing` or `ly` suffix based on rules |
| 7 | `07_replace_not_poor.go` | `ReplaceNotPoor(s string) string` | Replace `not...poor` span with `good` |
| 8 | `08_longest_word.go` | `LongestWord(words []string) (string, int)` | Find longest word and its length |
| 9 | `09_remove_nth_char.go` | `RemoveNthChar(s string, n int) string` | Remove character at index n |
| 10 | `10_swap_first_last.go` | `SwapFirstLast(s string) string` | Swap first and last characters |
| 11 | `11_remove_odd_index_chars.go` | `RemoveOddIndexChars(s string) string` | Remove characters at odd indices |
| 12 | `12_count_word_occurrences.go` | `WordOccurrences(s string) map[string]int` | Count occurrences of each word |
| 13 | `13_upper_lower_case.go` | `UpperAndLower(s string) (upper, lower string)` | Return uppercase and lowercase versions |
| 14 | `14_sort_distinct_words.go` | `SortDistinctWords(s string) string` | Sort distinct words from comma-separated input |
| 15 | `15_wrap_html_tags.go` | `AddHTMLTags(tag, content string) string` | Wrap content in HTML tags |
| 16 | `16_insert_string_middle.go` | `InsertStringMiddle(outer, inner string) string` | Insert string into the middle of another |
| 17 | `17_repeat_last_two_chars.go` | `RepeatLastTwo(s string) string` | Repeat last 2 chars 4 times |
| 18 | `18_first_three_chars.go` | `FirstThree(s string) string` | Get first 3 characters |
| 19 | `19_substring_before_char.go` | `SubstringBeforeChar(s string, ch byte) string` | Get substring before last occurrence of a character |
| 20 | `20_reverse_if_multiple_of_4.go` | `ReverseIfMultipleOf4(s string) string` | Reverse string if length is multiple of 4 |
| 21 | `21_uppercase_if_two_upper_in_first_four.go` | `UppercaseIfTwoUpperInFirstFour(s string) string` | Uppercase if 2+ uppercase chars in first 4 |
| 22 | `22_sort_string_lexicographically.go` | `SortStringLex(s string) string` | Sort characters of string lexicographically |
| 23 | `23_remove_newline.go` | `RemoveNewline(s string) string` | Remove newline characters |
| 24 | `24_starts_with.go` | `StartsWith(s, prefix string) bool` | Check if string starts with prefix |
| 25 | `25_caesar_cipher.go` | `CaesarCipher(s string, shift int) string` | Caesar cipher encryption |
| 26 | `26_format_text_width50.go` | `FormatTextWidth(s string, width int) string` | Wrap text to specified width |
| 27 | `27_remove_indentation.go` | `RemoveIndentation(s string) string` | Remove indentation from all lines |
| 28 | `28_add_prefix_to_lines.go` | `AddPrefixToLines(s, prefix string) string` | Add prefix to every line |
| 29 | `29_set_first_line_indentation.go` | `IndentFirstLine(s, indent string) string` | Indent only the first line |
| 30 | `30_print_two_decimal_places.go` | `FormatTwoDecimals(f float64) string` | Format float to 2 decimal places |
| 31 | `31_print_with_sign.go` | `FormatWithSign(f float64) string` | Format float with explicit +/- sign |
| 32 | `32_print_no_decimal.go` | `FormatNoDecimal(f float64) string` | Format float with no decimal places |
| 33 | `33_left_pad_zeros.go` | `LeftPadZeros(n, width int) string` | Left-pad integer with zeros |
| 34 | `34_right_pad_stars.go` | `RightPadStars(n, width int) string` | Right-pad integer with `*` |
| 35 | `35_comma_separator.go` | `FormatWithCommas(n int) string` | Format integer with comma separators |
| 36 | `36_format_percentage.go` | `FormatPercentage(f float64) string` | Format float as percentage |
| 37 | `37_align_number.go` | `AlignNumber(n int) (left, right, center string)` | Left, right, center align in width 10 |
| 38 | `38_count_substring.go` | `CountSubstring(s, sub string) int` | Count occurrences of substring |
| 39 | `39_reverse_string.go` | `ReverseString(s string) string` | Reverse a string (unicode-safe) |
| 40 | `40_reverse_words.go` | `ReverseWords(s string) string` | Reverse words in a string |
| 41 | `41_strip_characters.go` | `StripChars(s, chars string) string` | Strip a set of characters from string |
| 42 | `42_count_repeated_chars.go` | `CountRepeatedChars(s string) map[rune]int` | Count characters that appear more than once |
| 43 | `43_area_volume_symbols.go` | `RectangleArea(l, w float64) string` / `CylinderVolume(r, h float64) string` | Area and volume with ² ³ symbols |
| 44 | `44_char_indices.go` | `CharIndices(s string) []struct{Ch rune; Index int}` | Get index of each character |
| 45 | `45_has_all_alphabet.go` | `HasAllAlphabet(s string) bool` | Check if string is a pangram |
| 46 | `46_string_to_words.go` | `StringToWords(s string) []string` | Convert string to slice of words |
| 47 | `47_lowercase_first_n.go` | `LowercaseFirstN(s string, n int) string` | Lowercase first n characters |
| 48 | `48_swap_commas_dots.go` | `SwapCommasAndDots(s string) string` | Swap commas and dots |
| 49 | `49_count_vowels.go` | `CountVowels(s string) map[rune]int` | Count and display vowels |
| 50 | `50_split_on_last_delimiter.go` | `SplitOnLastDelimiter(s string, delim rune) (string, string)` | Split on last occurrence of delimiter |
| 51 | `51_first_non_repeating_char.go` | `FirstNonRepeating(s string) rune` | Find first non-repeating character |
| 52 | `52_permutations_with_repetition.go` | `PermutationsWithRepetition(s string, r int) []string` | All permutations with repetition of length r |
| 53 | `53_first_repeated_char.go` | `FirstRepeatedChar(s string) rune` | Find first character that repeats |
| 54 | `54_first_repeated_char_smallest_index.go` | `FirstRepeatedCharSmallestIndex(s string) rune` | First repeated char with smallest first-occurrence index |
| 55 | `55_first_repeated_word.go` | `FirstRepeatedWord(s string) string` | Find first repeated word in sentence |
| 56 | `56_second_most_repeated_word.go` | `SecondMostRepeatedWord(s string) string` | Find second most repeated word |
| 57 | `57_remove_spaces.go` | `RemoveSpaces(s string) string` | Remove all spaces |
| 58 | `58_move_spaces_to_front.go` | `MoveSpacesToFront(s string) string` | Move all spaces to front |
| 59 | `59_max_occurring_char.go` | `MaxOccurringChar(s string) rune` | Find most frequent character |
| 60 | `60_capitalize_first_last.go` | `CapitalizeFirstLast(s string) string` | Capitalize first and last letter of each word |
| 61 | `61_remove_duplicate_chars.go` | `RemoveDuplicateChars(s string) string` | Remove duplicate chars (keep first occurrence) |
| 62 | `62_sum_digits.go` | `SumDigits(s string) int` | Sum all digits in string |
| 63 | `63_remove_leading_zeros_ip.go` | `RemoveLeadingZerosIP(ip string) string` | Remove leading zeros from IP address |
| 64 | `64_max_consecutive_zeros.go` | `MaxConsecutiveZeros(s string) int` | Max consecutive zeros in binary string |
| 65 | `65_common_chars.go` | `CommonChars(a, b string) string` | Common characters between two strings |
| 66 | `66_make_anagrams.go` | `MakeAnagrams(a, b string) string` | Characters needed to make two strings anagrams |
| 67 | `67_remove_consecutive_duplicates.go` | `RemoveConsecutiveDuplicates(s string) string` | Remove all consecutive duplicates |
| 68 | `68_separate_single_multiple_chars.go` | `SeparateSingleMultiple(s string) (single, multiple string)` | Separate single-occurrence and multi-occurrence chars |
| 69 | `69_longest_common_substring.go` | `LongestCommonSubstring(a, b string) string` | Find longest common substring |
| 70 | `70_concatenate_uncommon_chars.go` | `UncommonChars(a, b string) string` | Concatenate chars not common to both strings |
| 71 | `71_move_spaces_front_single_traversal.go` | `MoveSpacesFront(s string) string` | Move spaces to front in single traversal |
| 72 | `72_remove_all_except_char.go` | `RemoveAllExcept(s string, keep rune) string` | Remove all chars except a specified one |
| 73 | `73_count_char_types.go` | `CountCharTypes(s string) (upper, lower, special, digits int)` | Count uppercase, lowercase, special, numeric chars |
| 74 | `74_minimum_window_substring.go` | `MinWindowSubstring(str1, str2 string) string` | Minimum window containing all chars of another string |
| 75 | `75_smallest_window_all_chars.go` | `SmallestWindowAllChars(s string) string` | Smallest window containing all distinct chars |
| 76 | `76_substrings_k_distinct.go` | `SubstringsWithKDistinct(s string, k int) int` | Count substrings with exactly k distinct chars |
| 77 | `77_count_non_empty_substrings.go` | `CountNonEmptySubstrings(s string) int` | Count all non-empty substrings |
| 78 | `78_chars_at_alphabet_position.go` | `CharsAtAlphabetPosition(s string) int` | Count chars at same position as in alphabet |
| 79 | `79_smallest_largest_word.go` | `SmallestLargestWord(s string) (smallest, largest string)` | Find smallest and largest word |
| 80 | `80_substrings_same_first_last.go` | `SubstringsWithSameFirstLast(s string) int` | Count substrings with same first and last char |
| 81 | `81_index_of_substring.go` | `IndexOfSubstring(s, sub string) (int, string)` | Index of substring or "Not found" |
| 82 | `82_wrap_paragraph.go` | `WrapParagraph(s string, width int) string` | Wrap string into paragraph with given width |
| 83 | `83_number_in_bases.go` | `NumberInBases(n int) (decimal, octal, hex, binary string)` | Number in decimal, octal, hex, binary |
| 84 | `84_swap_cases.go` | `SwapCases(s string) string` | Swap upper/lower case of each character |
| 85 | `85_bytearray_to_hex.go` | `ByteArrayToHex(b []byte) string` | Convert byte slice to hex string |
| 86 | `86_delete_all_occurrences.go` | `DeleteAllOccurrences(s string, ch rune) string` | Delete all occurrences of a character |
| 87 | `87_common_values_two_strings.go` | `CommonValues(a, b string) string` | Find common leading/matching values in two strings |
| 88 | `88_validate_string.go` | `ValidateString(s string) []string` | Validate: has capital, lower, digit, min length |
| 89 | `89_remove_unwanted_chars.go` | `RemoveUnwantedChars(s string) string` | Remove non-alphanumeric, non-space characters |
| 90 | `90_remove_duplicate_words.go` | `RemoveDuplicateWords(s string) string` | Remove duplicate words (preserve order) |
| 91 | `91_heterogeneous_list_to_string.go` | `HeterogeneousListToString(items []interface{}) string` | Convert mixed-type slice to comma-separated string |
| 92 | `92_string_similarity.go` | `StringSimilarity(a, b string) float64` | Find similarity ratio between two strings |
| 93 | `93_extract_numbers.go` | `ExtractNumbers(s string) []int` | Extract all integers from a string |
| 94 | `94_hex_to_rgb.go` | `HexToRGB(hex string) (r, g, b int)` | Convert hex color code to RGB |
| 95 | `95_rgb_to_hex.go` | `RGBToHex(r, g, b int) string` | Convert RGB to hex color code |
| 96 | `96_to_camel_case.go` | `ToCamelCase(s string) string` | Convert string to camelCase |
| 97 | `97_to_snake_case.go` | `ToSnakeCase(s string) string` | Convert string to snake_case |
| 98 | `98_decapitalize_first.go` | `DecapitalizeFirst(s string) string` | Lowercase the first letter |
| 99 | `99_split_multiline.go` | `SplitMultiline(s string) []string` | Split multi-line string into slice of lines |
| 100 | `100_check_duplicate_chars_in_words.go` | `AllWordsUniqueChars(s string) bool` | Check if all words have unique characters |
| 101 | `101_add_strings_as_numbers.go` | `AddStringsAsNumbers(a, b string) string` | Add two strings as integer numbers |
| 102 | `102_remove_punctuation.go` | `RemovePunctuation(s string) string` | Remove all punctuation |
| 103 | `103_replace_long_words_hash.go` | `ReplaceLongWordsWithHash(s string, minLen int) string` | Replace long words with `#` per character |
| 104 | `104_capitalize_words.go` | `CapitalizeWords(s string) string` | Capitalize first letter of each word |
| 105 | `105_extract_name_from_email.go` | `ExtractNameFromEmail(email string) string` | Extract name from email address |
| 106 | `106_remove_repeated_consecutive.go` | `RemoveRepeatedConsecutive(s string) string` | Remove repeated consecutive characters |
| 107 | `107_count_three_letter_matches.go` | `CountThreeLetterMatches(a, b string) int` | Count matching 3-letter sequences at same index |
| 108 | `108_hash_around_non_vowels.go` | `HashAroundNonVowels(s string) string` | Add `-` around non-vowel characters |
| 109 | `109_count_leap_years.go` | `CountLeapYears(yearRange string) int` | Count leap years in a "YYYY-YYYY" range |
| 110 | `110_insert_space_before_capitals.go` | `InsertSpaceBeforeCapitals(s string) string` | Insert space before each capital letter |
| 111 | `111_replace_chars_with_position.go` | `ReplaceCharsWithPosition(s string) string` | Replace letters with their alphabet position numbers |
| 112 | `112_add_number_strings.go` | `AddNumberStrings(a, b string) string` | Add two very large numbers as strings |
| 113 | `113_sort_words_by_first_char.go` | `SortWordsByFirstChar(s string) string` | Sort words alphabetically by first character |

---

## Key Go Interview Concepts Used

### 1. `len(s)` vs `utf8.RuneCountInString(s)`
- `len(s)` → byte count (O(1))
- `utf8.RuneCountInString(s)` → character count (O(n))
- Classic interview trap: `len("café")` returns `5`, not `4`

### 2. Rune vs Byte
- Iterate bytes: `for i := 0; i < len(s); i++`
- Iterate characters (unicode-safe): `for i, ch := range s` — `ch` is a `rune`
- Convert: `[]rune(s)` to work with characters by index

### 3. `strings.Builder` for Efficient Concatenation
- Never use `+=` in a loop — it's O(n²) due to immutable strings
- Use `strings.Builder` (or `bytes.Buffer`) for O(n) assembly

### 4. Multiple Return Values
- Idiomatic Go: return `(result, error)` or `(val1, val2)`
- Named returns: `func f() (n int, err error)` — useful for clarity

### 5. `fmt.Sprintf` Formatting Verbs
| Verb | Meaning |
|------|---------|
| `%d` | decimal integer |
| `%05d` | zero-padded, width 5 |
| `%f` / `%.2f` | float / 2 decimal places |
| `%+.2f` | float with sign |
| `%o` | octal |
| `%x` / `%X` | hex lowercase / uppercase |
| `%b` | binary |
| `%-10d` | left-aligned, width 10 |
| `%10d` | right-aligned, width 10 |

### 6. `sort` Package
- `sort.Slice(slice, less func)` — sort with custom comparator
- `sort.Strings([]string)` — sort string slice in place
- `sort.SliceStable` — preserves order of equal elements

### 7. `strings` Package Essentials
```go
strings.Contains(s, sub)
strings.HasPrefix(s, prefix) / strings.HasSuffix(s, suffix)
strings.Index(s, sub) / strings.LastIndex(s, sub)
strings.Split(s, sep) / strings.Join(slice, sep)
strings.TrimSpace(s) / strings.Trim(s, cutset)
strings.Replace(s, old, new, n)  // n=-1 means all
strings.ToUpper(s) / strings.ToLower(s)
strings.Fields(s)  // split on any whitespace
```

### 8. Sliding Window Pattern (Problems 74, 75, 76)
- Used for minimum/maximum window substring problems
- Two pointers + frequency map = O(n) solution
- Common in FAANG interviews

### 9. Map for Frequency Counting (Problems 2, 12, 42, 49, 55, 56)
```go
freq := make(map[rune]int)
for _, ch := range s {
    freq[ch]++
}
```

### 10. `unicode` Package
```go
unicode.IsUpper(ch)  unicode.IsLower(ch)
unicode.IsDigit(ch)  unicode.IsLetter(ch)
unicode.IsPunct(ch)  unicode.IsSpace(ch)
unicode.ToUpper(ch)  unicode.ToLower(ch)
```
