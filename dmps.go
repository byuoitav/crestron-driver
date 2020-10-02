package dmps

import "context"

type DMPS struct{}

func (vs *DMPS) GetAudioVideoInputs(ctx context.Context, output string) (map[string]string, error) {
	return nil, nil
}

func (vs *DMPS) SetAudioVideoInput(ctx context.Context, output, input string) error {
	return nil
}

func (vs *DMPS) GetVolumes(ctx context.Context, blocks []string) (map[string]int, error) {
	return nil, nil
}

func (vs *DMPS) GetMutes(ctx context.Context, blocks []string) (map[string]bool, error) {
	return nil, nil
}

func (vs *DMPS) SetVolume(ctx context.Context, block string, volume int) error {
	return nil
}

func (vs *DMPS) SetMuted(ctx context.Context, block string, muted bool) error {
	return nil
}

func (vs *DMPS) GetInfo(ctx context.Context) (interface{}, error) {
	return nil, nil
}
