module github.com/evolbioinf/fur

go 1.22.5

require (
	github.com/evolbioinf/clio v0.0.0-20240827074707-cb9ff755a85b
	github.com/evolbioinf/esa v0.0.0-20240208112648-445905ef2b6d
	github.com/evolbioinf/fasta v0.0.0-20220329100526-fa625fa59b5d
	github.com/evolbioinf/sus v0.0.0-20230320163303-b6d16dd4ec1f
	github.com/ivantsers/chr v0.0.0-20240904113825-8dfc26d67657
)

require github.com/ivantsers/fasta v0.0.0-20240830081231-39b1ecbb3ca0 // indirect

replace github.com/evolbioinf/fur/util => ../util
