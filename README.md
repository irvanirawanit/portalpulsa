## Example
```sh
gateway := portalpulsa{
	PortalUserId: "P20123x",
	PortalKey: "9634a97ccxxxxxxxxxx437674df7ae9c",
	PortalSecret: "228ab559ceea8a0ba9bc014ed1f5977edfe97cd1754e6b45485d54xxxxxxxxxx",
}
gateway.BeliTokenPLN(code string, phone string, idcust string, trxid_api string, no string) map[string]interface{}

gateway.CekSaldo() map[string]interface{}

gateway.Harga(code string) map[string]interface{}

gateway.RequestSaldoDeposit(bank string, nominal string) map[string]interface{}

gateway.StatusTransaksi(trxid_api string) map[string]interface{}

gateway.TopUp(code string, phone string, idcust string, trxid_api string, no string) map[string]interface{}
```