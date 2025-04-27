package music

import "mh-api/app/internal/domain/monsters"

type Music struct {
	musicId   MusicId
	monsterId monsters.MonsterId
	name      MusicName
	Url       MusicUrl
}

func newMusic(musicId MusicId, monsterId monsters.MonsterId, name MusicName, url MusicUrl) *Music {
	return &Music{musicId, monsterId, name, url}
}

func NewMusic(musicId string, monsterId string, name string, url string) *Music {
	return newMusic(
		MusicId{value: musicId},
		monsters.MonsterId{Value: monsterId},
		MusicName{value: name},
		MusicUrl{value: url},
	)
}
