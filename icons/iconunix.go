//go:build linux || darwin

package icons

import "embed"

//go:embed iconunix.png
var file embed.FS

const filename = "iconunix.png"
