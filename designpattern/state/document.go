package main

type document struct {
	nowState                   state
	draft, moderation, publish state
}

func newDocument() *document {
	d := &document{}
	draft := draftState{d}
	moder := moderationState{d}
	publish := publishState{d}
	d.nowState = draft
	d.draft = draft
	d.moderation = moder
	d.publish = publish
	return d
}

func (d *document) setState(s state) {
	d.nowState = s
}

func (d *document) execute() {
	d.nowState.execute()
}
