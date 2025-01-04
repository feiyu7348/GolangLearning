// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 目标接口
type AudioPlayer interface {
	Play(filename string)
}

type MP3Player struct {
}

func (mp3 *MP3Player) Play(filename string) {
	fmt.Println(filename)
}

// 被适配者：WAV播放器
type WAVPlayer struct {
}

func (w *WAVPlayer) PlayWAV(filename string) {
	fmt.Println("播放WAV音频：", filename)
}

// 适配器
type WAVAdapter struct {
	wavPlayer *WAVPlayer
}

func (wa *WAVAdapter) Play(filename string) {
	wa.wavPlayer.PlayWAV(filename)
}

// 客户端代码
func main() {
	mp3Player := &MP3Player{}

	// 播放MP3音频
	mp3Player.Play("music.mp3")

	wavPlayer := &WAVPlayer{}

	// 使用适配器播放WAV音频
	wavAdapter := &WAVAdapter{wavPlayer}
	wavAdapter.Play("music.wav")
}
