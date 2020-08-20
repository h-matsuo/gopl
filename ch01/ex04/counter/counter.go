package counter

type LineCounter struct {
	counterByLine map[string]*Counter // key: line
}

type Counter struct {
	line        string
	occurrences uint
	fileNames   map[string]struct{} // Use map[K]struct{} as Set<K>
}

func NewLineCounter() *LineCounter {
	return &LineCounter{
		counterByLine: make(map[string]*Counter),
	}
}

func (l *LineCounter) ToSlice() []*Counter {
	result := make([]*Counter, 0, len(l.counterByLine))
	for _, counter := range l.counterByLine {
		result = append(result, counter)
	}
	return result
}

func (l *LineCounter) AddLine(line string, fileName string) {
	c, ok := l.counterByLine[line]
	if !ok {
		c = &Counter{
			line:        line,
			occurrences: 0,
			fileNames:   make(map[string]struct{}),
		}
		l.counterByLine[line] = c
	}
	c.increment(fileName)
}

func (c *Counter) increment(fileName string) {
	c.occurrences++
	c.fileNames[fileName] = struct{}{}
}

func (c *Counter) Line() string {
	return c.line
}

func (c *Counter) Occurrences() uint {
	return c.occurrences
}

func (c *Counter) FileNames() []string {
	result := make([]string, 0, len(c.fileNames))
	for fileName := range c.fileNames {
		result = append(result, fileName)
	}
	return result
}
