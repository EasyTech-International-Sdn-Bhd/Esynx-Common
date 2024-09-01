package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointDebtorBranch      = "debtor/branch"
	EndPointDebtorBranchBulk  = "debtor/branch/bulk"
	EndPointDebtorBranchQuery = "debtor/branch/query"
)

type DebtorBranchResponse struct {
	Data    []entities.CmsCustomerBranch `json:"data"`
	Status  bool                         `json:"status"`
	Message string                       `json:"message"`
}

func (i *ApiClient) PostDebtorBranch(form request.DebtorBranchCreateOrUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointDebtorBranch)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkDebtorBranch(form request.DebtorBranchCreateOrUpdateBulkForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointDebtorBranchBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetDebtorBranch(branchCode string) (*DebtorBranchResponse, error) {
	var response DebtorBranchResponse
	resp, err := i.Reqwest.R().
		SetPathParam("branchCode", branchCode).
		SetResult(&response).
		Get(fmt.Sprintf("%s/{branchCode}", EndPointDebtorBranch))

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkDebtorBranch(branchCodes string) (*DebtorBranchResponse, error) {
	var response DebtorBranchResponse
	resp, err := i.Reqwest.R().
		SetPathParam("branchCode", branchCodes).
		SetResult(&response).
		Get(fmt.Sprintf("%s/{branchCode}", EndPointDebtorBranchBulk))

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutDebtorBranch(form request.DebtorBranchCreateOrUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointDebtorBranch)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkDebtorBranch(form request.DebtorBranchCreateOrUpdateBulkForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointDebtorBranchBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteDebtorBranch(form request.DebtorBranchDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointDebtorBranch)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkDebtorBranch(form request.DebtorBranchDeleteBulkForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointDebtorBranchBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) QueryDebtorBranch(query interface{}) (*DebtorBranchResponse, error) {
	var response DebtorBranchResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointDebtorBranchQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}
