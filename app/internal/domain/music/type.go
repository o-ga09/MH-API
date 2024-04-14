package music

type Musics []Music

type MusicId struct{ value string }
type MusicName struct{ value string }
type MusicImageUrl struct{ value string }

func (f *Music) GetMonsterID() string { return f.monsterId.Value }
func (f *Music) GetID() string        { return f.musicId.value }
func (f *Music) GetName() string      { return f.name.value }
func (f *Music) GetURL() string       { return f.imageUrl.value }
