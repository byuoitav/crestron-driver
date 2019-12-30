package dmps

import "context"

type DMPS struct{}

func (vs *DMPS) GetInputByOutput(ctx context.Context, output string) (string, error) {}

func (vs *DMPS) SetInputByOutput(ctx context.Context, output, input string) error {}

func (vs *DMPS) GetVolumeByBlock(ctx context.Context, block string) (int, error) {}

func (vs *DMPS) GetMutedByBlock(ctx context.Context, block string) (bool, error) {}

func (vs *DMPS) SetVolumeByBlock(ctx context.Context, block string, volume int) error {}

func (vs *DMPS) SetMutedByBlock(ctx context.Context, block string, muted bool) error {}

func (vs *DMPS) GetInfo(ctx context.Context) (interface{}, error) {
	return nil, nil
}
