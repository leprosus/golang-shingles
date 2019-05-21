# Golang shingles algorithm implementation for english, french, norwegian, russian, spanish and swedish

## Settings and usage

```go
SetLanguage(English)
SetShinglesLength(3)

line1 := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
line2 := `Lorem is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

percent, err := Compare(line1, line2)
if err != nil {
	panic(err)
}
```

## List of all functions

* golang_shingles.SetLanguage - sets language to stem
* golang_shingles.SetShinglesLength - sets shingles length
* golang_shingles.Compare - computes percent of difference