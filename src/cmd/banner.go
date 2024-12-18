// Author: fluffelpuff
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"runtime"

	cenvxcore "github.com/custodia-cenv/cenvx-core/src"
	"github.com/custodia-cenv/cenvx-core/src/host"
	"github.com/custodia-cenv/cenvx-core/src/utils"
)

// Welcome Banner
func ShowBanner(mode cenvxcore.BannerMode) {
	osName := runtime.GOOS
	arch := runtime.GOARCH
	isAdmin := host.CheckAdmin()

	banner := fmt.Sprintf(`
 ██████╗██╗   ██╗███████╗████████╗ ██████╗ ██████╗ ██╗ █████╗
██╔════╝██║   ██║██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗██║██╔══██╗
██║     ██║   ██║███████╗   ██║   ██║   ██║██║  ██║██║███████║
██║     ██║   ██║╚════██║   ██║   ██║   ██║██║  ██║██║██╔══██║
╚██████╗╚██████╔╝███████║   ██║   ╚██████╔╝██████╔╝██║██║  ██║
 ╚═════╝ ╚═════╝ ╚══════╝   ╚═╝    ╚═════╝ ╚═════╝ ╚═╝╚═╝  ╚═╝
                                                                             
	WASM VM                  
	Version: %s
	Author: %s
	OS: %s
	Architecture: %s
	User is Admin: %t
----------------------------------------------------------------------------------`, utils.FormatNumberWithDots(int(cenvxcore.CoreVersion)), "fluffelpuff", osName, arch, isAdmin)
	fmt.Println(banner)

}
