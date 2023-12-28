package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/AllenDang/giu"
)

var (
	fontSize    float32 = 64.0
	currentTime time.Time
	colorBg     color.Color = color.RGBA{50, 50, 50, 255}
	colorText   color.Color = color.RGBA{241, 93, 14, 255}
)

func refresh() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		currentTime = time.Now()
		giu.Update()

		<-ticker.C
	}
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Style().SetFontSize(fontSize).To(
			giu.Style().SetColor(giu.StyleColorText, colorText).To(
				giu.Style().SetColor(giu.StyleColorWindowBg, colorText).To(
					giu.Align(giu.AlignCenter).To(
						giu.Label(fmt.Sprintf("%s", currentTime.Format(time.Kitchen))),
						giu.Label(fmt.Sprintf("%s", currentTime.Format(time.DateOnly))),
					)))),
	)
}

func onAnyKeyPressed(key giu.Key, mod giu.Modifier, action giu.Action) {
	if action == giu.Press {
		os.Exit(1)
	}
}

func main() {
	wnd := giu.NewMasterWindow("Clock", 420, 160, giu.MasterWindowFlagsNotResizable)
	wnd.SetAdditionalInputHandlerCallback(onAnyKeyPressed)
	go refresh()
	giu.Context.FontAtlas.SetDefaultFont("RubikDoodleShadow.ttf", fontSize)
	giu.Context.FontAtlas.AddFont("RubikDoodleShadow.ttf", fontSize)
	giu.PushColorWindowBg(colorBg)

	wnd.Run(loop)
}
