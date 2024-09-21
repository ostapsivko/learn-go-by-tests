package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
)

type FileSystemPlayerStore struct {
	league  League
	encoder *json.Encoder
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file %s, %v", file.Name(), err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		league:  league,
		encoder: json.NewEncoder(&tape{file}),
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	slices.SortFunc(f.league, func(a Player, b Player) int {
		return b.Score - a.Score
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	if player := f.league.Find(name); player != nil {
		return player.Score
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.encoder.Encode(f.league)
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating a file system player store, %v", err)
	}

	return store, closeFunc, nil
}
