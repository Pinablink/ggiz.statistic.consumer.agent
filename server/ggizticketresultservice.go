package server

import ("strconv"
	    "strings"
        "errors")

//
type GgizTicketServer struct {
	TituloLinhaMsg string
	DataStrProcessamento string
	ErrorsVerIp int
	ErrorsEnvio int
	SucessoEnvio int
}

//
func (ticket *GgizTicketServer) SetTituloLinhaMsg (strTituloMsg string) {
	ticket.TituloLinhaMsg = strTituloMsg
}

//
func (ticket *GgizTicketServer) SetDataProcessamento(strDataProc string) {
	ticket.DataStrProcessamento = strDataProc
}

//
func (ticket *GgizTicketServer) SetErrorsVerIp() {
	ticket.ErrorsVerIp = ticket.ErrorsVerIp + 1	
}

//
func (ticket *GgizTicketServer) SetErrorsEnvio() {
	ticket.ErrorsEnvio = ticket.ErrorsEnvio + 1
}

//
func (ticket *GgizTicketServer) SetSucessoEnvio() {
	ticket.SucessoEnvio = ticket.SucessoEnvio + 1
}

//
func (ticket *GgizTicketServer) GetMsg() (string, error) {
	
	var tituloLinhaMsgExist bool =  (len(strings.TrimSpace(ticket.TituloLinhaMsg)) > 0)
	var dataStrProc bool = (len(strings.TrimSpace(ticket.DataStrProcessamento)) > 0)
    var strMsgReturn string = ""

	if (tituloLinhaMsgExist && dataStrProc) {
		var statusProc bool = (ticket.ErrorsVerIp == 0 && ticket.ErrorsEnvio == 0 && ticket.SucessoEnvio == 0)

		if (statusProc) {
			return "", errors.New("Nenhum status de operação foi informado no ticket")
		} else {
			strMsgReturn = strings.TrimSpace(ticket.TituloLinhaMsg + ";DtPr:" + ticket.DataStrProcessamento + ";ErIp:" + strconv.Itoa(ticket.ErrorsVerIp) + ";ErEnv:" + strconv.Itoa(ticket.ErrorsEnvio) + ";OkEnv:" + strconv.Itoa(ticket.SucessoEnvio))
		}

	} else {
		return "", errors.New("Os dados de titulo e datetime precisam ser informados")
	}

	return strMsgReturn, nil
}