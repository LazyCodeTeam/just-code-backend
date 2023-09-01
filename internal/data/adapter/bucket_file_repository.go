package adapter

import (
	"context"
	"io"
	"log/slog"

	"cloud.google.com/go/storage"

	"github.com/LazyCodeTeam/just-code-backend/internal/config"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

const (
	profileAvatarDir = "profile/avatar/"
	contentAssetDir  = "content/asset/"
)

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

	return r.uploadObject(ctx, path, reader)
}

func (r *BucketFileRepository) DeleteProfileAvatar(ctx context.Context, profileId string) error {
	path := profileAvatarDir + profileId

	return r.deleteObject(ctx, path)
}

func (r *BucketFileRepository) UploadContentAsset(
	ctx context.Context,
	reader io.Reader,
	assetId string,
) (string, error) {
	path := contentAssetDir + assetId

	return r.uploadObject(ctx, path, reader)
}

func (r *BucketFileRepository) DeleteContentAsset(ctx context.Context, assetId string) error {
	path := contentAssetDir + assetId
	return r.deleteObject(ctx, path)
}

func (r *BucketFileRepository) deleteObject(ctx context.Context, path string) error {
	err := r.bucketHandle.Object(path).Delete(ctx)
	if err == storage.ErrObjectNotExist {
		return port.FileNotFoundError
	}
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete object from bucket", "err", err)
		return err
	}

	return nil
}

func (r *BucketFileRepository) uploadObject(
	ctx context.Context,
	path string,
	reader io.Reader,
) (string, error) {
	writer := r.bucketHandle.Object(path).NewWriter(ctx)
	defer writer.Close()
	_, err := io.Copy(writer, reader)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to upload file to bucket", "err", err)
		return "", err
	}

	return r.config.CdnBaseUrl + "/" + path, nil
}
