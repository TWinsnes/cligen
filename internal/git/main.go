package git

import (
	"net/url"
	"os/exec"
	"strings"
)

func ModuleFromGit() string {
	remote := RemoteURL()
	if remote == "" {
		return ""
	}
	return NormalizeRemoteToModule(remote)
}

func RemoteURL() string {
	// Prefer the common origin first
	if out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output(); err == nil {
		s := strings.TrimSpace(string(out))
		if s != "" {
			return s
		}
	}

	// Discover remotes and pick the first one if present
	if out, err := exec.Command("git", "remote").Output(); err == nil {
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, ln := range lines {
			name := strings.TrimSpace(ln)
			if name == "" {
				continue
			}
			// Try to resolve URL for the first remote name found
			if o, err := exec.Command("git", "config", "--get", "remote."+name+".url").Output(); err == nil {
				s := strings.TrimSpace(string(o))
				if s != "" {
					return s
				}
			}
			if o, err := exec.Command("git", "remote", "get-url", name).Output(); err == nil {
				s := strings.TrimSpace(string(o))
				if s != "" {
					return s
				}
			}
			// As a fallback, parse `git remote -v` and take the first URL line for this remote
			if o, err := exec.Command("git", "remote", "-v").Output(); err == nil {
				for _, l := range strings.Split(string(o), "\n") {
					fields := strings.Fields(strings.TrimSpace(l))
					if len(fields) >= 2 && fields[0] == name {
						url := strings.TrimSpace(fields[1])
						if url != "" {
							return url
						}
					}
				}
			}
			// Only consider the first remote by requirement
			break
		}
	}

	return ""
}

// normalizeRemoteToModule converts various Git remote URL formats to a module path like "github.com/user/repo".
func NormalizeRemoteToModule(remote string) string {
	s := strings.TrimSpace(remote)
	if s == "" {
		return ""
	}
	// Trim trailing .git if present
	if strings.HasSuffix(s, ".git") {
		s = strings.TrimSuffix(s, ".git")
	}

	var host, path string

	// Handle SCP-like SSH form: git@github.com:user/repo
	if (strings.HasPrefix(s, "git@") || strings.Contains(s, "@")) && strings.Contains(s, ":") && !strings.Contains(s, "://") {
		at := strings.Index(s, "@")
		colon := strings.Index(s, ":")
		if at >= 0 && colon > at {
			host = s[at+1 : colon]
			path = s[colon+1:]
		}
	} else if strings.Contains(s, "://") {
		// Handle URL forms: https://github.com/user/repo or ssh://git@github.com/user/repo
		if u, err := url.Parse(s); err == nil {
			host = u.Hostname()
			path = strings.TrimPrefix(u.Path, "/")
		}
	} else {
		// Possibly already in host/user/repo form
		parts := strings.Split(s, "/")
		if len(parts) >= 3 {
			return strings.Join(parts[:3], "/")
		}
	}

	path = strings.TrimPrefix(path, "/")
	if host == "" || path == "" {
		return ""
	}
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return ""
	}
	return host + "/" + parts[0] + "/" + parts[1]
}
