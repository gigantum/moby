package daemon

import (
	"runtime"
	"testing"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/container"
	"github.com/docker/docker/volume"
)

func TestParseVolumesFrom(t *testing.T) {
	cases := []struct {
		spec    string
		expID   string
		expMode string
		fail    bool
	}{
		{"", "", "", true},
		{"foobar", "foobar", "rw", false},
		{"foobar:rw", "foobar", "rw", false},
		{"foobar:ro", "foobar", "ro", false},
		{"foobar:baz", "", "", true},
	}

	parser := volume.NewParser(runtime.GOOS)

	for _, c := range cases {
		id, mode, err := parser.ParseVolumesFrom(c.spec)
		if c.fail {
			if err == nil {
				t.Fatalf("Expected error, was nil, for spec %s\n", c.spec)
			}
			continue
		}

		if id != c.expID {
			t.Fatalf("Expected id %s, was %s, for spec %s\n", c.expID, id, c.spec)
		}
		if mode != c.expMode {
			t.Fatalf("Expected mode %s, was %s for spec %s\n", c.expMode, mode, c.spec)
		}
	}
}

func TestMountSubdir(t *testing.T) {
	d := Daemon{containers: container.NewMemoryStore()}
	mountSubdir := mount.Mount{
		mount.TypeVolume,
		"check",
		"check2",
		false,
		mount.ConsistencyDelegated,
		nil,
		&mount.VolumeOptions{
			Subpath: "test",
		},
		nil,
	}
	container := container.Container{}
	hostConfig := containertypes.HostConfig{
		Mounts: []mount.Mount{
			mountSubdir,
		},
	}

	err := d.registerMountPoints(&container, &hostConfig)

	if err != nil {
		t.Fatalf("unexpected error")
	}
}
