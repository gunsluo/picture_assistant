package assistant

import (
	"bytes"
	"fmt"

	backblaze "gopkg.in/kothar/go-backblaze.v0"
)

type B2PictureWrite struct {
	accountID      string
	applicationKey string
	bucketName     string
	b2             *backblaze.B2
	bucket         *backblaze.Bucket
	nameGenerator  func(dir string, info *PictureInfo) string
}

func NewB2PictureWrite(accountID, applicationKey, bucketName string) *B2PictureWrite {
	return &B2PictureWrite{
		accountID:      accountID,
		applicationKey: applicationKey,
		bucketName:     bucketName,
		nameGenerator:  defaultNameGenerator,
	}
}

func (write *B2PictureWrite) build() error {

	if write.b2 == nil {

		b2, err := backblaze.NewB2(backblaze.Credentials{
			AccountID:      write.accountID,
			ApplicationKey: write.applicationKey,
		})
		if err != nil {
			return err
		}

		write.b2 = b2
	}

	if write.bucket == nil {
		bucket, err := write.b2.Bucket(write.bucketName)
		if err != nil {
			return err
		}

		if bucket == nil {
			bucket, err = write.b2.CreateBucket(write.bucketName, backblaze.AllPrivate)
			if err != nil {
				return err
			}
		}

		write.bucket = bucket
	}

	return nil
}

func (write *B2PictureWrite) SetNameGenerator(nameGenerator func(dir string, info *PictureInfo) string) {
	write.nameGenerator = nameGenerator
}

func (write *B2PictureWrite) Write(buffer []byte, info *PictureInfo) error {

	err := write.build()
	if err != nil {
		return err
	}

	name, err := AutoName()
	if err != nil {
		return err
	}

	reader := bytes.NewReader(buffer)
	file, err := write.bucket.UploadFile(name, nil, reader)
	if err != nil {
		return err
	}

	//modify picture info
	info.Path = file.ID
	info.Name = file.Name
	info.Size = file.ContentLength
	info.Sha1 = file.ContentSha1

	return nil
}

func AutoNameGenerator(dir string, info *PictureInfo) string {
	name, err := AutoName()
	if err != nil {
		return defaultNameGenerator(dir, info)
	}

	return fmt.Sprintf("%s%s.%s", dir, name, info.Ext)
}
