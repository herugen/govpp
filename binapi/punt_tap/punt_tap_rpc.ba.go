// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package punt_tap

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service punt_tap.
type RPCService interface {
	PuntTapPairAddDel(ctx context.Context, in *PuntTapPairAddDel) (*PuntTapPairAddDelReply, error)
	PuntTapPairGet(ctx context.Context, in *PuntTapPairGet) (RPCService_PuntTapPairGetClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PuntTapPairAddDel(ctx context.Context, in *PuntTapPairAddDel) (*PuntTapPairAddDelReply, error) {
	out := new(PuntTapPairAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PuntTapPairGet(ctx context.Context, in *PuntTapPairGet) (RPCService_PuntTapPairGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PuntTapPairGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PuntTapPairGetClient interface {
	Recv() (*PuntTapPairDetails, *PuntTapPairGetReply, error)
	api.Stream
}

type serviceClient_PuntTapPairGetClient struct {
	api.Stream
}

func (c *serviceClient_PuntTapPairGetClient) Recv() (*PuntTapPairDetails, *PuntTapPairGetReply, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, nil, err
	}
	switch m := msg.(type) {
	case *PuntTapPairDetails:
		return m, nil, nil
	case *PuntTapPairGetReply:
		if err := api.RetvalToVPPApiError(m.Retval); err != nil {
			return nil, nil, err
		}
		err = c.Stream.Close()
		if err != nil {
			return nil, nil, err
		}
		return nil, m, io.EOF
	default:
		return nil, nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
