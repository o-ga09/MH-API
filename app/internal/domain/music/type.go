package music

type MusicId struct{ value string }
type MusicName struct{ value string }
type MusicImageUrl struct{ value string }

func (f *MusicId) GetID() string        { return f.value }
func (f *MusicName) GetName() string    { return f.value }
func (f *MusicImageUrl) GetURL() string { return f.value }
