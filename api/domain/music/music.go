package music

import "mh-api/api/domain/monsters"

type Music struct {
	MusicId   MusicId
	monsterId monsters.MonsterId
	name      MusicName
	imageUrl  MusicImageUrl
}

func newMusic(MusicId MusicId, monsterId monsters.MonsterId, name MusicName, imageUrl MusicImageUrl) *Music {
	return &Music{MusicId, monsterId, name, imageUrl}
}

func NewFiled(musicId string, monsterId string, name string, imageUrl string) *Music {
	return newMusic(
		MusicId{value: musicId},
		monsters.MonsterId{Value: monsterId},
		MusicName{value: name},
		MusicImageUrl{value: imageUrl},
	)
}
