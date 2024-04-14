package music

import "mh-api/app/internal/domain/monsters"

type Music struct {
	musicId   MusicId
	monsterId monsters.MonsterId
	name      MusicName
	imageUrl  MusicImageUrl
}

func newMusic(MusicId MusicId, monsterId monsters.MonsterId, name MusicName, imageUrl MusicImageUrl) *Music {
	return &Music{MusicId, monsterId, name, imageUrl}
}

func NewMusic(musicId string, monsterId string, name string, imageUrl string) *Music {
	return newMusic(
		MusicId{value: musicId},
		monsters.MonsterId{Value: monsterId},
		MusicName{value: name},
		MusicImageUrl{value: imageUrl},
	)
}
