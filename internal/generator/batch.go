package generator

func GenerateBatch(n int) [][]byte {
	var batch [][]byte

	for i := 0; i < n; i++ {
		strat := GenerateRandomStrategy()
		js, _ := ToJSON(strat)
		batch = append(batch, js)
	}

	return batch
}
