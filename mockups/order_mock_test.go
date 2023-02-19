package mock_proto

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/blablatov/stream-notls-grpc/proto"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moclProdInfoClient := NewMockOrderManagementClient(ctrl)

	id := "Sumsung S999"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2039"
	price := float32(777.0)
	req := &pb.Order{Id: id, Description: description, Price: price}

	moclProdInfoClient.EXPECT().AddOrder(gomock.Any(), &rpcMsg{msg: req}).
		Return(&pb.Order{Id: "Product:" + id}, nil)
	testAddProduct(t, moclProdInfoClient)
}

func testAddProduct(t *testing.T, client pb.OrderManagementClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := "Sumsung S999"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2039"
	price := float32(777.0)

	r, err := client.AddOrder(ctx, &pb.Order{Id: id, Description: description, Price: price})

	if err != nil || r.GetValue() == "Sumsung S999" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.GetValue())
}
