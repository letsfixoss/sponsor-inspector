package packagemanagers

import "github.com/letsfixoss/gh-sponsor-grabber/repositories"

type PackageManager interface {
	AppliesToRepo(r repositories.Repository) bool
}
