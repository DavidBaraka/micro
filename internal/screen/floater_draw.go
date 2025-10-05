package screen

import "github.com/micro-editor/tcell/v2"

func DrawAllFloaters() {
    floaterMutex.Lock()
    defer floaterMutex.Unlock()
    
    if len(floatingWindows) == 0 {
        return
    }

    // Sort by ZIndex
    sorted := make([]*FloatingWindow, len(floatingWindows))
    copy(sorted, floatingWindows)
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].ZIndex < sorted[j].ZIndex
    })

    for _, floater := range sorted {
        if floater.Visible {
            drawFloatingWindow(floater)
        }
    }
}

func drawFloatingWindow(f *FloatingWindow) {
    drawFloaterBorder(f)
    if f.Title != "" {
        drawFloaterTitle(f)
    }
    drawFloaterContent(f)
}

// ... (drawFloaterBorder, drawFloaterTitle, drawFloaterContent implementations)