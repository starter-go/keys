package keys

// Driver 代表密钥的驱动
type Driver interface {
	Algorithm() string

	Class() Class

	ListRegistrations() []*DriverRegistration
}

// DriverRegistration 包含驱动的注册信息
type DriverRegistration struct {
	Algorithm string
	Class     Class
	Enabled   bool
	Priority  int
	Driver    Driver
}
