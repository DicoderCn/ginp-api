package email

// EmailConfig 定义邮件配置结构体
type EmailConfig struct {
	Email string
	Pwd   string
	Port  int
	Host  string
}

// Init 初始化邮件客户端
func NewEemailClient(config EmailConfig) *EmailClient {
	return &EmailClient{config: config}
}
