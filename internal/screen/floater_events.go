package screen

import "github.com/micro-editor/tcell/v2"

func ProcessFloaterEvents(ev tcell.Event) bool {
    floaterMutex.Lock()
    defer floaterMutex.Unlock()
    
    if len(floatingWindows) == 0 {
        return false
    }

    switch ev := ev.(type) {
    case *tcell.EventKey:
        return handleFloaterKeyEvent(ev)
    case *tcell.EventMouse:
        return handleFloaterMouseEvent(ev)
    }
    return false
}

func handleFloaterKeyEvent(ev *tcell.EventKey) bool {
    // Key event handling logic
    if activeFloater != nil {
        if ev.Key() == tcell.KeyESC {
            CloseFloatingWindow(activeFloater.ID)
            return true
        }
        return true
    }
    return false
}

func handleFloaterMouseEvent(ev *tcell.EventMouse) bool {
    // Mouse event handling logic
    x, y := ev.Position()
    for i := len(floatingWindows) - 1; i >= 0; i-- {
        floater := floatingWindows[i]
        if floater.Visible && isPointInFloater(x, y, floater) {
            bringToFront(floater)
            if ev.Buttons()&tcell.Button1 != 0 {
                activeFloater = floater
            }
            return true
        }
    }
    return false
}