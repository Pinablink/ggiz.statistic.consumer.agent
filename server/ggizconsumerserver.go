package server

import (
       "log"
       "bytes"
       "github.com/Pinablink/g2ssms/send"
       "io/ioutil"
       "encoding/json"
       "net/http"
       "ggiz.statistic.consumer.agent/consutil"
       "ggiz.statistic.consumer.agent/consuos"
       "ggiz.statistic.consumer.agent/cartaz"
       "github.com/streadway/amqp"
)

type dataCli struct {
   Ip string `json:"IP"`
   Usw int   `json:"USW"`
   Ush int   `json:"USH"`
   Sucess bool `json:"SUCESS"`
   Date string  `json:"DATE"`
}

type dataMsg struct {
   DataProc string `json:"DATAPROC"`
   ListDataCli []dataCli `json:"LISTDATACLI"`
}

// 
type GgizConsumer struct {
   RefSys *consuos.Status
   RefCartaz cartaz.Cartaz
}

//
type GgizCliDatageo struct {
   Codigopais string `json:"country_code"`
   Codigoregiao string `json:"region_code"`
   Timezone  string `json:"time_zone"`
}

var iRefCartaz cartaz.Cartaz
var iRefSys *consuos.Status
var iTicket *GgizTicketServer

//
type ggizEnvListInput struct {
   ListInput []ggizConsumerDataAMQP `json:"listInput"`
}

//
type ggizConsumerDataAMQP struct {
    DataAcesso string `json:"dataAcesso"`
    Ip string `json:"ip"`
    Paisdeacesso string `json:"paisdeacesso"`
    Regiao string `json:"regiao"`
    Timezone string `json:"timezone"`
    Sw int `json:"sw"`
    Sh int `json:"sh"`
}

//
func getMessages(w http.ResponseWriter, r *http.Request) {
   
   var i int

   iTicket = &GgizTicketServer{}
   iTicket.TituloLinhaMsg = "TK:Ggiz Consumer Statistic"
   iTicket.DataStrProcessamento = consutil.GetDateTime()
   iTicket.ErrorsVerIp = 0
	iTicket.ErrorsEnvio = 0
	iTicket.SucessoEnvio = 0 
   
   for i = 0; i < 8; i++ {
      iRefCartaz.ShowMessageGetAMQP((i + 1))
      solicInputStatisticDB(w, r)
   }

   msg, error := iTicket.GetMsg()
   
   if error != nil {
      iRefCartaz.ShowMessageErrGerMsgSms(error) 
   } else {
      obSSMS := &send.SSendSMS{}
	   obSSMS.UrlService = iRefSys.UrlSmsService
	   obSSMS.Acao = send.SendSms
	   obSSMS.Login = iRefSys.LoginSmsService
	   obSSMS.Token = iRefSys.TokenSmsService
	   obSSMS.Numero = iRefSys.NumDestSmsService
	   obSSMS.Msg = send.SMsg {
		   Msg: msg,
	   }

      idResponse, err := obSSMS.SendSMS()

      if err != nil {
         iRefCartaz.ShowMessageErrorSendSMS(err)
      } else {
         iRefCartaz.ShowMessageRespTicket(idResponse)
      }
   }

}

// Recebe o estimulo para se conectar a fila e obter os dados de acesso
func solicInputStatisticDB(w http.ResponseWriter, r *http.Request) {
   iRefCartaz.ShowMessageConsumerAMQP()
   amqpcconn, err := amqp.Dial(iRefSys.AmqpConn)

   if err != nil {
      iRefCartaz.ShowErrorConnAmqp()
      panic(err)
   }

   defer amqpcconn.Close()
   
   ch, err := amqpcconn.Channel()

   if err != nil {
      iRefCartaz.ShowErrorChAmqp()
      panic(err)
   }

   defer ch.Close()

   gMessages, ok, err := ch.Get(
     "input_statistic", //queue
      true, //auto-ack
   )

   if err != nil {
      iRefCartaz.ShowErrorGetMsgAmqp()
      panic(err)
   }

   if ok {
     refDataMsg := dataMsg{}
     strData := string(gMessages.Body)
     jsonerr := json.Unmarshal([]byte(strData), &refDataMsg) 

     if jsonerr != nil {
        iRefCartaz.ShowErrorParseAmqp()
     } else {
       var qtCli int = len(refDataMsg.ListDataCli)
       
       if qtCli > 0 {
         iRefCartaz.ShowMessageDataMsgDetc(qtCli)
         ggizCons := getDataGeo(refDataMsg)
         errorInputHttp := inputHttpDatStatistic(ggizCons)

        if errorInputHttp != nil {
            iTicket.SetErrorsEnvio()
        } else {
            iTicket.SetSucessoEnvio()
        }

       } else {
          iRefCartaz.ShowMessageZeroDataDetcAmqp()
       }
         
     }

   } else {
      iRefCartaz.ShowNoGetMsgAmqp()
   }
   
}

