package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

type ConcurrentFreqMap struct {
	fm FreqMap
	mu *sync.Mutex
}

func (cfm *ConcurrentFreqMap) Inc(r rune) {
	defer cfm.mu.Unlock()
	cfm.mu.Lock()
	cfm.fm[r]++
}

func NewConcurrentFreqMap() ConcurrentFreqMap {
	return ConcurrentFreqMap{
		fm: make(FreqMap),
		mu: &sync.Mutex{},
	}
}

func ConcurrentFrequency(inputs []string) FreqMap {
	m := NewConcurrentFreqMap()
	var wg sync.WaitGroup
	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(s string) {
			for _, r := range s {
				m.Inc(r)
			}
			wg.Done()
		}(input)
	}
	wg.Wait()
	return m.fm
}
