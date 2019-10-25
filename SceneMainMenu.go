package OpenDiablo2

import (
	"image/color"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/UI"

	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/hajimehoshi/ebiten"
)

// MainMenu represents the main menu
type MainMenu struct {
	engine              *Engine
	trademarkBackground *Common.Sprite
	background          *Common.Sprite
	diabloLogoLeft      *Common.Sprite
	diabloLogoRight     *Common.Sprite
	diabloLogoLeftBack  *Common.Sprite
	diabloLogoRightBack *Common.Sprite
	copyrightLabel      *UI.Label
	copyrightLabel2     *UI.Label
	showTrademarkScreen bool
	leftButtonHeld      bool
}

// CreateMainMenu creates an instance of MainMenu
func CreateMainMenu(engine *Engine) *MainMenu {
	result := &MainMenu{
		engine:              engine,
		showTrademarkScreen: true,
	}

	return result
}

// Load is called to load the resources for the main menu
func (v *MainMenu) Load() {
	v.engine.PlayBGM(ResourcePaths.BGMTitle)
	go func() {
		loadStep := 1.0 / 8.0
		v.engine.LoadingProgress = 0
		{
			v.copyrightLabel = UI.CreateLabel(v.engine, ResourcePaths.FontFormal12, Palettes.Static)
			v.copyrightLabel.Alignment = UI.LabelAlignCenter
			v.copyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.copyrightLabel.ColorMod = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel.MoveTo(400, 500)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.copyrightLabel2 = UI.CreateLabel(v.engine, ResourcePaths.FontFormal12, Palettes.Static)
			v.copyrightLabel2.Alignment = UI.LabelAlignCenter
			v.copyrightLabel2.SetText("All Rights Reserved.")
			v.copyrightLabel2.ColorMod = color.RGBA{188, 168, 140, 255}
			v.copyrightLabel2.MoveTo(400, 525)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.background = v.engine.LoadSprite(ResourcePaths.GameSelectScreen, Palettes.Sky)
			v.background.MoveTo(0, 0)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.trademarkBackground = v.engine.LoadSprite(ResourcePaths.TrademarkScreen, Palettes.Sky)
			v.trademarkBackground.MoveTo(0, 0)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.diabloLogoLeft = v.engine.LoadSprite(ResourcePaths.Diablo2LogoFireLeft, Palettes.Units)
			v.diabloLogoLeft.Blend = true
			v.diabloLogoLeft.Animate = true
			v.diabloLogoLeft.MoveTo(400, 120)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.diabloLogoRight = v.engine.LoadSprite(ResourcePaths.Diablo2LogoFireRight, Palettes.Units)
			v.diabloLogoRight.Blend = true
			v.diabloLogoRight.Animate = true
			v.diabloLogoRight.MoveTo(400, 120)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.diabloLogoLeftBack = v.engine.LoadSprite(ResourcePaths.Diablo2LogoBlackLeft, Palettes.Units)
			v.diabloLogoLeftBack.MoveTo(400, 120)
			v.engine.LoadingProgress += loadStep
		}
		{
			v.diabloLogoRightBack = v.engine.LoadSprite(ResourcePaths.Diablo2LogoBlackRight, Palettes.Units)
			v.diabloLogoRightBack.MoveTo(400, 120)
			v.engine.LoadingProgress = 1.0
		}
	}()
}

// Unload unloads the data for the main menu
func (v *MainMenu) Unload() {

}

// Render renders the main menu
func (v *MainMenu) Render(screen *ebiten.Image) {
	if v.showTrademarkScreen {
		v.trademarkBackground.DrawSegments(screen, 4, 3, 0)
	} else {
		v.background.DrawSegments(screen, 4, 3, 0)
	}
	v.diabloLogoLeftBack.Draw(screen)
	v.diabloLogoRightBack.Draw(screen)
	v.diabloLogoLeft.Draw(screen)
	v.diabloLogoRight.Draw(screen)

	if v.showTrademarkScreen {
		v.copyrightLabel.Draw(screen)
		v.copyrightLabel2.Draw(screen)
	} else {

	}
}

// Update runs the update logic on the main menu
func (v *MainMenu) Update() {
	if v.showTrademarkScreen {
		if v.engine.CursorButtonPressed(CursorButtonLeft) {
			v.leftButtonHeld = true
			v.showTrademarkScreen = false
		}
		return
	}
}
