package mlib

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

/**
 * 创建一个空的管理器 构造函数 用来实例化管理者的类
 */
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//获取管理器中的音乐长度
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//获取某个位置的音乐实体
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("index out of range.")
	}
	return &m.musics[index], nil
}

//根据名称 查找某个音乐实体
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

// 添加一首音乐
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//移动一首音乐
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}
