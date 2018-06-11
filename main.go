package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	if err := mix.Init(0); err != nil {
		return fmt.Errorf("could not initialize MIX: %v", err)
	}
	defer mix.Quit()

	sndFmt := uint16(mix.DEFAULT_FORMAT)
	if mix.OpenAudio(44100, sndFmt, 2, 1024); err != nil {
		return fmt.Errorf("could not open audio mixer: %v", err)
	}

	music, err := mix.LoadMUS("res/music/Crimson_Nights_Track_02.mp3")
	if err != nil {
		return fmt.Errorf("could not load music: %v", err)
	}

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window %v", err)
	}
	defer w.Destroy()

	if !mix.PlayingMusic() {
		music.Play(10)
	}

	paintTitleScreen(r)
	event := sdl.WaitEvent()
	var mouseEvent *sdl.MouseButtonEvent
	for reflect.TypeOf(event) != reflect.TypeOf(mouseEvent) {
		event = sdl.WaitEvent()
	}

	s, err := newScene(r)
	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}
	defer s.destroy()

	events := make(chan sdl.Event)
	errc := s.run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}

func drawTitle(r *sdl.Renderer, text string) error {
	r.Clear()

	f, err := ttf.OpenFont("res/fonts/Rockwell Extra Bold.ttf", 20)

	if err != nil {
		return fmt.Errorf("could not return font: %v", err)
	}
	defer f.Close()

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := f.RenderUTF8Solid(text, c)
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}
	r.Present()

	return nil
}

func paintTitleScreen(r *sdl.Renderer) error {
	title, err := img.LoadTexture(r, "res/img/START-GAME.png")
	if err != nil {
		return fmt.Errorf("could not load title image: %v", err)
	}
	if err := r.Copy(title, nil, nil); err != nil {
		return fmt.Errorf("could not copy title: %v", err)
	}

	r.Present()
	return nil
}
