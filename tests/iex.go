package tests

import iex "github.com/goinvest/iexcloud/v2"

func GetIEXSandboxClient() *iex.Client {
	return iex.NewClient("Tsk_c1ee184113dc42bdae6907741c07d6ac", iex.WithBaseURL("https://sandbox.iexapis.com/v1"))
}
