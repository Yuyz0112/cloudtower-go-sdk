package main

import (
	apiclient "github.com/Yuyz0112/cloudtower-go-sdk/client"
	"github.com/Yuyz0112/cloudtower-go-sdk/client/operations"
	"github.com/Yuyz0112/cloudtower-go-sdk/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func strPtr(s string) *string {
	return &s
}

func main() {
	transport := httptransport.New("192.168.31.209", "/v2/api/", []string{"http"})
	c := apiclient.New(transport, strfmt.Default)

	where := &models.VMWhereInput{
		NameContains: strPtr("test"),
	}
	getVmsParams := operations.NewGetVmsParams()
	getVmsParams.RequestBody = &models.GetVmsRequestBody{
		Where: where,
	}
	_, err := c.Operations.GetVms(getVmsParams)
	if err != nil {
		panic(err)
	}
	vmDiskWhere := &models.VMDiskWhereInput{
		VM: where,
	}
	getVmDiskParams := operations.NewGetVMDisksParams()
	getVmDiskParams.RequestBody = &models.GetVMDisksRequestBody{
		Where: vmDiskWhere,
	}
	diskRsp, err := c.Operations.GetVMDisks(getVmDiskParams)
	if err != nil {
		panic(err)
	}
	for _, v := range diskRsp.Payload {
		println(*v.ID, *v.Type)
	}
}
