
func Jssdk(ctx *corporation.App, params url.Values) (resp *http.Response, err error) {
	accessToken, err := ctx.AccessToken.GetAccessTokenHandler(ctx)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, corporation.WXServerUrl+apiJssdk+"?access_token="+accessToken+"&"+params.Encode(), nil)
	if err != nil {
		return
	}

	return http.DefaultClient.Do(req)
}