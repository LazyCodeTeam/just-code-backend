package port

import (
	"context"
	"io"
)

type FileRepository interface {
	UploadProfileAvatar(ctx context.Context, reader io.Reader, profileId string) (string, error)

	DeleteProfileAvatar(ctx context.Context, profileId string) error

	UploadContentAsset(ctx context.Context, reader io.Reader, assetId string) (string, error)

	DeleteContentAsset(ctx context.Context, assetId string) error
}
