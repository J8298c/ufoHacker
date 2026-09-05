# ufoHacker

A terminal "hacking" game built with [Bubble Tea](https://charm.land/bubbletea) and [Lipgloss](https://charm.land/lipgloss). Browse a list of real UFO sighting reports from the National UFO Reporting Center (NUFORC), then crack the encrypted summary of each one by guessing its Caesar cipher shift key.

## Gameplay

- **List view** — browse sightings with their city, state, coordinates, and reported shape.
- **Detail view** — pick a sighting to see its scrambled (encrypted) summary and try to decrypt it.
- Enter a 2-digit shift key and press `Enter` to attempt decryption.
- You get 3 attempts before the terminal auto-decrypts the transmission for you.

## Controls

| Key(s)          | Action                              |
|-----------------|--------------------------------------|
| `↑`/`k`, `↓`/`j` | Move cursor in the list view          |
| `Enter`         | Select a record / submit a guess     |
| `0`-`9`         | Enter digits for the shift key guess |
| `Backspace`     | Delete last digit entered            |
| `Esc`           | Return to the list view              |
| `q`             | Quit (list view) / back to list      |
| `Ctrl+C`        | Quit                                 |

## Running

Requires Go 1.27+.

```sh
go run .
```

## Project layout

- [main.go](main.go) — Bubble Tea model, update loop, and views
- [banner.go](banner.go) — ASCII-art terminal banner
- [styles.go](styles.go) — Lipgloss color/style definitions
- [records/records.go](records/records.go) — UFO sighting records (summaries stored ROT13-encoded)
