package html

import "fmt"

// MediaElement represents the <media> HTML element.

type MediaElement[C HTMLContent] basicElement[C]


// Media creates a new <media> element.
func Media(children ...HTMLContent) *MediaElement[HTMLContent] {
	return &MediaElement[HTMLContent]{
		tagName: "media",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *MediaElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *MediaElement[C]) Attr(key, value string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// AudioTracks sets the audioTracks attribute.
func (e *MediaElement[C]) AudioTracks(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "audioTracks", Value: v})
	return e
}



// Autoplay sets the autoplay attribute.
func (e *MediaElement[C]) Autoplay(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "autoplay", Value: "autoplay"})
	}
	return e
}



// Buffered sets the buffered attribute.
func (e *MediaElement[C]) Buffered(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "buffered", Value: v})
	return e
}



// Controls sets the controls attribute.
func (e *MediaElement[C]) Controls(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "controls", Value: "controls"})
	}
	return e
}



// CurrentSrc sets the currentSrc attribute.
func (e *MediaElement[C]) CurrentSrc(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "currentSrc", Value: v})
	return e
}



// CurrentTime sets the currentTime attribute.
func (e *MediaElement[C]) CurrentTime(v float64) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "currentTime", Value: fmt.Sprintf("%g", v)})
	return e
}



// DefaultMuted sets the defaultMuted attribute.
func (e *MediaElement[C]) DefaultMuted(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "defaultMuted", Value: "defaultMuted"})
	}
	return e
}



// DefaultPlaybackRate sets the defaultPlaybackRate attribute.
func (e *MediaElement[C]) DefaultPlaybackRate(v float64) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "defaultPlaybackRate", Value: fmt.Sprintf("%g", v)})
	return e
}



// Ended sets the ended attribute.
func (e *MediaElement[C]) Ended(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "ended", Value: "ended"})
	}
	return e
}



// Loop sets the loop attribute.
func (e *MediaElement[C]) Loop(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "loop", Value: "loop"})
	}
	return e
}



// Muted sets the muted attribute.
func (e *MediaElement[C]) Muted(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "muted", Value: "muted"})
	}
	return e
}



// Paused sets the paused attribute.
func (e *MediaElement[C]) Paused(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "paused", Value: "paused"})
	}
	return e
}



// PlaybackRate sets the playbackRate attribute.
func (e *MediaElement[C]) PlaybackRate(v float64) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "playbackRate", Value: fmt.Sprintf("%g", v)})
	return e
}



// Played sets the played attribute.
func (e *MediaElement[C]) Played(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "played", Value: v})
	return e
}



// Preload sets the preload attribute.
func (e *MediaElement[C]) Preload(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "preload", Value: v})
	return e
}



// PreservesPitch sets the preservesPitch attribute.
func (e *MediaElement[C]) PreservesPitch(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "preservesPitch", Value: "preservesPitch"})
	}
	return e
}



// Seekable sets the seekable attribute.
func (e *MediaElement[C]) Seekable(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "seekable", Value: v})
	return e
}



// Seeking sets the seeking attribute.
func (e *MediaElement[C]) Seeking(v bool) *MediaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "seeking", Value: "seeking"})
	}
	return e
}



// Src sets the src attribute.
func (e *MediaElement[C]) Src(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// TextTracks sets the textTracks attribute.
func (e *MediaElement[C]) TextTracks(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "textTracks", Value: v})
	return e
}



// VideoTracks sets the videoTracks attribute.
func (e *MediaElement[C]) VideoTracks(v string) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "videoTracks", Value: v})
	return e
}



// Volume sets the volume attribute.
func (e *MediaElement[C]) Volume(v float64) *MediaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "volume", Value: fmt.Sprintf("%g", v)})
	return e
}


