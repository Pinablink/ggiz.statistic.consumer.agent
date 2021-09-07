package cartaz

import ( "fmt"
         "ggiz.statistic.consumer.agent/consutil"
		 "ggiz.statistic.consumer.agent/consuos"
)

// 
type Cartaz struct {
	Refsys *consuos.Status

}

//
func (c Cartaz) ShowCartazIni() {
	fmt.Println(`------------------------------------------`)
	fmt.Println(`- GGIZ - CONSUMER STATISTIC WEBCLI       -`)
	fmt.Println(`------------------------------------------`)
}

//
func (c Cartaz) ShowDetailIni() {
	fmt.Println(consutil.GetDateTime(), ` - Inicializando...`)
	fmt.Println(consutil.GetDateTime(), ` - Lendo Recursos do Ambiente`)
	c.Refsys.ReadResourceEnv()
	fmt.Println(consutil.GetDateTime(),` - Recursos do Ambiente Processados`)
	fmt.Println(consutil.GetDateTime(),` - URL do serviço de inserção de   `)
	fmt.Println(consutil.GetDateTime(),`   de dados estatisticos na base   `)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.Urldatabaseservice)
	fmt.Println(consutil.GetDateTime(),` - Porta de escuta do serviço      `)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.Port)
	fmt.Println(consutil.GetDateTime(),` - Dados para Conexão e Obtenção de `)
	fmt.Println(consutil.GetDateTime(),`   mensagem na fila `)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.AmqpConn)
	fmt.Println(consutil.GetDateTime(),` - Recurso para obter dados geograficos `)
	fmt.Println(consutil.GetDateTime(),`   com finalidade estatistica `)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.GeoIpService)
	fmt.Println(consutil.GetDateTime(),` - Token para solicitar envio de SMS `)
	fmt.Println(consutil.GetDateTime(),`   a API `)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.TokenSmsService)
	fmt.Println(consutil.GetDateTime(),` - Login para validação no consumo do `)
	fmt.Println(consutil.GetDateTime(),`   serviço de envio de SMS`)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.LoginSmsService)
	fmt.Println(consutil.GetDateTime(),` - Numero para envio informativo do `)
	fmt.Println(consutil.GetDateTime(),`   TicketSMS ao administrador`)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.NumDestSmsService)
	fmt.Println(consutil.GetDateTime(),` - URL do microserviço de envio de `)
	fmt.Println(consutil.GetDateTime(),`   SMS`)
	fmt.Println(consutil.GetDateTime(),`:>>> `, c.Refsys.UrlSmsService)

	c.Refsys.ReadFinResourceEnv()
}

//
func (c Cartaz) ReportIni() {
   fmt.Println(`------------------------------------------`)	
   fmt.Println(` - Report de Inicialização`)
   fmt.Println(` - Sistema Inicializado em : `, c.Refsys.Datetimeinit)
   fmt.Println(` - Inicialização concluída em : `, c.Refsys.Datetimefin)
}

//
func (c Cartaz) ReportIniServer() {
	fmt.Println(`------------------------------------------`)
	fmt.Println(consutil.GetDateTime(),` - Inicializando e Configurando Servidor`)
}

//
func (c Cartaz) ShowSolicServer() {
	fmt.Println(consutil.GetDateTime(),` - Solicitação de Serviço Recebida`)
}

// 
func (c Cartaz) ShowMessageConsumerAMQP() {
	fmt.Println(consutil.GetDateTime(),` - Obter dados da Fila`)
}

//
func (c Cartaz) ShowErrorConnAmqp() {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu um problema na solicitação de conexão com a fila`)
}

//
func (c Cartaz) ShowErrorChAmqp() {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu um problema na obtenção do canal da fila`)
}

//
func (c Cartaz) ShowErrorGetMsgAmqp() {
	fmt.Println(consutil.GetDateTime(),` - Não foi possivel obter os dados na fila`)
}

//
func (c Cartaz) ShowNoGetMsgAmqp() {
	fmt.Println(consutil.GetDateTime(),` - Não existe mensagem na fila`)
}

//
func (c Cartaz) ShowErrorParseAmqp() {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu um erro na execução do parse dos dados da fila para um objeto interno`)
}

//
func (c Cartaz) ShowMessageDataMsgDetc(len int) {
	fmt.Println(consutil.GetDateTime(),` - Foram encontrado(s) `, len, ` recurso(s) para processamento`)
}

// 
func (c Cartaz) ShowMessageZeroDataDetcAmqp() {
	fmt.Println(consutil.GetDateTime(),` - Foi encontrado mensagem na fila. Mas a mensagem esta sem registro de acesso ao webcli`)
}

//
func (c Cartaz) ShowMessageErrorGeoData() {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu erro na obtenção dos dados geograficos para relatorio estatistico`)
}

//
func (c Cartaz) ShowMessageErrorParser() {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu um erro no Parser da mensagem obtida na fila para o tratamento interno do sistema`)
}

//
func (c Cartaz) ShowErrorRequestGeoData(responseCode int) {
	fmt.Println(consutil.GetDateTime(),` - Bad Response `, responseCode)
}

//
func (c Cartaz) ShowMessageGetAMQP(refd int) {
	fmt.Println(`================================================`)
	fmt.Println(consutil.GetDateTime(),` - Consultando dia `, refd)
}

//
func (c Cartaz) ShowMessageErrGerMsgSms(err error) {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu erro na obtenção da mensagem de envio do ticket`)
	fmt.Println(consutil.GetDateTime(),` - `,err)
	fmt.Println(`                                                                                 `)
}

//
func (c Cartaz) ShowMessageErrorSendSMS(err error) {
	fmt.Println(consutil.GetDateTime(),` - Ocorreu erro no envio do Ticket SMS`)
	fmt.Println(consutil.GetDateTime(),` - `, err)
	fmt.Println(`                                                                                 `)
}

//
func (c Cartaz) ShowMessageRespTicket(strTicket string) {
	fmt.Println(consutil.GetDateTime(),` - Ticket enviado. Num: `, strTicket)
}