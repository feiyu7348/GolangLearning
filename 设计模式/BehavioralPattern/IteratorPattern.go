// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 迭代器模式 https://developer.aliyun.com/article/1268022

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// MusicPlayerIterator 具体迭代器：音乐播放器迭代器
type MusicPlayerIterator struct {
	musicPlayer *MusicPlayer
	index       int
}

func NewMusicPlayerIterator(musicPlayer *MusicPlayer) *MusicPlayerIterator {
	return &MusicPlayerIterator{
		musicPlayer: musicPlayer,
		index:       0,
	}
}

func (it *MusicPlayerIterator) HasNext() bool {
	return it.index < len(it.musicPlayer.songs)
}

func (it *MusicPlayerIterator) Next() interface{} {
	if it.HasNext() {
		song := it.musicPlayer.songs[it.index]
		it.index++
		return song
	}
	return nil
}

// MusicPlayer 容器：音乐播放器
type MusicPlayer struct {
	songs []string
}

func (mp *MusicPlayer) AddSong(song string) {
	mp.songs = append(mp.songs, song)
}

func (mp *MusicPlayer) GetIterator() Iterator {
	return NewMusicPlayerIterator(mp)
}

// 客户端代码
func main() {
	player := &MusicPlayer{}
	player.AddSong("Song 1")
	player.AddSong("Song 2")
	player.AddSong("Song 3")

	iterator := player.GetIterator()
	for iterator.HasNext() {
		song := iterator.Next().(string)
		fmt.Println("Playing:", song)
	}
}
