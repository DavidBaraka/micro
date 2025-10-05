package screen

import (
    "sort"
    "sync"

    "github.com/micro-editor/tcell/v2"
    "github.com/zyedidia/micro/v2/internal/config"
)

var (
    floatingWindows []*FloatingWindow
    activeFloater   *FloatingWindow
    floaterZIndex   int = 1000
    floaterMutex    sync.Mutex
)

type FloatingWindow struct {
    ID       uint64
    X, Y     int
    Width, Height int
    Title    string
    Content  []string
    ZIndex   int
    Modal    bool
    Visible  bool
    Style    tcell.Style
    Callback func(string) // For interactive floaters
}

func CreateFloatingWindow(title string, content []string, x, y, w, h int, modal bool) uint64 {
    floaterMutex.Lock()
    defer floaterMutex.Unlock()
    
    floater := &FloatingWindow{
        ID:      uint64(len(floatingWindows) + 1),
        X:       x, Y: y, Width: w, Height: h,
        Title: title, Content: content,
        ZIndex: floaterZIndex, Modal: modal, Visible: true,
        Style: config.DefStyle,
    }
    
    floaterZIndex++
    floatingWindows = append(floatingWindows, floater)
    
    if modal || activeFloater == nil {
        activeFloater = floater
    }
    
    Redraw()
    return floater.ID
}

func CloseFloatingWindow(id uint64) {
    floaterMutex.Lock()
    defer floaterMutex.Unlock()
    
    for i, floater := range floatingWindows {
        if floater.ID == id {
            floatingWindows = append(floatingWindows[:i], floatingWindows[i+1:]...)
            updateActiveFloater()
            Redraw()
            return
        }
    }
}