//
func (g *GgizConsumer) InitServer() {
   var port string = g.RefSys.Port
   iRefCartaz = g.RefCartaz
   iRefSys = g.RefSys
   g.RefCartaz.ReportIniServer()
   http.HandleFunc("/ggiz_consumer_statistic", getMessages)
   log.Fatal(http.ListenAndServe(port, nil))
}

//
func getDataGeo(refDataMsg dataMsg) ggizEnvListInput {
   arrRet := make([]ggizConsumerDataAMQP, len(refDataMsg.ListDataCli))
   returnInput := ggizEnvListInput{}

   for i, data := range refDataMsg.ListDataCli {
      strIp := data.Ip
      strUsw := data.Usw
      strUsh := data.Ush
      strDate := data.Date
      gDataCons :=  ggizConsumerDataAMQP{}
      gDataCons.DataAcesso = strDate
      gDataCons.Ip = strIp
      gDataCons.Sw  = strUsw
      gDataCons.Sh  = strUsh
      
      dataGeo, ok := httpGeoData(strIp)

      if ok {
         gDataCons.Paisdeacesso = dataGeo.Codigopais
         gDataCons.Regiao = dataGeo.Codigoregiao
         gDataCons.Timezone = dataGeo.Timezone
      } else {
         iTicket.SetErrorsVerIp()
      }


      arrRet[i] = gDataCons
   }

   returnInput.ListInput = arrRet

   return returnInput
}

//
func httpGeoData(refStrIp string) (GgizCliDatageo, bool) {

   refGeoIpService := iRefSys.GeoIpService
   urlGeoIpService := refGeoIpService + refStrIp
   resp, err := http.Get(urlGeoIpService)
   
   if err != nil {
      iRefCartaz.ShowMessageErrorGeoData()
      return GgizCliDatageo{}, false
   } else {

      defer resp.Body.Close() 
      mData, errParse := ioutil.ReadAll(resp.Body)

      if errParse != nil {
         
         iRefCartaz.ShowMessageErrorParser()
         return GgizCliDatageo{}, false

      } else {
         if resp.StatusCode == 200 {
            var dados string = string(mData)
            var ggizCliDataGeo GgizCliDatageo =  GgizCliDatageo{}
            parser := json.Unmarshal([]byte(dados), &ggizCliDataGeo)

            if parser != nil {
               iRefCartaz.ShowMessageErrorParser()
               return GgizCliDatageo{}, false
            } else {
               return ggizCliDataGeo, true
            }

         } else {
            iRefCartaz.ShowErrorRequestGeoData(resp.StatusCode)
            return GgizCliDatageo{}, false
         }
      }

   }

}

//
func inputHttpDatStatistic(inputList ggizEnvListInput) error {
   strJson, errMarshal := json.Marshal(inputList)

   if errMarshal != nil {
      return errMarshal
   } else {
      var inpByte = []byte(strJson)
      strUrl := iRefSys.Urldatabaseservice
      request, err := http.NewRequest("POST", strUrl, bytes.NewBuffer(inpByte))

      if err != nil {
         return err
      } else {
         request.Header.Set("Content-Type", "application/json")

         client := &http.Client{}
         _, err := client.Do(request)

         if err != nil {
            return err
         }
      }

   }

   return nil
}
