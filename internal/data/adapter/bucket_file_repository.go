package adapter

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"log/slog"

	"github.com/LazyCodeTeam/just-code-backend/internal/config"
)

const profileAvatarDir = "profile/avatar/"

type BucketFileRepository struct {
	bucketHandle *storage.BucketHandle
	config       *config.Config
}

func NewBucketFileRepository(
	bucketHandle *storage.BucketHandle,
	config *config.Config,
) *BucketFileRepository {
	return &BucketFileRepository{bucketHandle: bucketHandle, config: config}
}

func (r *BucketFileRepository) UploadProfileAvatar(
	ctx context.Context,
	reader io.Reader,
	profileId string,
) (string, error) {
	path := profileAvatarDir + profileId

	writer := r.bucketHandle.Object(path).NewWriter(ctx)
	defer writer.Close()
	_, err := io.Copy(writer, reader)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to upload file to bucket", "err", err)
		return "", err
	}

	return r.config.CdnBaseUrl + "/" + path, nil
}

func (r *BucketFileRepository) DeleteProfileAvatar(ctx context.Context, profileId string) error {
	path := profileAvatarDir + profileId
	err := r.bucketHandle.Object(path).Delete(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete file from bucket", "err", err)
		return err
	}

	return nil
}
