package validator

//验证接口
type Validate interface {
	ValidateParam() (int, interface{})
}
