package main


type Pattern struct {
	matched bool
	pattern [] byte
	plen int

}

func New(pattern string) *Pattern {
	b := []byte(pattern)
	return &Pattern{false,b,len(b)}
}

func (p *Pattern) Match(text []byte, tlen int) bool {
	p.matched = false
	p.rmatch(0,0,tlen,text)
	return p.matched
}

func (p *Pattern) rmatch( ti , pj, tlen int , text []byte )  {
	if p.matched {
		return
	}
	if pj == p.plen {
		if ti == tlen {
			p.matched = true
			return
		}
	}
	if p.pattern[pj] == '*' { //匹配任意个字符
		for k := 0; k < tlen - ti; k++ {
			p.rmatch(ti+k,pj+1,tlen,text)
		}
	}else if p.pattern[pj] == '?' {
		p.rmatch(ti,pj+1,tlen,text)
		p.rmatch(ti+1,pj+1,tlen,text)
	} else if ti < tlen && p.pattern[pj] == text[ti] {
		p.rmatch(ti+1,pj+1,tlen,text)
	}
}



