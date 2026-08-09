package deej

import _ "embed"

// Logo is a binary representation of the deej logo; used for notifications and tray icon
//
//go:embed assets/logo-512.png
var Logo []byte

// EditConfig is the cog icon in the edit config menu option
//
//go:embed assets/menu-items/edit-config.ico
var EditConfig []byte

// RefreshSessions is the reload icon in the refresh sessions menu option
//
//go:embed assets/menu-items/refresh-sessions.ico
var RefreshSessions []byte
