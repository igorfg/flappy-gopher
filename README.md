# flappy-tiago
Trying to learn how to make a game using golang

This game was made based on flappy-gopher from "JustForFunc: Programming in Go" youtube channel
https://www.youtube.com/watch?v=aYkxFbd6luY

# Requirements
* [Install golang](https://golang.org/doc/install)
* [SDL2](http://libsdl.org/download-2.0.php)
* [SDL2_image](http://www.libsdl.org/projects/SDL_image/)
* [SDL2_mixer](http://www.libsdl.org/projects/SDL_mixer/)
* [SDL2_ttf](http://www.libsdl.org/projects/SDL_ttf/)
* [SDL2_gfx (optional)](http://www.ferzkopp.net/wordpress/2016/01/02/sdl_gfx-sdl2_gfx/)
* [Golang SDL Wrapper](https://github.com/veandco/go-sdl2)

# Building
If you whish to build flappy-tiago you need to:

1. Install golang\
https://golang.org/doc/install

2. Install SDL on your machine\
If you use a distribuition that update their repositories frequently, the easiest way is to install SDL is:\
`$ apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev`

3. Get SDL wrapper for golang\
`$ go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}`

4. Get flappy-tiago's source code\
`$ go get -v github.com/igorfg/flappy-tiago`

5.Build the executable file\
Go to your /home/yousername/go/src/github.com/igorfg/flappy-tiago
`$ go build`
`$ ./flappy-tiago`
