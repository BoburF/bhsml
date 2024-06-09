package keys

type BHSML struct {
	Head
	Body
	Foot
}

func (bhl *BHSML) SetHead(head Head) {
	bhl.Head = head
}
