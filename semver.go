package semver

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrorWrongVersion = errors.New("wrong version string")
)

type Semver struct {
	Major int
	Minor int
}

func (s *Semver) MajorGreatThen(sem Semver) bool {
	return s.Major > sem.Major
}

func (s *Semver) MinorGreatThen(sem Semver) bool {
	return s.Minor > sem.Minor
}

func Parse(version string) (Semver, error) {
	versions := strings.Split(version, ".")
	if len(versions) < 2 {
		return Semver{}, ErrorWrongVersion
	}

	major, err := strconv.Atoi(versions[0])
	if err != nil {
		return Semver{}, errors.Wrap(err, "parse major version")
	}

	minor, err := strconv.Atoi(versions[1])
	if err != nil {
		return Semver{}, errors.Wrap(err, "parse minor version")
	}

	return Semver{
		Major: major,
		Minor: minor,
	}, nil
}
