package config

type AppConfig struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Database int
	}
	SendSms struct {
		Account  string
		Password string
	}
	QiNiuYun struct {
		AccessKey string
		SecretKey string
		Bucket    string
		URL       string
	}
	RealName struct {
		SecretID  string
		SecretKey string
	}
	Email struct {
		SMTP string
	}
	AliPay struct {
		AppId      string
		PrivateKey string
		NotifyURL  string
		ReturnURL  string
	}
	Es struct {
		Address string
	}
	Jwt struct {
		SecretKey string
	}
	BaiDuYun struct {
		Ak string
		Sk string
	}
}
