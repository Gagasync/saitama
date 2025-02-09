package main

type SourceText struct {
	text     []rune
	location int
	length   int
}

func NewSourceText(text string) *SourceText {
	return &SourceText{text: []rune(text), length: len(text), location: 0}
}

func (st *SourceText) Current() rune {
	return st.Peek(0)
}

func (st *SourceText) LA() rune {
	return st.Peek(1)
}

func (st *SourceText) Peek(n int) rune {
	if st.location+n >= st.length {
		return '\000'
	}
	return st.text[st.location+n]
}

func (st *SourceText) Consume() {
	st.location++
}

func (st *SourceText) IsEnd() bool {
	return st.location >= st.length
}
