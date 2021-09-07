package consuos

import ("os"
        "ggiz.statistic.consumer.agent/consutil"
)

// 
type Status struct {
	Datetimeinit string
	Urldatabaseservice string
	Datetimefin string
	Port string
	AmqpConn string
	GeoIpService string
	TokenSmsService string
	LoginSmsService string
	NumDestSmsService string
	UrlSmsService string
}

// 
func (s *Status) Init() {
	s.Datetimeinit = consutil.GetDateTime()
}

// 
func (s *Status) ReadResourceEnv() {
	s.Urldatabaseservice  = os.Getenv("GGIZ_CONSUMER_STATISTIC_URL_CAD_DATABASE_SERVER_AGENT") 
	s.Port         = os.Getenv("GGIZ_CONSUMER_STATISTIC_PORT_SERVER_AGENT")	
	s.AmqpConn     = os.Getenv("GGIZ_CONSUMER_STATISTIC_URL_PATH")
	s.GeoIpService = os.Getenv("GGIZ_CONSUMER_STATISTIC_URL_GEO_IP")
	s.TokenSmsService = os.Getenv("GGIZ_CONSUMER_STATISTIC_TOKEN_SMS_SERVICE")
	s.LoginSmsService = os.Getenv("GGIZ_CONSUMER_STATISTIC_LOGIN_SMS_SERVICE")
	s.NumDestSmsService = os.Getenv("GGIZ_CONSUMER_STATISTIC_NUM_DEST_SMS_SERVICE")
	s.UrlSmsService = os.Getenv("GGIZ_CONSUMER_STATISTIC_URL_SMS_SERVICE")
} 

// 
func (s *Status) ReadFinResourceEnv() {
	s.Datetimefin = consutil.GetDateTime()
}
