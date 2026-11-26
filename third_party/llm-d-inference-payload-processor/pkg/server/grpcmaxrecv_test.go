package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolveGRPCMaxRecvMsgSizeDefault(t *testing.T) {
	n, err := resolveGRPCMaxRecvMsgSizeFrom("")
	require.NoError(t, err)
	require.Equal(t, defaultGRPCMaxRecvMsgSize, n)
	require.Equal(t, 32*1024*1024, n)
}

func TestResolveGRPCMaxRecvMsgSizeOverrideBytes(t *testing.T) {
	n, err := resolveGRPCMaxRecvMsgSizeFrom("8388608")
	require.NoError(t, err)
	require.Equal(t, 8*1024*1024, n)
}

func TestResolveGRPCMaxRecvMsgSizeOverrideQuantity(t *testing.T) {
	n, err := resolveGRPCMaxRecvMsgSizeFrom("16Mi")
	require.NoError(t, err)
	require.Equal(t, 16*1024*1024, n)
}

func TestResolveGRPCMaxRecvMsgSizeInvalid(t *testing.T) {
	_, err := resolveGRPCMaxRecvMsgSizeFrom("not-a-size")
	require.Error(t, err)
}
