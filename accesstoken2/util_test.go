package accesstoken2

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

func Test_base64EncodeStr(t *testing.T) {
	encodeStr := base64EncodeStr([]byte("hello"))
	testutil.Equal(t, "aGVsbG8=", encodeStr)
	decodeStr, err := base64DecodeStr(encodeStr)
	testutil.Nil(t, err)
	testutil.Equal(t, "hello", string(decodeStr))
}

func Test_compressZlib(t *testing.T) {
	compressed := compressZlib([]byte("hello"))
	testutil.Equal(t, "hello", string(decompressZlib(compressed)))
}

func assertTokenPayloadEqual(t *testing.T, expectedToken string, actualToken string) {
	t.Helper()
	if len(expectedToken) < VersionLength || len(actualToken) < VersionLength {
		t.Fatalf("token is shorter than the %d-byte version prefix", VersionLength)
	}
	testutil.Equal(t, expectedToken[:VersionLength], actualToken[:VersionLength])

	expectedCompressed, err := base64DecodeStr(expectedToken[VersionLength:])
	testutil.Nil(t, err)
	actualCompressed, err := base64DecodeStr(actualToken[VersionLength:])
	testutil.Nil(t, err)

	expectedPayload, err := decompressZlibWithError(expectedCompressed)
	testutil.Nil(t, err)
	actualPayload, err := decompressZlibWithError(actualCompressed)
	testutil.Nil(t, err)
	testutil.Equal(t, expectedPayload, actualPayload)
}

func Test_packUint16(t *testing.T) {
	buf := new(bytes.Buffer)
	err := packUint16(buf, 600)
	testutil.Nil(t, err)
	testutil.Equal(t, "5802", fmt.Sprintf("%x", buf.Bytes()))

	i, err := unPackUint16(buf)
	testutil.Nil(t, err)
	testutil.Equal(t, uint16(600), i)
}

func Test_packUint32(t *testing.T) {
	buf := new(bytes.Buffer)
	err := packUint32(buf, 600)
	testutil.Nil(t, err)
	testutil.Equal(t, "58020000", fmt.Sprintf("%x", buf.Bytes()))

	i, err := unPackUint32(buf)
	testutil.Nil(t, err)
	testutil.Equal(t, uint32(600), i)
}

func Test_packInt16(t *testing.T) {
	buf := new(bytes.Buffer)
	err := packInt16(buf, int16(-1))
	testutil.Nil(t, err)

	i, err := unPackInt16(buf)
	testutil.Nil(t, err)
	testutil.Equal(t, int16(-1), i)
}

func Test_packString(t *testing.T) {
	buf := new(bytes.Buffer)
	err := packString(buf, "hello")
	testutil.Nil(t, err)
	testutil.Equal(t, "050068656c6c6f", fmt.Sprintf("%x", buf.Bytes()))

	s, err := unPackString(buf)
	testutil.Nil(t, err)
	testutil.Equal(t, "hello", s)
}

func Test_packMapUint32(t *testing.T) {
	buf := new(bytes.Buffer)
	err := packMapUint32(buf, map[uint16]uint32{uint16(1): uint32(2)})
	testutil.Nil(t, err)
	testutil.Equal(t, "0100010002000000", fmt.Sprintf("%x", buf.Bytes()))

	m, err := unPackMapUint32(buf)
	testutil.Nil(t, err)
	testutil.Equal(t, uint32(2), m[1])
	testutil.Equal(t, "map[1:2]", fmt.Sprintf("%x", m))
}
