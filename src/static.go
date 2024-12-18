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

package cenvxcore

const (
	// Gibt den Status des Core Osbjektes an
	NEW      CoreState = 1
	INITED   CoreState = 2
	SERVING  CoreState = 3
	SHUTDOWN CoreState = 4
	CLOSED   CoreState = 5

	// Legt die Aktuelle Version fest
	CoreVersion Vesion = 1000000000

	// Die Repo wird festgelegt
	CoreRepo CoreRepoUrl = "https://github.com/custodia-cenv/cenvx-core"

	// Gibt den Prefix der Core Socket Dateien an
	CoreIpcVmSocketPrefix CoreIpcVmSocketIdentifierPrefix = "cuspp"

	// Gibt die Verschiedenen Banner Modi an
	CoreBanner BannerMode = BannerMode("Core")
	VMBanner   BannerMode = BannerMode("VM")
)
