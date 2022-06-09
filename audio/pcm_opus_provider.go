package audio

import (
	"io"

	"github.com/disgoorg/disgo/audio/opus"
	"github.com/disgoorg/disgo/voice"
)

// NewPCMOpusProvider creates a new voice.OpusFrameProvider which gets PCM frames from the given PCMFrameProvider and encodes the PCM frames into Opus frames.
// You can pass your own *opus.Encoder or nil to use the default Opus encoder(48000hz sample rate, 2 channels, opus.ApplicationAudio & 64kbps bitrate).
func NewPCMOpusProvider(encoder *opus.Encoder, pcmProvider PCMFrameProvider) voice.OpusFrameProvider {
	if encoder == nil {
		var err error
		if encoder, err = opus.NewEncoder(48000, 2, opus.ApplicationAudio); err != nil {
			panic("NewPCMOpusProvider: " + err.Error())
		}
		if err = encoder.Ctl(opus.SetBitrate(64000)); err != nil {
			panic("SetBitrate: " + err.Error())
		}

	}
	return &pcmOpusProvider{
		encoder:     encoder,
		pcmProvider: pcmProvider,
		opusBuff:    make([]byte, 2000),
	}
}

type pcmOpusProvider struct {
	encoder     *opus.Encoder
	pcmProvider PCMFrameProvider
	opusBuff    []byte
}

func (p *pcmOpusProvider) ProvideOpusFrame() []byte {
	pcm := p.pcmProvider.ProvidePCMFrame()
	if len(pcm) == 0 {
		return nil
	}

	n, err := p.encoder.Encode(pcm, p.opusBuff)
	if err != nil {
		if err != io.EOF {
			panic("ProvideOpusFrame: " + err.Error())
		}
		return nil
	}
	return p.opusBuff[:n]
}

func (p *pcmOpusProvider) Close() {
	p.encoder.Destroy()
	p.pcmProvider.Close()
}
