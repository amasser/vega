package updater

import (
	"fmt"

	semver "github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

type Updater struct {
	Repo    string
	Version string
}

func NewUpdater(repo string, version string) *Updater {
	u := &Updater{
		Repo:    repo,
		Version: version,
	}
	return u
}

func (u *Updater) IsLatestAvailable() (string, bool, error) {
	latest, found, err := selfupdate.DetectLatest(u.Repo)
	if err != nil {
		return "", false, fmt.Errorf("error occurred while detecting version: %v", err)
	}

	v, err := semver.ParseTolerant(u.Version)
	if err != nil {
		return "", false, fmt.Errorf("not able to parse version: %v - %v", err ,u.Version)
	}
	if !found || latest.Version.LTE(v) {
		return "", false, nil
	} else {
		return latest.Version.String(), true, nil
	}

	return "", false, nil
}

func (u *Updater) SelfUpdate() error {
	currentVersion, _ := semver.ParseTolerant(u.Version)

	latest, err := selfupdate.UpdateSelf(currentVersion, u.Repo)
	if err != nil {
		return err
	}

	if currentVersion.Equals(latest.Version) {
		// Do Nothing
	} else {
		fmt.Println("Successfully updated to version", latest.Version)
		fmt.Println("Release Note:\n", latest.ReleaseNotes)
	}

	return nil
}
