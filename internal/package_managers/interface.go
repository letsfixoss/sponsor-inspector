package packagemanagers

import "github.com/letsfixoss/gh-sponsor-grabber/internal/repositories"

type PackageManager interface {
	AppliesToRepo(r repositories.Repository) bool
}
