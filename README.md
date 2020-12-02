# Advent of Code 2020 - Rust/WASM

[Advent of code](https://adventofcode.com/)

## Setup

Follow the `rust` [docs](https://www.rust-lang.org/) to install rust and `wasm-pack`. [MDN guide](https://developer.mozilla.org/en-US/docs/WebAssembly/Rust_to_wasm)

## Build

Compile to WASM

```bash
$ wasm-pack build --release --target web
```

Copy `advent_of_code_2020.js` & `advent_of_code_2020_bg.wasm` from the `/pkg` directory to the `/public` directory.

Run the web server
```bash
$ go run main.go
```

## TODO 
Configure `Go` to compile to wasm and compare. (bundle size etc)

## Misc
[Shrinking Rust for `wasm`](https://rustwasm.github.io/book/game-of-life/code-size.html#shrinking-wasm-size)