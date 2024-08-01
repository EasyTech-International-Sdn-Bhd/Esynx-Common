package contracts

type IAppSalesInvoice interface {
	RetrieveUptoFuture() error
}

type IAppArInvoice interface {
	RetrieveUptoFuture() error
}

type IAppSalesInvoiceDetails interface {
	RetrieveUptoFuture() error
}

type IAppArInvoiceDetails interface {
	RetrieveUptoFuture() error
}
