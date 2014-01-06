package filesystem

import (
	"os"
	"path/filepath"
)

func SymlinkVideos(videos []*Video, dir string) error {
	for _, video := range videos {
		err := os.Symlink(video.Path, filepath.Join(dir, video.Name))
		if err != nil {
			return err
		}
	}
	return nil
}
