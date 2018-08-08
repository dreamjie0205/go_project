package mlib

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager faild.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager faild. not empty")
	}

	m0 := &MusicEntry{
		"1", "My heart will go on", "Celion Dion", "http://qbox.me/23501234", "mp3",
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() faild.")
	}

	m := mm.Find(m0.Name)

	if m == nil {
		t.Error("MusicManager.Find() faild.")
	}

	if m.Id != m0.Id || m.Name != m0.Name || m.Artist != m0.Artist || m.Source != m0.Source ||
		m.Type != m0.Type {
		t.Error("MusicManager.Find() faild. Found item Mismatch")
	}
	m, err := mm.Get(0)

	if m == nil {
		t.Error("MusicManager.Get() failed", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() faild.", err)
	}
}
