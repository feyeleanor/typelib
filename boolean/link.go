package boolean

type link struct {
	value	bool
	next	*link
}

func (l link) String() string {
	if l.value {
		return "true"
	}
	return "false"
}

func (l *link) Append(v bool) (r *link) {
	r = &link{ value: v }
	if l != nil {
		l.next = r
	}
	return
}

func (l link) Next() *link {
	return l.next
}