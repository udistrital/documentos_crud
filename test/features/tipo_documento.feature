Feature: Validate API responses
  DOCUMENTOS_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tipo_documento
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route             |bodyreq               |codres       |
  |GET   |/v1/tipo_documento|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/tipo_document |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tipo_document |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tipo_document |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tipo_document |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /tipo_documento       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples:
  |method|route             |bodyreq               |codres         |bodyres                 |
  |GET   |/v1/tipo_documento|./files/req/Vacio.json|200 OK         |./files/res3/Vok1.json  |
  |POST  |/v1/tipo_documento|./files/req/Yt1.json  |201 Created    |./files/res3/Vok2.json  |
  |POST  |/v1/tipo_documento|./files/req/Vacio.json|400 Bad Request|./files/res3/Ierr1.json |
  |POST  |/v1/tipo_documento|./files/req/Nt1.json  |400 Bad Request|./files/res3/Ierr2.json |
  |POST  |/v1/tipo_documento|./files/req/Nt2.json  |400 Bad Request|./files/res3/Ierr3.json |
  |POST  |/v1/tipo_documento|./files/req/Nt3.json  |400 Bad Request|./files/res3/Ierr4.json |
  |POST  |/v1/tipo_documento|./files/req/Nt4.json  |400 Bad Request|./files/res3/Ierr5.json |
  |POST  |/v1/tipo_documento|./files/req/Nt5.json  |400 Bad Request|./files/res3/Ierr6.json |
  |POST  |/v1/tipo_documento|./files/req/Nt6.json  |400 Bad Request|./files/res3/Ierr7.json | 
  |POST  |/v1/tipo_documento|./files/req/Nt7.json  |400 Bad Request|./files/res3/Ierr8.json | 
  |POST  |/v1/tipo_documento|./files/req/Nt8.json  |400 Bad Request|./files/res3/Ierr9.json | 
  |POST  |/v1/tipo_documento|./files/req/Nt9.json  |400 Bad Request|./files/res3/Ierr10.json| 
  |POST  |/v1/tipo_documento|./files/req/Nt10.json |400 Bad Request|./files/res3/Ierr11.json| 
  |POST  |/v1/tipo_documento|./files/req/Yt2.json  |400 Bad Request|./files/res3/Ierr12.json| 
  |PUT   |/v1/tipo_documento|./files/req/Yt2.json  |200 OK         |./files/res3/Vok2.json  |
  |GETID |/v1/tipo_documento|./files/req/Vacio.json|200 OK         |./files/res3/Vok2.json  |
  |DELETE|/v1/tipo_documento|./files/req/Vacio.json|200 OK         |./files/res3/Ino.json   |
  |DELETE|/v1/tipo_documento|./files/req/Vacio.json|404 Not Found  |./files/res3/Ierr13.json|
