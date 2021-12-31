package header

type GenericHeader struct {
	HeaderName string
	Content    string
}

func (g *GenericHeader) String() string {
	return g.HeaderName + ":" + g.Content
}

func (g *GenericHeader) Name() string {
	return g.HeaderName
}

func (g *GenericHeader) Value() string {
	return g.Content
}

func (g *GenericHeader) Equals(other interface{}) bool {

	if h, ok := other.(*GenericHeader); ok {
		if g == h {
			return true
		}

		if h == nil && g != nil || h != nil && g == nil {
			return false
		}

		return h.Content == g.Content && h.HeaderName == g.HeaderName
	}
	return false
}

func (g *GenericHeader) Clone() Header {
	var newGenericHeader *GenericHeader

	if g == nil {
		return newGenericHeader
	}

	newGenericHeader = &GenericHeader{
		HeaderName: g.HeaderName,
		Content:    g.Content,
	}
	return newGenericHeader
}
