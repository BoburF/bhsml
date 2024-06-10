package keys

type BHSML struct {
	Head
	Body
	Foot
}

func (bhl *BHSML) SetHead(head Head) {
	bhl.Head = head
}

func (bhl *BHSML) SetBody(body Body) {
	bhl.Body = body
}

func (bhl *BHSML) SetFoot(foot Foot) {
	bhl.Foot = foot
}
