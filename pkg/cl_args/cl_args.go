package cl_args

import "flag"

const (
	BatchName_CreateManagementUser             = "create_management_user"
	BatchName_CreateSCStartOrganizationHistory = "create_sc_start_organization_history"
)

// コマンドライン引数を格納する構造体
type ClArgs struct {
	BatchName       string
	ArgUserName     string
	ArgEmailAddress string
	ArgPassword     string
	ArgSCFrameID    int64
}

// コマンドライン引数を読み取る
func ReadClArgs() ClArgs {
	batchName := flag.String("batch_name", "default", "string batch name")
	argUserName := flag.String("user_name", "", "username for create management user")
	argEmailAddress := flag.String("email_address", "", "emailaddress for create management user")
	argPassword := flag.String("password", "", "password for create management user")
	argSCFrameID := flag.Int64("sc_frame_id", 0, "stress_check_frame_id which want to forcibly associate with organization_history")

	flag.Parse()
	return ClArgs{
		BatchName:       *batchName,
		ArgUserName:     *argUserName,
		ArgEmailAddress: *argEmailAddress,
		ArgPassword:     *argPassword,
		ArgSCFrameID:    *argSCFrameID,
	}
}
