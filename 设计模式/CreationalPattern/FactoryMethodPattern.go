// Package CreationalPattern
// author: zfy  date: 2024/12/6
package CreationalPattern

import "fmt"

// 工厂方法模式 https://developer.aliyun.com/article/1271216

type MusicPlayer interface {
	Play()
}

type MP3Player struct {
}

func (mp3 *MP3Player) Play() {
	fmt.Println("播放MP3音乐")
}

type WAVPlayer struct {
}

func (wav *WAVPlayer) Play() {
	fmt.Println("播放WAV音乐")
}

type MusicPlayerFactory interface {
	CreatePlayer() MusicPlayer
}

type MP3PlayerFactory struct {
}

func (mp3Factory *MP3PlayerFactory) CreatePlayer() MusicPlayer {
	return &MP3Player{}
}

type WAVPlayerFactory struct {
}

func (wavFactory *WAVPlayerFactory) CreatePlayer() MusicPlayer {
	return &WAVPlayer{}
}

func main2() {
	mp3Factory := &MP3PlayerFactory{}
	mp3Player := mp3Factory.CreatePlayer()
	mp3Player.Play()

	wavFactory := &WAVPlayerFactory{}
	wavPlayer := wavFactory.CreatePlayer()
	wavPlayer.Play()
}
