
func Get(ctx *corporation.App, params url.Values, header http.Header) (resp *http.Response, err error) {
	accessToken, err := ctx.AccessToken.GetAccessTokenHandler(ctx)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, corporation.WXServerUrl+apiGet+"?access_token="+accessToken+"&"+params.Encode(), nil)
	if err != nil {
		return
	}

	req.Header = header

	return http.DefaultClient.Do(req)
}
