package intest

type InSet struct {
	wrods []uint64
}

func (s *InSet)Has(x int) bool  {
	word, bit := x/64, uint(x%64)
	return word<len(s.wrods) && s.wrods[word]&(1<<bit) != 0
}
func (s *InSet) UnionWith(t *InSet)  {
	for i, tword := range t.wrods{
		if i<len(s.wrods){
			s.wrods[i] |= tword
		}else {
			s.wrods = append(s.wrods, tword)
		}
	}
}
func (s *InSet)add(x int)  {
	word, bit := x/64, uint(x%64)
	for word>=len(s.wrods){
		s.wrods = append(s.wrods,0)
	}
	s.wrods[word] |= 1<<bit
}
