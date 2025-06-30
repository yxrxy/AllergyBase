package rpc

import (
	"github.com/yxrxy/AllergyBase/kitex_gen/biobank/biobankservice"
	"github.com/yxrxy/AllergyBase/kitex_gen/clinical/clinicalservice"
	"github.com/yxrxy/AllergyBase/kitex_gen/epidemiology/epidemiologyservice"
	"github.com/yxrxy/AllergyBase/kitex_gen/followup/followupservice"
	"github.com/yxrxy/AllergyBase/kitex_gen/user/userservice"
)

var (
	userClient         userservice.Client
	biobankClient      biobankservice.Client
	clinicalClient     clinicalservice.Client
	epidemiologyClient epidemiologyservice.Client
	followupClient     followupservice.Client
)

// Init 初始化所有RPC客户端
func Init() {
	InitUserRPC()
	InitBiobankRPC()
	InitClinicalRPC()
	InitEpidemiologyRPC()
	InitFollowupRPC()
}
