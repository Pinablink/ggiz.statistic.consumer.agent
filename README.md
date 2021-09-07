# ggiz.statistic.consumer.agent

## Dependências 
Pessoal, deveria tratar melhor o mecanismo de dependência do projeto. Mas esta versão precisa ser feito de forma manual a inclusão dos pacotes. Esse projeto utiliza o pacote ***github.com/streadway/amqp*** como dependência para interação com o Ambiente RabbitMQ. E para o envio de SMS, o pacote ***github.com/Pinablink/g2ssms/send***.

## Apresentação
Microsistema embarcado em um RaspberryPI que recebe um estimulo temporal com solicitação de execução de serviço. Esse serviço obtêm dados no ambiente RabbitMQ. Com esses dados, realiza uma segunda consulta a um serviço de terceiro. Então, formata os dados que contêm informações sobre as visitas realizadas na semana no WebClient do GGIZ. Em seguida é enviado um ticket de serviço informando o status da operação. O ticket de serviço é um sms, o sistema envia a solicitação a um microserviço de envio de mensagem sms. O último passo consiste no envio dos dados ao microserviço Spring Boot que persiste essa informação na base MariaDB do GGIZ.

## Fluxo Simplificado do Sistema

![Alt text](FluxoRaspStatistic.png)

⚠ ***ATENÇÃO*** ⚠
***O Fluxo Simplificado acima, omitiu o envio de SMS pelo sistema***

## Sobre o Código Fonte
O Código Fonte serve apenas para consulta técnica. Podendo servir de fonte de informação para quem quiser entender como funciona a interação de um sistema escito em ***Golang*** com ***RabbitMQ*** e ***Requisição HTTP WEB***.


⚠ ***ATENÇÃO*** ⚠
O Código Fonte e a Documentação estão em constante evolução.

👍 Obrigado por sua visita!
