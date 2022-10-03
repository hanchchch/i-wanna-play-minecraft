package args

type Config struct {
	Region        string
	Profile       string
	SSHPubKeyPath string
	InstanceType  string
}

func LoadConfig() (Config, error) {
	return Config{
		Region:        "ap-northeast-2",
		Profile:       "personal",
		SSHPubKeyPath: "minecraft-pub.pem",
		InstanceType:  "c6i.large",
	}, nil
}
