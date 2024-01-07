package main

type LabelLens struct {
	label string
	focal int
}

func findLabelLensIdx(lenses []LabelLens, label string) int {
	for idx, lens := range lenses {
		if lens.label == label {
			return idx
		}
	}
	return -1
}

type HashMap map[int][]LabelLens

func (hm *HashMap) Insert(boxNum int, labelLen LabelLens) {
	labelLenses := (*hm)[boxNum]
	if labelLenses == nil {
		(*hm)[boxNum] = []LabelLens{labelLen}
		return
	}
	idx := findLabelLensIdx(labelLenses, labelLen.label)
	if idx != -1 {
		labelLenses[idx] = labelLen
	} else {
		labelLenses = append(labelLenses, labelLen)
	}
	(*hm)[boxNum] = labelLenses
}

func (hm *HashMap) Remove(boxNum int, label string) {
	lenses := (*hm)[boxNum]
	if lenses == nil {
		return
	}
	idx := findLabelLensIdx(lenses, label)
	if idx == -1 {
		return
	}
	lenses = append(lenses[:idx], lenses[idx+1:]...)
	(*hm)[boxNum] = lenses
}

func NewHashMap() *HashMap {
	hashMap := make(HashMap)
	return &hashMap
}
